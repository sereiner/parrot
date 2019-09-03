package main

import (
	ct "context"
	"fmt"
	"github.com/sereiner/parrot/component"
	"github.com/sereiner/parrot/context"
	pb "github.com/sereiner/parrot/example/helloworld"
)

type TestModel struct {
	MerRefundID string `json:"mer_refund_id" valid:"required"`
	RequestID   string `json:"request_id" valid:"required"`
	OrderNo     string `json:"order_no" valid:"required"`
}

type QueryHandler struct {
	container component.IContainer
}

func NewQueryHandler(container component.IContainer) (u *QueryHandler) {
	return &QueryHandler{container: container}
}

func (u *QueryHandler) Handle(ctx *context.Context) (r interface{}) {
	// 从请求中获取参数

	v, ok := u.container.GetRpcClient("hello_service")
	if !ok {
		return fmt.Errorf("grpc 客户端错误")
	}
	res, err := v.(pb.GreeterClient).SayHello(ct.Background(), &pb.HelloRequest{Name: "world haha" })
	if err != nil {
		return err
	}
	return res
}
