package middlewares

import "github.com/gin-gonic/gin"

func Authenticate(ctx *gin.Context) {
	// TODO: authentication logic
	ctx.Next()
}
