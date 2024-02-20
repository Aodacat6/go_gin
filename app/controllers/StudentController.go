package controllers

import (
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"go_gin/app/service"
	"go_gin/models"
	"net/http"
	"strconv"
)

// 保存
func Save() gin.HandlerFunc {
	return func(context *gin.Context) {
		//application/x-www-form-urlencoded
		//name, _ := context.GetPostForm("name")
		//age, _ := context.GetPostForm("age")
		//address, _ := context.GetPostForm("address")
		//var name, address string
		//var age int
		student := new(models.Student)
		//入参json绑定结构体
		//application/json
		/*		err := context.ShouldBind(&student)
				if err == nil {
					name = student.Name
					age = student.Age
					address = student.Address
				}*/
		context.ShouldBind(&student)
		//创建一个切片
		students := []models.Student{}
		students = append(students, *student)
		students = append(students, *student)
		service.SaveBatch(students)
		context.JSON(http.StatusOK, map[string]any{
			"code": 0,
			"msg":  "",
		})
	}

}

func UpdateStudent() gin.HandlerFunc {
	return func(context *gin.Context) {
		//定义接收
		student := new(models.Student)
		//绑定变量到自定义
		context.ShouldBind(&student)
		service.UpdateStudent(student)
	}
}

func DeleteStudent() gin.HandlerFunc {
	return func(context *gin.Context) {
		defer func() {
			err := recover()
			if err != nil {
				context.JSON(http.StatusUnauthorized, gin.H{
					"msg": err,
				})
			}
		}()
		session := sessions.Default(context)
		fmt.Println(session.Get("auth"))
		if session.Get("auth") == nil || session.Get("auth") == "" || session.Get("auth") != "ok" {
			panic("请先登录")
		}
		value, _ := context.Get("aaa")
		fmt.Println("接收到中间件传入的变量aaa： ", value)
		//context.MustGet("")
		//定义接收
		student := new(models.Student)
		//绑定变量到自定义
		context.ShouldBind(&student)
		service.DeleteStudent(student)
	}
}

func GetOneStudentById() gin.HandlerFunc {
	return func(context *gin.Context) {
		id, _ := context.GetQuery("id")
		atoi, _ := strconv.Atoi(id)
		student := service.GetOneStudentById(atoi)
		context.JSON(http.StatusOK, student)
	}
}

func ListStudent() gin.HandlerFunc {
	return func(context *gin.Context) {

		students := service.ListStudent()
		context.JSON(http.StatusOK, students)
	}
}
