package engines

import (
	"fmt"
	"strings"
	"time"

	"github.com/sereiner/library/types"
	"github.com/sereiner/parrot/component"
	"github.com/sereiner/parrot/context"
)

//RPCProxy rpc 代理服务
func (r *ServiceEngine) RPCProxy() component.ServiceFunc {
	return func(ctx *context.Context) (r interface{}) {
		header, _ := ctx.Request.Http.GetHeader()
		cookie, _ := ctx.Request.Http.GetCookies()
		for k, v := range cookie {
			header[k] = v
		}
		header["method"] = strings.ToUpper(ctx.Request.GetMethod())
		input := ctx.Request.GetRequestMap()
		timeout := ctx.Request.Setting.GetInt("timeout", 3)
		response := ctx.RPC.AsyncRequest(ctx.Service, header, input, true)
		status, result, params, err := response.Wait(time.Second * time.Duration(timeout))
		if err != nil {
			err = fmt.Errorf("rpc.proxy %v(%d)", err, status)
		}
		ctx.Response.SetParams(types.GetIMap(params))
		if err != nil {
			ctx.Response.SetStatus(status)
			return err
		}
		ctx.Response.SetStatus(status)
		return result
	}
}
