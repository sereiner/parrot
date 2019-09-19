package main

import (
	"github.com/sereiner/parrot/component"
	"github.com/sereiner/parrot/context"

	//pb "github.com/sereiner/parrot/example/helloworld"
	"github.com/sereiner/parrot/parrot"
)

func main() {

	app := parrot.NewApp(
		parrot.WithPlatName("apiserver"),
		parrot.WithSystemName("apiserver"),
		parrot.WithServerTypes("once-rpc"),
		parrot.WithDebug())


	app.Conf.RPC.SetMainConf(`{"address":":8032"}`)

	app.Conf.Plat.SetVarConf("rpc", "rpc", `
		{
			"register":"http://localhost:2379"
	}`)

	app.Conf.ONCE.SetSubConf("task", `{"tasks":[{"cron":"@after 5s","service":"/order/query"}]}`)

	app.Initializing(func(c component.IContainer) error {

		//conn, err := c.GetConn("hello_service")
		//if err != nil {
		//	return err
		//}
		//
		//c.SetRpcClient("hello_service", pb.NewGreeterClient(conn))

		return nil
	})

	app.Once("/order/query", NewQueryHandler)
	app.Micro("/order", func(ctx *context.Context) (r interface{}) {
		return "success"
	})
	app.Start()
}
