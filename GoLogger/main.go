package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	SetupLogger()
	simpleHttpGet("www.baidu.com")
	simpleHttpGet("www.google.com")
	simpleHttpGet("http://www.google.com")
	simpleHttpGet("http://www.baidu.com")
}

func SetupLogger() {
	logFileLocation, _ := os.OpenFile("./test.log", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0744)
	log.SetOutput(logFileLocation)
}

func simpleHttpGet(url string) {
	resp, err := http.Get(url)
	if err != nil {
		log.Printf("Error fetching url %s: %s", url, err.Error())
	} else {
		log.Printf("Status Code for %s: %s", url, resp.Status)
		resp.Body.Close() //关闭http请求，为了提高效率，http.Get 等请求的 TCP 连接是不会关闭的（再次向同一个域名请求时，复用连接），所以必须要手动关闭。
	}
}
