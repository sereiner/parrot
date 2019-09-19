package component

import (
	"context"
	"fmt"
	"github.com/asaskevich/govalidator"
	"github.com/sereiner/library/balancer"
	"github.com/sereiner/parrot/conf"
	"github.com/sereiner/parrot/registry"
	"google.golang.org/grpc"
	"google.golang.org/grpc/balancer/roundrobin"
	"google.golang.org/grpc/resolver"
	"sync"
	"time"
)

const RPCTypeNameInVar = "rpc"

const RPCNameInVar = "rpc"

type IComponentRPC interface {
	GetConn(platName,service string) (*grpc.ClientConn, error)
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

func (c *CustomRPC) GetConn(platName,service string) (*grpc.ClientConn, error) {

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
	r := balancer.NewResolver(rpcConf.Register,platName, service)
	resolver.Register(r)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*1)
	conn, err := grpc.DialContext(
		ctx,
		r.Scheme()+"://authority/"+service,
		grpc.WithInsecure(),
		grpc.WithBalancerName(roundrobin.Name),
		grpc.WithBlock())
	defer cancel()
	if err != nil {
		return nil, fmt.Errorf("创建grpc客户端失败,请确保服务端存在")
	}
	return conn, nil
}

func (c *CustomRPC) SetRpcClient(name string, r interface{}) {
	c.rpcMap.Store(name, r)
}

func (c *CustomRPC) GetRpcClient(name string) (value interface{}, ok bool) {
	return c.rpcMap.Load(name)
}
