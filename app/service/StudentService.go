package service

import (
	"errors"
	"fmt"
	"go_gin/conf"
	"go_gin/models"
	"log"
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

func UpdateStudent(student *models.Student) {
	defer func() {
		a := recover()
		//fmt.Println("recover:  ..  ", a)
		if a != nil {
			log.Panicln(a)
		} else {
			log.Println("更新成功")
		}
	}()
	//获取数据库链接
	connection := conf.GetDBConnection()
	//1、根据id更新
	/*	update, err := connection.ID(10).Update(student)
		if err != nil {
			panic(errors.New("更新操作错误" + err.Error()))
			return
		}
		fmt.Printf("更新成功，更新数据量为：%d \n", update)*/

	//2、直接更新
	update, err := connection.Where("id = ?", student.Id).Update(student)
	if err != nil {
		panic(errors.New("更新操作错误" + err.Error()))
		return
	}
	fmt.Printf("更新成功，更新数据量为：%d \n", update)
}

func DeleteStudent(student *models.Student) {
	//获取数据库链接
	connection := conf.GetDBConnection()
	//delete(param)  param里是要删除的条件
	/*	i, err := connection.ID(student.Id).Delete(student)
		if err != nil {
			fmt.Println("删除失败： ", err.Error())
			return
		}
		fmt.Println("删除成功： ", i)
	*/
	//自定义sql语句
	exec, err := connection.Exec("delete from student where id = ?", student.Id)
	if err != nil {
		fmt.Println("执行错误： ", err.Error())
		return
	}
	fmt.Println("执行结果：", exec)
}
