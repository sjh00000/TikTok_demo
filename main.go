package main

import (
	"github.com/gin-gonic/gin"
	"tiktok/service"
)

func main() {

	//消息服务器
	go service.RunMessageServer()

	//创建一个默认的 Gin 引擎实例
	r := gin.Default()

	initRouter(r)

	_ = r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
