package main

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type UserInfo struct {
	ID     uint
	Name   string
	Gender string
	Hobby  string
}

func main() {
	//连接MySQL数据库
	db, err := gorm.Open("mysql", "root:123456@(localhost)/test?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	//创建表,自动迁移(把结构体和数据表相对应)
	db.AutoMigrate(&UserInfo{})
	//插入一行数据
	u1 := UserInfo{ID: 1, Name: "qimi", Gender: "男", Hobby: "游戏"}
	db.Create(&u1)

	//查询
	var u UserInfo
	db.First(&u) //查出来第一条数据放进结构体实例u中
	fmt.Printf("u:%#v\n", u)

	//更新
	db.Model(&u).Update("hobby", "双色球")
	fmt.Printf("u:%#v\n", u)

	//删除
	db.Delete(&u)
}
