package main

import (
	"fmt"
	"learningGo/http/controller"

	"github.com/gin-gonic/gin"
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
	list := []int{1, 2, 3, 4}
	for key := range list {
		fmt.Printf("我们打印了key:%d, 没有打印value \n", key)
	}

	// 设置api端口
	_ = apiServer.Run(":8099")
}
