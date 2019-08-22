package main

import (
	"errors"
	"github.com/sereiner/parrot/component"
	"github.com/sereiner/parrot/context"
)

type News struct {
	ID    string `json:"id"`
	Title string `json:"title"`
}
type QueryHandler struct {
	container component.IContainer
}

func NewQueryHandler(container component.IContainer) (u *QueryHandler) {
	return &QueryHandler{container: container}
}

func (u *QueryHandler) Handle(ctx *context.Context) (r interface{}) {
	// 从请求中获取参数
	tp := ctx.Request.GetInt("t", 0)
	ctx.Response.SetContentType(context.ContentTypes[tp])
	m := ctx.Request.GetInt("m", 0)
	switch m {
	case 0:
		return `{"id":1}`
	case 1:
		return map[string]interface{}{
			"a": "b",
		}

	case 2:
		return "success"
	case 3:
		return 100
	case 4:
		return `<?xml version="1.0" encoding="UTF-8"?>
		<note>
			<to>Tove</to>
			<from>Jani</from>
			<heading>Reminder</heading>
			<body>Don't forget me this weekend!</body>
		</note>`
	case 5:
		return `<!DOCTYPE html><html></html>`
	case 6:
		return News{
			ID:    "1000",
			Title: "最新新闻",
		}
	default:
		return errors.New("系统繁忙")
	}
}

func(u *QueryHandler) PostHandle(ctx *context.Context) (r interface{}) {

	return "POST"
}