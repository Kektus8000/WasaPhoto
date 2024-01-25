package api

type User struct {
	UserID   int
	Username string
}

type Comment struct {
	CommentID   int
	Comment     string
	PublisherID int
	PhotoID     int
}
