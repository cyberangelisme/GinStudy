package main

import (
	"html/template"

	"github.com/gin-gonic/gin"
)

func newa(c *gin.Context) {
	c.HTML(200, "templates/index.html", nil)
}

//静态文件
//html页面上的样式，css，js，图片

func main() {
	r := gin.Default()
	//gin框架中给模板添加自定义函数
	r.SetFuncMap(template.FuncMap{
		"safe": func(str string) template.HTML {
			return template.HTML(str)
		},
	})
	//加载静态文件,(当渲染的页面用到静态文件时
	r.Static("/xxx", "./statics") //xxx代替statics文件继承
	//模板解析
	//r.LoadHTMLFiles("templates/index.html")
	r.LoadHTMLFiles("templates/index.html")

	//模板渲染
	r.GET("/posts/index", func(c *gin.Context) {
		//http请求,gin.H表示map[string]interface{}类型
		c.HTML(200, "posts/index.html", gin.H{
			"title": "liwenzhou.com",
		})
	})

	r.GET("/users/index", func(c *gin.Context) {
		//http请求,gin.H表示map[string]interface{}类型，默认转译
		c.HTML(200, "users/index.html", gin.H{
			"title": "<a href='https://liwenzhou.com'>李文周博客</a>",
		})
	})

	r.GET("/newa", newa)
	r.Run(":8080") //启动server
}
