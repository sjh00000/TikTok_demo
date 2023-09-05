package database

func ChangeFavoriteCount(actionType string, userID int64) {
	if actionType == "1" {
		user := SearchAuthor(userID)
		user.FavoriteCount += 1
		db.Save(&user)
	} else {
		user := SearchAuthor(userID)
		user.FavoriteCount -= 1
		db.Save(&user)
	}
}
func ChangeTotalFavorited(actionType string, videoID int64) {
	if actionType == "1" {
		video := FindVideo(videoID)
		video.FavoriteCount += 1
		db.Save(&video)
		var author Author
		db.First(&author, video.AuthorId)
		author.TotalFavorite += 1
		db.Save(&author)
	} else {
		video := FindVideo(videoID)
		video.FavoriteCount -= 1
		db.Save(&video)
		var author Author
		db.First(&author, video.AuthorId)
		author.TotalFavorite -= 1
		db.Save(&author)
	}
}

func ChangeCommentCount(actionType string, id int64) {
	if actionType == "1" {
		video := FindVideo(id)
		video.CommentCount += 1
		db.Save(&video)
	} else {
		video := FindVideo(id)
		video.CommentCount -= 1
		db.Save(&video)

	}
}
