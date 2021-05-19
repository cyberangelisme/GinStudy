package main

import (
	"gowebstudy/ginstudy/Ginstudy_cookie/session"
	"net/http"

	"github.com/gin-gonic/gin"
)

//定义全局中间键
func AuthMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		//获取客户端cookie并校验,找不到cookie时err!=nil
		if cookie, err := c.Cookie("abc"); err == nil {
			//找到了cookie,value值为"123"
			if cookie == "123" {
				//先调用后续处理函数
				c.Next()
				//因为此中间件只做校验，所以之间返回
				return
			}
		}
		//找不到cookie，返回错误err，不进行后续处理
		c.JSON(http.StatusUnauthorized, gin.H{"error": "err"})
		c.Abort()
		return
	}
}

func main() {
	// r := gin.Default()
	// r.GET("/login", func(c *gin.Context) {
	// 	c.SetCookie("abc", "123", 60, "/", "localhost", false, true)
	// 	//注册成功，返回信息
	// 	c.String(200, "login sucesses！")
	// })
	// r.GET("/home", AuthMiddleWare(), func(c *gin.Context) {
	// 	c.JSON(200, gin.H{"data": "home"})
	// })
	// r.Run(":8080")

	//创建全局管理器SessionMgr
	var globalSessions *session.MemorySessionMgr
	//初始化globalSessions
	globalSessions = session.NewMemorySessionMgr()
	

}
