package main

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

func wtf(c *gin.Context) {
	fmt.Println("wtf ....in")

	//跨中间件存取值
	name, ok := c.Get("name")
	if !ok {
		name = "匿名用户"
	}
	c.JSON(200, gin.H{
		"msg": name,
	})

	fmt.Println("wtf..out")
}

func m1(c *gin.Context) {
	fmt.Println("m1 in...")
	//开始计时
	start := time.Now()
	//开始调用后续函数，调用完了在执行本函数剩余部分
	c.Next()
	cost := time.Since(start)
	fmt.Println("cost: ", cost)
	fmt.Println("m1 ...out")
}

func m2(c *gin.Context) {
	fmt.Println("m2.....in")
	c.Set("name", "hanser") //key,value
	//c.Abort()

	fmt.Println("m2.....out")
}

// //登录中间件
// func LoMiddleWare(c *gin.Context) {
// 	//if 是登录用户
// 	//r.Next()
// 	//else
// 	//r.Abort()
// }

//闭包形式
//LoMiddleWare(true)执行检查
//LoMiddleWare(false)不执行检查
func LoMiddleWare(doCheck bool) gin.HandlerFunc {
	//连接数据库
	//或一些其他准备工作

	return func(c *gin.Context) {
		//可以引用闭包外参数
		if doCheck {
			fmt.Println("执行了docheck检查")
		} else {
			c.Next()
		}				
	}
}
func main() {
	//Default默认使用两个中间件，Logger日志,Recovery报错返回500响应码
	//gin.New()创建一个没有任何中间件的路由
	r := gin.Default()
	r.Use(m1, m2) //全局注册中间件，每个请求都执行m1函数
	r.GET("/index", wtf)

	//路由组注册中间件方法1：
	xxgroup := r.Group("/xx", LoMiddleWare(true))
	{
		xxgroup.GET("/index", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"msg": "xxgroup",
			})
		})
	}

	//路由组注册中间件方法2：
	xxgroup2 := r.Group("/xx2")
	xxgroup2.Use(LoMiddleWare(false))
	{
		xxgroup2.GET("/index", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"msg": "xxgroup2",
			})
		})
	}
	r.Run()
}
