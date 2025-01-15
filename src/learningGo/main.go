package main

import (
	"learningGo/http/controller"

	"github.com/gin-gonic/gin"
)

func main() {
	apiServer := gin.New()
	// 加入日志
	apiServer.Use(gin.Logger())
	// apiServer.Use(handle.GinLogHandle())
	/**
	假装代码
	假装代码
	假装代码
	假装代码
	假装代码
	假装代码
	// 注释2
	// 注释2
	// 注释2
	// 注释2
	// 注释2
	// 注释2
	// 注释2
	// 注释2
	注释2
	注释2
	注释2
	注释2
	*/

	imageGroup := apiServer.Group("/image")

	imageGroup.GET("/:name", controller.ImageShow)
	// 路由声明
	// routers.NewApiSss(r.Group(""))
	// routers.NewPCApi(r.Group(""))

	// 设置api端口
	_ = apiServer.Run(":8099")
	// 注释
	// 注释
	// 注释
	// 注释
	// 注释
	// 注释
	// 注释
	// 注释
	// 注释
	// 注释
	// 注释
}
