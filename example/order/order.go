package order

import (
	"github.com/sereiner/parrot/component"
	"github.com/sereiner/parrot/context"
	"github.com/sereiner/parrot/example/queues"
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

	return queues.Send(u.container,"test",`{"name":"jack","age":18}`)

}
