package controller

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm/utils"
	"net/http"
	"strconv"
	"sync/atomic"
	"tiktok/database"
	"tiktok/pjdata"
	"time"
)

var commentIdSequence int64 = database.AddCommentNum()

type CommentListResponse struct {
	pjdata.Response
	CommentList []pjdata.Comment `json:"comment_list,omitempty"`
}

type CommentActionResponse struct {
	pjdata.Response
	Comment pjdata.Comment `json:"comment,omitempty"`
}

// CommentAction 评论操作
func CommentAction(c *gin.Context) {
	token := c.Query("token")
	actionType := c.Query("action_type")
	videoId, _ := strconv.Atoi(c.Query("video_id"))
	if user, exist := usersLoginInfo[token]; exist {
		if actionType == "1" {
			text := c.Query("comment_text")
			atomic.AddInt64(&commentIdSequence, 1)
			now := time.Now()
			month := now.Format("01")
			day := utils.ToString(now.Format("02"))
			nowTime := month + "-" + day
			newComment := database.Comment{
				Id:         commentIdSequence,
				VideoId:    int64(videoId),
				UserId:     user.Id,
				Content:    text,
				CreateDate: nowTime,
			}
			database.SaveComment(newComment)
			c.JSON(http.StatusOK, CommentActionResponse{Response: pjdata.Response{StatusCode: 0},
				Comment: pjdata.Comment{
					Id:         commentIdSequence,
					User:       user,
					Content:    text,
					CreateDate: nowTime,
				}})
			database.ChangeCommentCount(actionType, int64(videoId))
			return
		} else if actionType == "2" {
			id, _ := strconv.Atoi(c.Query("comment_id"))
			database.DeleteComment(int64(id))
			c.JSON(http.StatusOK, CommentActionResponse{Response: pjdata.Response{StatusCode: 0}})
			database.ChangeCommentCount(actionType, int64(videoId))
		}
	} else {
		c.JSON(http.StatusOK, pjdata.Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
	}
}

// CommentList 评论列表
func CommentList(c *gin.Context) {
	token := c.Query("token")
	videoId, _ := strconv.Atoi(c.Query("video_id"))
	if _, exist := usersLoginInfo[token]; exist {
		commentList := database.GetCommentList(int64(videoId))
		c.JSON(http.StatusOK, CommentListResponse{Response: pjdata.Response{StatusCode: 0}, CommentList: commentList})
	} else {
		c.JSON(http.StatusOK, pjdata.Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
	}
	c.JSON(http.StatusOK, CommentListResponse{
		Response: pjdata.Response{StatusCode: 0},
		//CommentList: DemoComments,
	})
}
