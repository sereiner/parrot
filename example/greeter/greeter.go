package greeter

import (
	"context"
	"github.com/sereiner/parrot/component"
	"github.com/sereiner/parrot/example/helloworld"
)

type Greeter struct {
	c component.IContainer
}

func NewGreeter(c component.IContainer) *Greeter {
	return &Greeter{
		c: c,
	}
}

func (g *Greeter) SayHello(ctx context.Context, req *helloworld.HelloRequest) (*helloworld.HelloReply, error) {
	return &helloworld.HelloReply{
		Message: "hello: " + req.GetName(),
	}, nil
}
