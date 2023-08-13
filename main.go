package main

import (
	"github.com/gin-gonic/gin"
	"tiktok/database"
)

func main() {

	//消息服务器
	//go service.RunMessageServer()

	//创建一个默认的 Gin 实例
	r := gin.Default()

	initRouter(r)

	_ = r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

func init() {
	//连接数据库
	go database.InitDatabase()
}
