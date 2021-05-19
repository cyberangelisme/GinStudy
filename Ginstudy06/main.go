package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserInfo struct {
	Username string `form:username json:"username"`
	Password string `form:password json:"password"`
}

func main() {
	r := gin.Default()
	r.GET("/user", func(c *gin.Context) {
		// username := c.Query("username")
		// password := c.Query("password")
		// u := UserInfo{
		// 	username: username,
		// 	password: password,
		// }
		var u UserInfo
		err := c.ShouldBind(&u) //!!!需要传地址
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"status": "ok",
			})
		}
	})

	r.POST("/form", func(c *gin.Context) {
		var u UserInfo
		err := c.ShouldBind(&u) //!!!需要传地址
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		} else {
			fmt.Printf("%#v\n", u)
			c.JSON(http.StatusOK, gin.H{
				"status": "ok",
			})
		}
	})

	r.POST("/json", func(c *gin.Context) {
		var u UserInfo
		err := c.ShouldBind(&u) //!!!需要传地址
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		} else {
			fmt.Printf("%#v\n", u)
			c.JSON(http.StatusOK, gin.H{
				"status": "ok",
			})
		}
	})
	r.Run(":9090")
}
