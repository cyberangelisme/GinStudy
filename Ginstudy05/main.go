package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.LoadHTMLFiles("./login.html", "./index.html")
	r.GET("/login", func(c *gin.Context) {
		c.HTML(200, "login.html", nil)
	})
	//处理login页面下的post请求
	r.POST("/login", func(c *gin.Context) {
		//获取post中form提交的数据，name关键字
		username := c.PostForm("username")
		password := c.PostForm("password")
		c.HTML(200, "index.html", gin.H{
			"Name": username,
			"Pass": password,
		})
	})
	r.Run(":8080")
}
