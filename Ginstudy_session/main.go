package main

//gin demo
import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin_learing/session/gin_session"
)

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")
	//初始化全局的mgrobj
	gin_session.InitMgr("redis", "127.0.0.1:6379")
	//session作为全局的中间件
	r.Use(gin_session.SessionMiddleware(gin_session.MgrObj))
	r.Any("/login", loginHandler)
	r.GET("/index", indexHandler)
	r.GET("/home", homeHandler)
	r.GET("/vip", AuthMiddleware, vipHandler)

	//没有匹配到走下面
	r.NoRoute(func(c *gin.Context) {
		c.HTML(http.StatusNotFound, "404.html", nil)
	})
	r.Run()
}
