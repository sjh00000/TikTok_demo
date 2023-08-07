package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"tiktok/controller"
)

func initRouter(r *gin.Engine) {

	//将 /static 路径与本地的 public 目录关联，用于提供静态资源的访问
	r.Static("/static", "./public")
	//加载templates目录下的html文件作为模板（加载并非渲染出来）
	r.LoadHTMLGlob("templates/*")

	//渲染html模板
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title": "Main website",
		})
	})

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

	/***社交api***/
	//关注
	apiRouter.POST("/relation/action/", controller.RelationAction)
	//关注列表
	apiRouter.GET("/relation/follow/list/", controller.FollowList)
	//粉丝列表
	apiRouter.GET("/relation/follower/list/", controller.FollowerList)
	//好友列表
	apiRouter.GET("/relation/friend/list/", controller.FriendList)
	//聊天记录
	apiRouter.GET("/message/chat/", controller.MessageChat)
	//发送信息
	apiRouter.POST("/message/action/", controller.MessageAction)
}
