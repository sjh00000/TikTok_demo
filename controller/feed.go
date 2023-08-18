package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"tiktok/database"
	"tiktok/pjdata"
	"time"
)

type FeedResponse struct {
	//嵌入字段
	pjdata.Response
	VideoList []pjdata.Video `json:"video_list,omitempty"`
	NextTime  int64          `json:"next_time,omitempty"`
}

// Feed same demo video list for every request
func Feed(c *gin.Context) {
	//从数据库中获取视频列表
	VideoListData := database.SearchVideo()
	c.JSON(http.StatusOK, FeedResponse{
		Response:  pjdata.Response{StatusCode: 0},
		VideoList: VideoListData,
		NextTime:  time.Now().Unix(),
	})
}
