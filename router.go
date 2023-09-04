package main

import (
	"github.com/gin-gonic/gin"
	"tiktok/controller"
	"tiktok/public"
)

func initRouter(r *gin.Engine) {

	//将 /static 路径与本地的 public 目录关联，用于提供静态资源的访问
	r.Static("/static", "./public")

	r.GET("/:filename", public.FileHandler)

	//创建抖音路由组apiRouter
	apiRouter := r.Group("/douyin")

	/***基本api***/
	//视频流
	apiRouter.GET("/feed/", controller.Feed)
	//用户信息
	apiRouter.GET("/user/", controller.UserInfo)
	//用户注册
	apiRouter.POST("/user/register/", controller.Register)
	//用户登录
	apiRouter.POST("/user/login/", controller.Login)
	//投稿接口
	apiRouter.POST("/publish/action/", controller.Publish)
	//发布列表
	apiRouter.GET("/publish/list/", controller.PublishList)

	/***互动api***/
	//点赞
	apiRouter.POST("/favorite/action/", controller.FavoriteAction)
	//喜欢列表
	apiRouter.GET("/favorite/list/", controller.FavoriteList)
	//评论
	apiRouter.POST("/comment/action/", controller.CommentAction)
	//评论列表
	apiRouter.GET("/comment/list/", controller.CommentList)

}
