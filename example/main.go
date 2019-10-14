package main

import (
	"github.com/sereiner/parrot/component"
	"github.com/sereiner/parrot/context"
	"github.com/sereiner/parrot/example/order"
	"github.com/sereiner/parrot/parrot"
)

func main() {

	app := parrot.NewApp(
		parrot.WithPlatName("apiserver"),
		parrot.WithSystemName("apiserver"),
		parrot.WithServerTypes("once-rpc-api"),
		parrot.WithDebug())

	app.Conf.RPC.SetMainConf(`{"address":":8032"}`)

	app.Conf.ONCE.SetSubConf("task", `{"tasks":[{"cron":"@after 5s","service":"/order/query"}]}`)

	//app.Conf.Plat.SetVarConf("ding", "ding", `{
	//	"webhook":"https://oapi.dingtalk.com/robot/send?access_token=3340852f9ce446e6bed2cb8b32ea1a2fb30b8ded538dc1c2735d4d07730c5bc6",
	//	"monitor":"http://monitor.manyoujing.net"
	//}`)

	app.Initializing(func(c component.IContainer) error {

		return nil
	})

	app.Once("/order/query", order.NewQueryHandler)
	app.Micro("/order", func(ctx *context.Context) (r interface{}) {
		return "success"
	})
	app.Start()
}
