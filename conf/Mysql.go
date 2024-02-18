package conf

import (
	"errors"
	"fmt"
	_ "github.com/go-sql-driver/mysql" //_ 表示导入了 github.com/go-sql-driver/mysql 包，但是没有使用该包中的任何函数或变量
	"go_gin/models"
	"xorm.io/xorm"
)

type MYSQL_CONF struct {
	UserName  string
	Password  string
	IpAddress string
	Port      int
	DbName    string
	Charset   string
}

// 定义mysql配置
var mysqlConfig = MYSQL_CONF{
	"root",
	"123456",
	"127.0.0.1",
	3306,
	"mytest",
	"utf8mb4",
}

var engine *xorm.Engine

// 初始化
func init() {
	GetDBConnection()
}

// 统一管理数据库链接
func GetDBConnection() *xorm.Engine {
	if engine == nil {
		//构建数据库链接信息
		//root:123456@tcp(127.0.0.1:3306)/go_test?charset=utf8mb4
		dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s", mysqlConfig.UserName, mysqlConfig.Password, mysqlConfig.IpAddress, mysqlConfig.Port, mysqlConfig.DbName, mysqlConfig.Charset)
		newEngine, err := xorm.NewEngine("mysql", dataSourceName)
		if err != nil {
			panic("获取数据库链接失败：" + err.Error())
		}
		fmt.Println("数据库链接创建成功")

		//配置数据库engine
		//打印sql
		newEngine.ShowSQL(true)
		//连接池：设置最大支持链接数
		newEngine.SetMaxOpenConns(10)
		//连接池：设置最大存活链接数
		newEngine.SetMaxIdleConns(3)
		engine = newEngine
		return newEngine
	} else {
		return engine
	}
}

// 同步表结构
func SyncTable() int {
	connection := GetDBConnection()
	if connection == nil {
		errors.New("mysql链接获取失败")
	}
	err := connection.Sync(new(models.Student))
	if err != nil {
		errors.New("建表失败")
	}
	fmt.Println("建表成功！")
	return 0
}
