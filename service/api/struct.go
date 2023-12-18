package api

type User struct {
	UserID   int
	Username string
}

type Banned struct {
	BannerID int
	BannedID int
}

type Follower struct {
	FollowerID  int
	FollowingID int
}

type Photo struct {
	File            string
	PhotoID         int
	Likes           []Like
	PublicationDate string
	PublisherID     int
}

type Comment struct {
	CommentID   int
	Comment     string
	PublisherID int
	PhotoID     int
}

type Like struct {
	LikerID      int
	LikedPhotoID int
}
