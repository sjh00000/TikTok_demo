package database

import "tiktok/pjdata"

func AddCommentNum() int64 {
	var comment Comment
	db.Last(&comment)
	return comment.Id
}

func SaveComment(comment Comment) {
	db.Create(&comment)
}

func DeleteComment(id int64) {
	var comment Comment
	db.Delete(&comment, id)
}

func GetCommentList(id int64) []pjdata.Comment {
	var comments []Comment
	var commentList []pjdata.Comment
	db.Where("video_id = ?", id).Find(&comments)
	for i := len(comments) - 1; i >= 0; i-- {
		commentNow := comments[i]
		author := SearchAuthor(commentNow.UserId)
		comment := pjdata.Comment{
			Id:         commentNow.Id,
			User:       pjdata.Author(author),
			Content:    commentNow.Content,
			CreateDate: commentNow.CreateDate,
		}
		commentList = append(commentList, comment)
	}
	return commentList
}
