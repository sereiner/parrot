package component

import (
	"context"
	"fmt"
	"github.com/asaskevich/govalidator"
	"github.com/coreos/etcd/clientv3"
	"github.com/coreos/etcd/mvcc/mvccpb"
	logger "github.com/sereiner/library/log"
	"github.com/sereiner/parrot/conf"
	"github.com/sereiner/parrot/registry"
	"google.golang.org/grpc"
	"google.golang.org/grpc/balancer/roundrobin"
	"google.golang.org/grpc/resolver"
	"strings"
	"sync"
	"time"
)

const RPCTypeNameInVar = "rpc"

const RPCNameInVar = "rpc"

type IComponentRPC interface {
	GetConn(service string) (*grpc.ClientConn, error)
	SetRpcClient(name string, r interface{})
	GetRpcClient(name string) (value interface{}, ok bool)
}

type CustomRPC struct {
	IContainer
	name   string
	rpcMap sync.Map
}

func NewCustomRPC(c IContainer, name ...string) *CustomRPC {
	if len(name) > 0 {
		return &CustomRPC{IContainer: c, name: name[0], rpcMap: sync.Map{}}
	}
	return &CustomRPC{IContainer: c, name: DBNameInVar, rpcMap: sync.Map{}}
}

func (c *CustomRPC) GetConn(service string) (*grpc.ClientConn, error) {

	cacheConf, err := c.IContainer.GetVarConf(RPCTypeNameInVar, RPCNameInVar)
	if err != nil {
		return nil, fmt.Errorf("%s %v", registry.Join("/", c.GetPlatName(), "var", RPCTypeNameInVar, RPCNameInVar), err)
	}

	var rpcConf conf.RPCConf
	if err = cacheConf.Unmarshal(&rpcConf); err != nil {
		return nil, err
	}
	if b, err := govalidator.ValidateStruct(&rpcConf); !b {
		return nil, err
	}
	r := NewResolver(rpcConf.Register, service)
	resolver.Register(r)
	ctx , cancel := context.WithTimeout(context.Background(),time.Second * 2)
	conn, err := grpc.DialContext(
		ctx,
		r.Scheme()+"://authority/"+service,
		grpc.WithInsecure(),
		grpc.WithBalancerName(roundrobin.Name),
		grpc.WithBlock())
	defer cancel()
	if err != nil {
		panic(fmt.Errorf("创建grpc客户端失败,请确保服务端存在"))
	}
	return conn, nil
}

func (c *CustomRPC) SetRpcClient(name string, r interface{}) {
	c.rpcMap.Store(name, r)
}

func (c *CustomRPC) GetRpcClient(name string) (value interface{}, ok bool) {
	return c.rpcMap.Load(name)
}

// resolver is the implementaion of grpc.resolve.Builder
type Resolver struct {
	schema  string
	target  string
	service string
	cli     *clientv3.Client
	cc      resolver.ClientConn
	logger  *logger.Logger
}

// NewResolver return resolver builder
// target example: "http://127.0.0.1:2379,http://127.0.0.1:12379,http://127.0.0.1:22379"
// service is service name
func NewResolver(target string, service string) resolver.Builder {
	return &Resolver{
		schema:  service,
		target:  target,
		service: service,
		logger:  logger.GetSession("parrot", logger.CreateSession()),
	}
}

// Scheme return etcdv3 schema
func (r *Resolver) Scheme() string {
	return r.schema
}

// ResolveNow
func (r *Resolver) ResolveNow(rn resolver.ResolveNowOption) {
}

// Close
func (r *Resolver) Close() {
}

// Build to resolver.Resolver
func (r *Resolver) Build(target resolver.Target, cc resolver.ClientConn, opts resolver.BuildOption) (resolver.Resolver, error) {
	var err error

	r.cli, err = clientv3.New(clientv3.Config{
		Endpoints: strings.Split(r.target, ","),
	})
	if err != nil {
		return nil, fmt.Errorf("grpclb: create clientv3 client failed: %v", err)
	}

	r.cc = cc

	go r.watch(fmt.Sprintf("/%s/%s/", r.schema, r.service))

	return r, nil
}

func (r *Resolver) watch(prefix string) {
	addrDict := make(map[string]resolver.Address)

	update := func() {
		addrList := make([]resolver.Address, 0, len(addrDict))
		for _, v := range addrDict {
			addrList = append(addrList, v)
		}
		r.cc.NewAddress(addrList)
	}

	resp, err := r.cli.Get(context.Background(), prefix, clientv3.WithPrefix())
	if err == nil {
		for i := range resp.Kvs {
			addrDict[string(resp.Kvs[i].Value)] = resolver.Address{Addr: string(resp.Kvs[i].Value)}
		}
	}

	update()

	rch := r.cli.Watch(context.Background(), prefix, clientv3.WithPrefix(), clientv3.WithPrevKV())
	for n := range rch {
		for _, ev := range n.Events {
			switch ev.Type {
			case mvccpb.PUT:
				addrDict[string(ev.Kv.Key)] = resolver.Address{Addr: string(ev.Kv.Value)}
				r.logger.Info(" grpc节点更新 ", string(ev.Kv.Key))
			case mvccpb.DELETE:
				delete(addrDict, string(ev.PrevKv.Key))
				r.logger.Info(" grpc节点更新下线 ", string(ev.Kv.Key))
			}
		}
		update()
	}
}
