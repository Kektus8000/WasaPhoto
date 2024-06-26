package api

import "time"

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

type Photo struct {
	File            string
	PhotoID         int
	PublisherID     int
	PublicationDate time.Time
	Comments        []Comment
}

type UserProfile struct {
	UserID          int
	Username        string
	Followers       []int   // ID degli utenti seguaci
	Followings      []int   // ID degli utenti seguiti
	Banneds         []int   // ID degli utenti bannati
	PublishedPhotos []Photo // Foto pubblicate
}
