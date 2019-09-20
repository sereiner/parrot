package component

import (
	"context"
	"fmt"
	"github.com/sereiner/library/balancer"
	"google.golang.org/grpc"
	"google.golang.org/grpc/balancer/roundrobin"
	"google.golang.org/grpc/resolver"
	"sync"
	"time"
)

type IComponentRPC interface {
	GetConn(platName, service string) (*grpc.ClientConn, error)
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

func (c *CustomRPC) GetConn(platName, serverName string) (*grpc.ClientConn, error) {

	r := balancer.NewResolver("", platName, serverName)
	resolver.Register(r)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*1)
	conn, err := grpc.DialContext(
		ctx,
		r.Scheme()+"://authority/",
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
