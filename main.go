package main

import (
	"github.com/gin-gonic/gin"
	"go_gin/router"
)

func main() {
	//全局启动对象，类型与springboot  application
	engine := gin.Default()

	//gin配置
	//[GIN] 2024/02/07 - 13:44:33 | 200 |            0s |             ::1 | GET      "/test"
	gin.SetMode(gin.DebugMode)
	//[GIN] 2024/02/07 - 13:45:51 | 200 |            0s |             ::1 | GET      "/test"
	//gin.SetMode(gin.ReleaseMode)

	/*	engine.GET("/test", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"message": "hello world",
		})
	})*/

	//注册路由
	router.Router(&engine.RouterGroup)

	//配置数据库

	//engine.Run()
	engine.Run(":8081")

	//D:\my_exe
}
