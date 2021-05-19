package main

import (
	"net/http"
	"path"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.LoadHTMLFiles("./index.html")
	r.GET("/index", func(c *gin.Context) {
		c.HTML(200, "index.html", nil)
	})
	r.POST("/upload", func(c *gin.Context) {
		//从请求中读取文件，保存在本地
		f, err := c.FormFile("f1") //从请求中获取参数
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		} else {
			//保存在服务端本地
			//dst := fmt.Sprintf("./%s", f.Filename)
			dst := path.Join("./", f.Filename)
			c.SaveUploadedFile(f, dst)
			c.JSON(http.StatusOK, gin.H{
				"status": "ok",
			})
		}
	})
	r.Run(":8080")
}
