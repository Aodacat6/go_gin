package controllers

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

// 获取用户
func GetById() gin.HandlerFunc {
	return func(context *gin.Context) {
		// your code here
		/*		user := new(models.User)
				user.Username = "damiao"
				user.Password = "112233"
				context.JSON(http.StatusOK, user)*/
		name := "dadada"
		//返回string
		context.String(http.StatusOK, "返回了字符串，%s", name)
	}

}

// param 传参
func GetAndRtn() gin.HandlerFunc {
	return func(context *gin.Context) {
		context.Header("languate", "ch")
		//获取param参数
		//id := context.Query("id")
		//增加判断
		id, ok := context.GetQuery("id")
		if !ok {
			context.Error(errors.New("id不能为空"))
			return
		}
		context.JSON(http.StatusOK, gin.H{
			"id": id,
		})
	}
}
