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
		parrot.WithServerTypes("api-mqc"),
		parrot.WithDebug())

	app.Conf.RPC.SetMainConf(`{"address":":9001"}`)


	app.Conf.MQC.SetSubConf("server", `
		{
			"proto":"nsq",
			"addrs":["localhost:4151"],
			"db":1,
			"dial_timeout":10,
			"read_timeout":10,
			"write_timeout":10,
			"pool_size":10
	}
	`)


	app.Conf.MQC.SetSubConf("queue", `{
	     "queues":[
			{
	           "queue":"test#ch",
	           "service":"/mqc"
			}
	   ]
	}`)

	app.Conf.Plat.SetVarConf("queue", "queue", `
		{
			"proto":"nsq",
			"addrs":["localhost:4151"],
			"db":1,
			"dial_timeout":10,
			"read_timeout":10,
			"write_timeout":10,
			"pool_size":10
	}
	`)


	app.Initializing(func(c component.IContainer) error {

		return nil
	})


	app.Micro("/order", order.NewQueryHandler)

	app.MQC("/mqc", func(ctx *context.Context) (r interface{}) {
		ctx.Log.Info(ctx.Request.GetBodyMap())
		return "success"
	})

	app.Start()
}
