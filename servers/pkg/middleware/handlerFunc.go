package middleware

import "github.com/sereiner/parrot/servers/pkg/dispatcher"

type HandlerFunc func(ctx *dispatcher.Context)

func (h HandlerFunc) Handle(ctx *dispatcher.Context) {
	h(ctx)
}
