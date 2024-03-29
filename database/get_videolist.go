package database

import (
	"tiktok/pjdata"
)

// SearchVideo VideoList 读取 video 表中所有数据
func SearchVideo() []pjdata.Video {

	//数据库里的数据结构"database"
	var videos []Video

	//常用的数据结构“common”
	var videoList []pjdata.Video

	//查找视频信息
	db.Find(&videos)

	// 合并用户信息以及视频信息
	for i := len(videos) - 1; i >= 0; i-- {

		videoNow := videos[i]
		//定义common数据结构
		var video pjdata.Video

		author := SearchAuthor(videoNow.AuthorId)

		video.Author = pjdata.Author(author)
		GetVideo(&video, videoNow)
		videoList = append(videoList, video)
	}

	return videoList
}

// SearchAuthor 通过id寻找authors表
func SearchAuthor(authorId int64) Author {
	var author Author
	db.Last(&author, authorId)

	return author
}
func FindVideo(videoId int64) Video {
	var video Video
	db.Last(&video, videoId)

	return video
}

// GetVideo 合并
func GetVideo(video *pjdata.Video, videoNow Video) {
	video.Id = videoNow.Id
	video.PlayUrl = videoNow.PlayUrl
	video.CoverUrl = videoNow.CoverUrl
	video.FavoriteCount = videoNow.FavoriteCount
	video.CommentCount = videoNow.CommentCount
	video.IsFavorite = videoNow.IsFavorite
	video.Title = videoNow.Title
}

func SearchAuthorVideo(userId int64, user pjdata.Author) []pjdata.Video {
	var videos []Video

	var pjVideo []pjdata.Video
	db.Where("author_id = ?", userId).Find(&videos)
	for _, video := range videos {
		pjVideo = append(pjVideo, pjdata.Video{
			Id:            video.Id,
			Author:        user,
			PlayUrl:       video.PlayUrl,
			CoverUrl:      video.CoverUrl,
			FavoriteCount: video.FavoriteCount,
			CommentCount:  video.CommentCount,
			IsFavorite:    video.IsFavorite,
			Title:         video.Title,
		})
	}
	return pjVideo
}
