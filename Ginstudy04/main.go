package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/web", func(c *gin.Context) {
		//获取浏览器get请求携带的querystring参数
		name := c.Query("query")
		c.JSON(200, gin.H{
			"name": name,
		})
	})
	r.Run(":8080")
}
