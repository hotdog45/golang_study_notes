package main

import (
	"fmt"
	"gin_xm/global"
	"gin_xm/initialize"

	"github.com/gin-gonic/gin"
)

func main() {
	//1.初始化yaml配置
	initialize.InitConfig()

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run(fmt.Sprintf(":%d", global.Settings.Port))
}
