package main

import (
	"github.com/gin-gonic/gin"
	"tiktok/controller"
)

func main() {
	controller.MapInit()
	//创建一个默认的 Gin 实例
	r := gin.Default()

	initRouter(r)

	_ = r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")

}
