package boot

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-gin-api/routers"
	"net/http"
)

type route struct{}

var Route = route{}

func (*route) Init() {
	err := SetRouter().Run(":8000")
	if err != nil {
		fmt.Println(err.Error())
	}
}

// SetRouter 设置路由与服务启动分开方便单元测试
func SetRouter() *gin.Engine {
	r := gin.Default()
	routers.Init(r)
	setup404Handler(r)
	return r
}

// 处理404请求
func setup404Handler(r *gin.Engine) {
	r.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(http.StatusNotFound, gin.H{
			"code":    404,
			"message": "not found",
		})
	})
}
