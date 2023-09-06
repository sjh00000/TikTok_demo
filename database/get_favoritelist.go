package database

import (
	"tiktok/pjdata"
)

var pjVideo pjdata.Video

func FavoriteList(userID int64) []pjdata.Video {
	var favorites []Favorite
	var videoListId []int64
	db.Where("user_id = ?", userID).Find(&favorites)
	for i := 0; i <= len(favorites)-1; i++ {
		favoriteNow := favorites[i]
		videoListId = append(videoListId, favoriteNow.VideoId)

	}
	var videoList []pjdata.Video

	for i := 0; i <= len(videoListId)-1; i++ {
		var video Video
		db.First(&video, videoListId[i])
		change(video)
		videoList = append(videoList, pjVideo)
	}

	return videoList
}

func change(video Video) {
	pjVideo = pjdata.Video{
		Id:            video.Id,
		Author:        pjdata.Author(SearchAuthor(video.AuthorId)),
		PlayUrl:       video.PlayUrl,
		CoverUrl:      video.CoverUrl,
		FavoriteCount: video.FavoriteCount,
		CommentCount:  video.CommentCount,
		IsFavorite:    video.IsFavorite,
		Title:         video.Title,
	}
}
