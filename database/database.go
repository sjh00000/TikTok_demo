package database

type Video struct {
	Id            int64  `gorm:"column:id"`
	AuthorId      int64  `gorm:"column:author_id"`
	PlayUrl       string `gorm:"column:play_url"`
	CoverUrl      string `gorm:"column:cover_url"`
	FavoriteCount int64  `gorm:"column:favorite_count"`
	CommentCount  int64  `gorm:"column:comment_count"`
	IsFavorite    bool   `gorm:"column:is_favorite"`
	Title         string `gorm:"column:title"`
}

type Comment struct {
	Id         int64  `gorm:"column:"`
	User       Author `gorm:"column:"`
	Content    string `gorm:"column:"`
	CreateDate string `gorm:"column:"`
}

type Author struct {
	Id              int64  `gorm:"column:id"`
	Name            string `gorm:"column:name"`
	FollowCount     int64  `gorm:"column:follow_count"`
	FollowerCount   int64  `gorm:"column:follower_count"`
	IsFollow        bool   `gorm:"column:is_follow"`
	Avatar          string `gorm:"column:avatar"`
	BackgroundImage string `gorm:"column:background_image"`
	Signature       string `gorm:"column:signature"`
	TotalFavorite   int64  `gorm:"column:total_favorited"`
	WorkCount       int64  `gorm:"column:work_count"`
	FavoriteCount   int64  `gorm:"column:favorite_count"`
	Token           string `gorm:"column:token"`
}

type Message struct {
	Id         int64  `gorm:"column:"`
	Content    string `gorm:"column:"`
	CreateTime string `gorm:"column:"`
}

type MessageSendEvent struct {
	UserId     int64  `gorm:"column:"`
	ToUserId   int64  `gorm:"column:"`
	MsgContent string `gorm:"column:"`
}

type MessagePushEvent struct {
	FromUserId int64  `json:"user_id,omitempty"`
	MsgContent string `json:"msg_content,omitempty"`
}
