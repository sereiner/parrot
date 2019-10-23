package main

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"github.com/sereiner/parrot/component"
	"github.com/sereiner/parrot/context"
	"github.com/sereiner/parrot/example/order"
	"github.com/sereiner/parrot/parrot"
	"go.uber.org/zap"
	"net/http"
	"time"
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
	app.Micro("/download", func(ctx *context.Context) (r interface{}) {

		bytesBuffer := &bytes.Buffer{}

		writer := csv.NewWriter(bytesBuffer)
		writer.Write([]string{
			"书名",
		})

		writer.Write([]string{
			"heheh",
		})

		writer.Flush()

		ctx.GinContext.Writer.Header().Set("Content-Disposition", fmt.Sprintf("attachment;filename=回收端上架报表-%s.csv", time.Now().Format("2006-01-02 15:04:05")))

		zap.L().Info("3. 返回数据")
		ctx.GinContext.Data(http.StatusOK, "text/csv", bytesBuffer.Bytes())
		return
	})
	app.Start()
}
