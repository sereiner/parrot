package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/sereiner/parrot/conf"
)

func JwtWriter(cnf *conf.MetadataConf) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Next()
		context := getCTX(ctx)
		if context == nil {
			return
		}
		setJwtResponse(ctx, cnf, context.Response.GetParams()["__jwt_"])
	}
}
