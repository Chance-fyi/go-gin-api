package routers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Init(r *gin.Engine) {
	r.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
}
