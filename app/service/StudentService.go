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
	session := connection.NewSession()
	//开启事务
	session.Begin()
	//异常处理
	defer func() {
		err1 := recover()
		if err1 != nil {
			fmt.Println("执行回滚")
			session.Rollback()
		} else {
			//提交事务
			fmt.Println("执行提交")
			session.Commit()
		}
		//必须要关闭，否则回滚不生效
		session.Close()
	}()
	//delete(param)  param里是要删除的条件
	/*	i, err := connection.ID(student.Id).Delete(student)
		if err != nil {
			fmt.Println("删除失败： ", err.Error())
			return
		}
		fmt.Println("删除成功： ", i)
	*/
	//自定义sql语句
	exec, err := session.Exec("delete from student where id = ?", student.Id)
	if err != nil {
		fmt.Println("执行错误： ", err.Error())
		return
	}
	panic("error")

	fmt.Println("执行结果：", exec)
}

func GetOneStudentById(id int) models.Student {
	//获取数据库链接
	/*	connection := conf.GetDBConnection()
		quote, err := connection.QueryString("select * from student where id = ?", id)
		if err != nil {
			fmt.Println("查询报错：", err.Error())
			panic(err)
		}
		var jsonstr []byte
		for i := 0; i < len(quote); i++ {
			jsonstr, _ = json.Marshal(quote[i])
			fmt.Printf("index ： %d , 查到了：%s \n", i, jsonstr)
		}
		//todo 先看看返回数据的结构
		student := new(models.Student)
		json.Unmarshal(jsonstr, student)
		return *student*/
	/*
		connection := conf.GetDBConnection()
		quote, err := connection.QueryInterface("select * from student where id = ?", id)
		if err != nil {
			fmt.Println("查询报错：", err.Error())
			panic(err)
		}
		//var jsonstr []byte
		student := new(models.Student)
		for i := 0; i < len(quote); i++ {
			m := quote[i]
			student.Id = int(m["id"].(int32))
			student.Name = m["name"].(string)
			student.Age = int(m["age"].(int32))
			student.Address = m["address"].(string)
			fmt.Printf("index ： %d , 查到了：%s \n", i, student)
		}
		//todo 先看看返回数据的结构
		//student := new(models.Student)
		//json.Unmarshal(jsonstr, student)
		return *student*/

	connection := conf.GetDBConnection()
	m := models.Student{Name: "damiao"}
	//如果有多条记录，也只会返回一条
	get, _ := connection.Where("id = ?", id).Get(&m)
	if get == true {

		return m
	}
	fmt.Println("没有找到")
	return *new(models.Student)
}

func ListStudent() []models.Student {
	connection := conf.GetDBConnection()
	var students []models.Student
	//list
	//err := connection.Cols("name").Distinct("name").Limit(1, 1).Find(&students)
	err := connection.Limit(10).Find(&students)
	if err != nil {
		panic(err)
	}

	return students
}
