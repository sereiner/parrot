package main

import (
	"fmt"
	"github.com/sereiner/parrot/component"
	"github.com/sereiner/parrot/context"
	"github.com/sereiner/parrot/example/greeter"
	"github.com/sereiner/parrot/example/helloworld"
	"github.com/sereiner/parrot/example/order"
	"github.com/sereiner/parrot/parrot"
	"google.golang.org/grpc"
)

func main() {

	app := parrot.NewApp(
		parrot.WithPlatName("apiserver"),
		parrot.WithSystemName("apiserver"),
		parrot.WithServerTypes("once-rpc"),
		parrot.WithDebug(),
		parrot.WithPbRegister(RegisterPb))

	app.Conf.RPC.SetMainConf(`{"address":":8032"}`)

	app.Conf.ONCE.SetSubConf("task", `{"tasks":[{"cron":"@after 5s","service":"/order/query"}]}`)

	app.Initializing(func(c component.IContainer) error {

		conn, err := c.GetConn("hello_service", "hello_service")
		if err != nil {
			return err
		}

		c.SetRpcClient("hello_service", helloworld.NewGreeterClient(conn))

		return nil
	})

	app.Once("/order/query", order.NewQueryHandler)
	app.Micro("/order", func(ctx *context.Context) (r interface{}) {
		return "success"
	})
	app.Start()
}

func RegisterPb(c component.IContainer, s *grpc.Server) {

	helloworld.RegisterGreeterServer(s, greeter.NewGreeter(c))
	fmt.Println("注册完成")
}
