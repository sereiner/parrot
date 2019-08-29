package main

import (
	"github.com/sereiner/parrot/component"
	"github.com/sereiner/parrot/context"
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

	ctx.Log.Info(ctx.Request.GetString("order_no"))

	input := &TestModel{}
	if err := ctx.Request.Bind(input); err != nil {
		return err
	}

	return "success"
}
