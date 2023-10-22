package main

import (
	"github.com/gin-gonic/gin"
	"learningGo/http/controller"
)

func main() {
	apiServer := gin.New()
	// 加入日志
	apiServer.Use(gin.Logger())
	// apiServer.Use(handle.GinLogHandle())

	imageGroup := apiServer.Group("/image")

	imageGroup.GET("/:name", controller.ImageShow)
	// 路由声明
	// routers.NewApis(r.Group(""))
	// routers.NewPCApi(r.Group(""))

	// 设置api端口
	_ = apiServer.Run(":8099")
}
