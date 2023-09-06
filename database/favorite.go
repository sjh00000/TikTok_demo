package database

func AddFavorite(videoId int64, authorId int64) {
	var dbFavorite = Favorite{
		Id:      AddFavoriteNum(),
		VideoId: videoId,
		UserId:  authorId,
	}

	db.Create(&dbFavorite)
}

func AddFavoriteNum() int64 {
	var favorite Favorite
	db.Last(&favorite)
	return favorite.Id + 1
}

func Favorites(videoId int64, userId int64, actionType string) {
	var favorite Favorite
	db.Last(&favorite, userId)
	if actionType == "1" {
		AddFavorite(videoId, userId)
	}
	if actionType == "2" {
		var result Favorite
		db.Where("video_id = ? AND user_id = ?", videoId, userId).First(&result) // 替换为你的列名和值
		db.Delete(&result)
	}
}

func IsFavorites(videoId int64, userId int64) {
	var video Video
	var favorite []Favorite
	db.Find(&favorite, "video_id = ? AND user_id = ?", videoId, userId)
	db.Find(&video, "id = ?", videoId)
	// 检查查询结果
	if len(favorite) == 0 {
		video.IsFavorite = false
	} else {
		video.IsFavorite = true
	}
	db.Save(&video)
}

func IsFavoriteUpdate(userId int64) {
	var videos []Video
	var favorite []Favorite
	db.Find(&videos)
	for _, video := range videos {
		db.Find(&favorite, "video_id = ? AND user_id = ?", video.Id, userId)
		var videoNow Video
		db.Find(&videoNow, "id = ?", video.Id)
		if len(favorite) == 0 {
			videoNow.IsFavorite = false
		} else {
			videoNow.IsFavorite = true
		}
		db.Save(&videoNow)
	}
}
