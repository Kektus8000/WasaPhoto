package api

import "time"

const PHOTOFOLDER = "/tmp/userProfile/"
const PUBLISHEDFOLDER = "/publishedPhotos/"

type User struct {
	UserID   int
	Username string
}

type Comment struct {
	CommentID   int
	Text        string
	PublisherID int
	PhotoID     int
}

type Photo struct {
	File            string
	PhotoID         int
	PublisherID     int
	PublicationDate time.Time
	Comments        []Comment
	Likes           []User
}

type UserProfile struct {
	UserID          int
	Username        string
	Followers       []User  // Utenti seguaci
	Followings      []User  // Utenti seguiti
	Banneds         []User  // Utenti bannati
	PublishedPhotos []Photo // Foto pubblicate
}
