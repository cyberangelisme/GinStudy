package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/json", func(c *gin.Context) {
		//方法一，使用map
		// data := map[string]interface{}{
		// 	"name":    "小娃子",
		// 	"message": "helloworld",
		// 	"age":     18,
		// }
		//gin.H就等于map[string]interface{}
		data := gin.H{
			"name":    "小娃子",
			"message": "helloworld",
			"age":     18,
		}

		c.JSON(http.StatusOK, data)
	})
	//方法二，结构体
	type msg struct {
		Name    string
		Age     int
		Message string
	}
	r.GET("another_json", func(c *gin.Context) {
		data := msg{
			"小网址",
			18,
			"hello,yep",
		}

		c.JSON(http.StatusOK, data)
	})
	r.Run(":8080")
}
