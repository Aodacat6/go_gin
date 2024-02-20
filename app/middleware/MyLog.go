package middleware

import (
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

// 自定义接口日志中间件
func InterfaceLog() gin.HandlerFunc {
	return func(context *gin.Context) {
		start := time.Now()
		//设置上下文变量
		context.Set("aaa", "aaa,aaa")
		//到达控制器之前的处理
		log.Println("接收到请求：")

		//调用下一个中间件，或者控制器的方法
		context.Next()

		//end := time.Now().UnixMilli()
		//useTime := (end - start) / 1000
		elapsed := time.Since(start)
		log.Println("请求处理结束，耗时：", elapsed)

	}
}
