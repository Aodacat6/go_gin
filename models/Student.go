package models

// 学生结构
type Student struct {
	Id      int    `json:"id" xorm:"'id' pk autoincr"`
	Name    string `json:"name" :"name"`
	Age     int    `json:"age" :"age"`
	Address string `json:"address" :"address"`
}
