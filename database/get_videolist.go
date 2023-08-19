package database

import (
	"fmt"
	"strconv"
	"tiktok/pjdata"
)

// SearchVideo VideoList 读取 video 表中所有数据
func SearchVideo() []pjdata.Video {

	var videos []Video
	var videoList []pjdata.Video

	//查找视频信息
	db.Find(&videos)

	// 合并用户信息以及视频信息
	for _, videoNow := range videos {
		fmt.Print("id是" + strconv.FormatInt(videoNow.AuthorId, 10) + "\n")
		fmt.Print(videoNow)
		var video pjdata.Video
		author := SearchAuthor(videoNow.AuthorId)
		video.Author = pjdata.Author(author)
		GetVideo(&video, videoNow)
		videoList = append(videoList, video)
	}

	return videoList
}

func SearchAuthor(authorId int64) Author {
	var author Author
	db.First(&author, authorId)
	return author
}

func GetVideo(video *pjdata.Video, videoNow Video) {
	video.Id = videoNow.Id
	video.PlayUrl = videoNow.PlayUrl
	video.CoverUrl = videoNow.CoverUrl
	video.FavoriteCount = videoNow.FavoriteCount
	video.CommentCount = videoNow.CommentCount
	video.IsFavorite = videoNow.IsFavorite
	video.Title = videoNow.Title
}
