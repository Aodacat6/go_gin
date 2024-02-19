package controllers

import (
	"github.com/gin-gonic/gin"
	"go_gin/app/service"
	"go_gin/models"
	"net/http"
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
		//定义接收
		student := new(models.Student)
		//绑定变量到自定义
		context.ShouldBind(&student)
		service.DeleteStudent(student)
	}
}
