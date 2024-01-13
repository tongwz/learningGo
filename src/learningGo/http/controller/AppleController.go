package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func AppleShow(c *gin.Context) {
	appleName := c.Param("name")

	fmt.Printf("apple名称是：%s \n", appleName)
}
