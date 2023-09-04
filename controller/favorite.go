package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"tiktok/database"
	"tiktok/pjdata"
)

// FavoriteAction no practical effect, just check if token is valid
func FavoriteAction(c *gin.Context) {
	token := c.Query("token")
	userID := usersLoginInfo[token].Id
	videoID, _ := strconv.Atoi(c.Query("video_id"))
	actionType := c.Query("action_type")
	if _, exist := usersLoginInfo[token]; exist {
		c.JSON(http.StatusOK, pjdata.Response{StatusCode: 0})
	} else {
		c.JSON(http.StatusOK, pjdata.Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
	}
	database.ChangeFavoriteCount(actionType, userID)
	database.ChangeTotalFavorited(actionType, int64(videoID))
}

// FavoriteList all users have same favorite video list
func FavoriteList(c *gin.Context) {
	c.JSON(http.StatusOK, VideoListResponse{
		Response: pjdata.Response{
			StatusCode: 0,
		},
		//VideoList: DemoVideos,
	})
}
