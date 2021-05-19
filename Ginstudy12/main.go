package main

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

//定义模型
type User struct {
	ID   int64
	Name string
	Age  int64
}

func main() {

	db, err := gorm.Open("mysql", "root:123456@(localhost)/test?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	//模型与数据表对应
	db.AutoMigrate(&User{})

	//创建
	u := User{Name: "qimi", Age: 18}
	//返回bool,为空返回true
	fmt.Println(db.NewRecord(u))
	db.Create(&u)
	fmt.Println(db.NewRecord(u))

	var x = new(User)
	db.First(x)
	fmt.Println(x)
}
