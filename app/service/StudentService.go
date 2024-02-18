package service

import (
	"errors"
	"fmt"
	"go_gin/conf"
	"go_gin/models"
)

// p批量保存
func SaveBatch(students []models.Student) {
	if students == nil || len(students) == 0 {
		errors.New("空集合不保存")
	}
	connection := conf.GetDBConnection()
	//测试数据库链接数是否生效
	for i := 0; i < 100; i++ {
		go func() {
			session := connection.NewSession()
			defer func() {
				session.Close()
				fmt.Println("===关闭了 数据库 链接 的session")
			}()
			insert, err := session.Insert(students)
			if err != nil {
				errors.New("插入报错了" + err.Error())
			}
			fmt.Printf("插入成功，共插入 %d 条数据 \n", insert)
		}()
	}

}

func UpdateStudent(student models.Student) {
	//获取数据库链接
	connection := conf.GetDBConnection()
	//1、根据id更新
	connection.ID().Update()

}
