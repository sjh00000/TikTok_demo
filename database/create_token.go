package database

import "tiktok/pjdata"

func CreateToken(user pjdata.Author) {
	var dbAuthor = Author{
		Id:            user.Id,
		Name:          user.Name,
		FollowCount:   0,
		FollowerCount: 0,
		IsFollow:      false,
		TotalFavorite: 0,
		WorkCount:     0,
		FavoriteCount: 0,
		Token:         user.Token,
	}

	db.Create(&dbAuthor)
}
