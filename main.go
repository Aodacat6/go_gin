package main

import (
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
	"go_gin/app/middleware"
	"go_gin/router"
	"reflect"
)

func main() {
	//全局启动对象，类型与springboot  application
	engine := gin.Default()

	//Use中间件，gin的default里默认带了
	//engine.Use(gin.Logger())
	//出问题不会崩溃服务
	//engine.Use(gin.Recovery())
	//注册我的日志中间件
	engine.Use(middleware.InterfaceLog())
	//注册session中间
	//生成session存储store
	//方法一：基于cookie的存储，在服务器重启后，依旧有效
	//store := cookie.NewStore([]byte("123456"))
	//方法二：使用redis存储session
	//
	store, _ := redis.NewStore(10, "tcp", "localhost:6379", "", []byte("123456"))
	engine.Use(sessions.Sessions("mySession", store))

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

func main11() {
	//var num int = 10
	//
	//var a *int = &num
	////*int：*放在变量类型上，表示指针类型，指向地址
	////* ： 放在变量上，表示解引用，将地址值映射为对应的真实值
	//fmt.Println(*a)

	//i := new(int)
	//*i = 19
	//fmt.Println(i)
	//fmt.Println(*i)

	//反射
	var a int = 10

	//fmt.Println(reflect.TypeOf(a))
	//fmt.Println(reflect.ValueOf(a))
	//反射重新赋值
	of := reflect.ValueOf(&a)
	of.Elem().SetInt(12)
	fmt.Println(a)
}
