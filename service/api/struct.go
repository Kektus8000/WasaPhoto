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
	Followers       []User   // Utenti seguaci
	Followings      []User   // Utenti seguiti
	Banneds         []User   // Utenti bannati
	PublishedPhotos []Photo  // Foto pubblicate
}
