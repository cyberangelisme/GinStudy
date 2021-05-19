package main

import (
	"encoding/gob"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

type user struct {
	Name string
	Age  int
}

func main() {
	//注册结构体，使其可以跨路由存取
	gob.Register(user{})
	r := gin.Default()
	store := cookie.NewStore([]byte("secret"))

	r.Use(sessions.Sessions("mysess", store))

	r.GET("/hello", func(c *gin.Context) {
		session := sessions.Default(c)

		//第一次访问写入
		if session.Get("hello") != "world" {
			session.Set("hello", "world")
			//记得写入
			session.Save()
		}
		session.Set("user", user{"hanyun", 30})
		session.Save()
		c.JSON(200, gin.H{"hello": session.Get("hello")})
	})

	r.GET("/user", func(c *gin.Context) {
		session := sessions.Default(c)
		user := session.Get("user")
		c.JSON(200, gin.H{"user": user})
	})

	r.Run(":8080")
}
