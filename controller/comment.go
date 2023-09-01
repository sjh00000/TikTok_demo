package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"tiktok/pjdata"
)

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

	if user, exist := usersLoginInfo[token]; exist {
		if actionType == "1" {
			text := c.Query("comment_text")
			c.JSON(http.StatusOK, CommentActionResponse{Response: pjdata.Response{StatusCode: 0},
				Comment: pjdata.Comment{
					Id:         1,
					User:       user,
					Content:    text,
					CreateDate: "05-01",
				}})
			return
		}
		c.JSON(http.StatusOK, pjdata.Response{StatusCode: 0})
	} else {
		c.JSON(http.StatusOK, pjdata.Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
	}
}

// CommentList 评论列表
func CommentList(c *gin.Context) {
	c.JSON(http.StatusOK, CommentListResponse{
		Response:    pjdata.Response{StatusCode: 0},
		CommentList: DemoComments,
	})
}
