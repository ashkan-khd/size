package main

import (
	limits "github.com/gin-contrib/size"
	"github.com/gin-gonic/gin"
	"net/http"
)

func handler(ctx *gin.Context) {
	val := ctx.PostForm("b")
	if len(ctx.Errors) > 0 {
		return
	}
	ctx.String(http.StatusOK, "got %s\n", val)
}

func main() {
	rtr := gin.Default()
	rtr.Use(limits.RequestSizeLimiter(2750), limits.AbortIfTooLarge("b", limits.TextFormFile))
	rtr.POST("/", handler)
	rtr.Run(":8080")
}
