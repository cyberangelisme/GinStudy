package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/index", func(c *gin.Context) {
		// c.JSON(http.StatusOK, gin.H{
		// 	"status": "ok",
		// })
		//跳转至某个网址
		c.Redirect(http.StatusMovedPermanently, "http://www.sogo.com")
	})

	r.GET("/a", func(c *gin.Context) {
		//跳转至/b另一个路由对应的函数
		c.Request.URL.Path = "/b"
		r.HandleContext(c) //继续后续处理
	})

	r.GET("/b", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "b",
		})
	})
	r.Run(":8080")
}
