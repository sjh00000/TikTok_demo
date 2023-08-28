package database

import (
	"tiktok/pjdata"
)

func AddVideo(video pjdata.Video) {
	var dbVideo = Video{
		Id:            video.Id,
		AuthorId:      video.Author.Id,
		PlayUrl:       video.PlayUrl,
		FavoriteCount: video.FavoriteCount,
		CommentCount:  video.CommentCount,
		IsFavorite:    video.IsFavorite,
		Title:         video.Title,
	}

	db.Create(&dbVideo)
}

func AddVideoNum() int64 {
	var video Video
	db.Last(&video)
	return video.Id + 1
}
