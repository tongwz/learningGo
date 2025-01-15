package controller

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func ImageShow(c *gin.Context) {
	imageName := c.Param("name")

	fmt.Printf("图片名称是：%s \n", imageName)
}
