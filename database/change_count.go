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
	} else {
		video := FindVideo(videoID)
		video.FavoriteCount -= 1
		db.Save(&video)
	}
}
