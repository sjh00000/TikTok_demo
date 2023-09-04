package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"path/filepath"
	"strconv"
	"tiktok/database"
	"tiktok/pjdata"
	"tiktok/public"
)

type VideoListResponse struct {
	pjdata.Response
	VideoList []pjdata.Video `json:"video_list"`
}

// Publish check token then save upload file to public directory
func Publish(c *gin.Context) {
	token := c.PostForm("token")
	title := c.PostForm("title")

	//查询当前用户是否存在
	if _, exist := usersLoginInfo[token]; !exist {
		c.JSON(http.StatusOK, pjdata.Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
		return
	}

	//获取视频数据
	data, err := c.FormFile("data")

	//上传出错，返回
	if err != nil {
		c.JSON(http.StatusOK, pjdata.Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}

	//获取文件的名字以及userID来创建视频文件在public文件夹下
	filename := filepath.Base(data.Filename)
	user := usersLoginInfo[token]
	finalName := fmt.Sprintf("%d_%s", user.Id, filename)
	saveFile := filepath.Join("./public/", finalName)
	if err := c.SaveUploadedFile(data, saveFile); err != nil {
		c.JSON(http.StatusOK, pjdata.Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, pjdata.Response{
		StatusCode: 0,
		StatusMsg:  finalName + " uploaded successfully",
	})

	public.CutFeed(finalName)

	//合成common数据结构,将video传到数据库中
	var video = pjdata.Video{
		Id:            database.AddVideoNum(),
		Author:        user,
		PlayUrl:       "http://47.109.78.46:8080/zip_" + finalName,
		CoverUrl:      "http://47.109.78.46:8080/" + public.GetFeedCover("zip_"+finalName),
		FavoriteCount: 0,
		CommentCount:  0,
		IsFavorite:    false,
		Title:         title,
	}
	database.AddVideo(video)

}

// PublishList all users have same publish video list
func PublishList(c *gin.Context) {
	token := c.Query("token")
	userId, _ := strconv.Atoi(c.Query("user_id"))

	if _, exist := usersLoginInfo[token]; !exist {
		c.JSON(http.StatusOK, pjdata.Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
		return
	}

	user := usersLoginInfo[token]
	userVideos := database.SearchAuthorVideo(int64(userId), user)

	c.JSON(http.StatusOK, VideoListResponse{
		Response: pjdata.Response{
			StatusCode: 0,
		},
		VideoList: userVideos,
	})
}
