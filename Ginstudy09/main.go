package main

import "github.com/gin-gonic/gin"

func axxx(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "ok",
	})
}

func main() {
	r := gin.Default()
	userGroup := r.Group("/user")
	{
		userGroup.GET("/index", axxx)
		userGroup.GET("/login", axxx)
		userGroup.POST("/login", axxx)

	}
	shopGroup := r.Group("/shop")
	{
		shopGroup.GET("/index", axxx)
		shopGroup.GET("/cart", axxx)
		shopGroup.POST("/checkout", axxx)
	}
	r.Run()
}
