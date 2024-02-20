package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go_gin/app/controllers"
)

func Router(group *gin.RouterGroup) {

	defer func() {
		e := recover()
		fmt.Println(e)
	}()
	//登录页
	loginGroup := group.Group("auth")
	loginGroup.GET("/login", controllers.Login())

	//user 页面
	//user := router.Group("/user")
	userGroup := group.Group("user")
	userGroup.GET("/getbyid", controllers.GetById())
	userGroup.DELETE("/deletebyid")
	userGroup.GET("/getAndRtn", controllers.GetAndRtn())

	//student 页面
	studentGroup := group.Group("student")
	//student := router.Group("/student")
	studentGroup.GET("/getone", controllers.GetOneStudentById())
	studentGroup.GET("/list", controllers.ListStudent())
	studentGroup.POST("/save", controllers.Save())
	studentGroup.PUT("/edit", controllers.UpdateStudent())
	studentGroup.DELETE("/delete", controllers.DeleteStudent())

}
