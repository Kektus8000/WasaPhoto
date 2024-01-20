package database

import (
	"database/sql"
	"errors"
)

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

type AppDatabase interface {
	AddUser(username string) (int, error)

	GetUserByID(userID int) (User, error)

	GetUserByUsername(username string) (User, error)

	GetFollowers(userID int) ([]int, error)

	GetFollowings(userID int) ([]int, error)

	GetBanList(userID int) ([]int, error)

	CheckBanned(bannerID int, bannedID int) (bool, error)

	SetMyUsername(userID int, newUsername string) error

	FollowUser(followerID int, tofollowID int) error

	UnFollowUser(followerID int, followingID int) error

	GetPublishedPhotos(userID int) ([]int, error)

	GetPhotoPublisher(photoID int) (int, error)

	BanUser(bannerID int, bannedId int) error

	UnbanUser(bannerID int, bannedID int) error

	UploadPhoto(filename string, userID int) (int, error)

	DeletePhoto(photoID int) (string, error)

	LikePhoto(userID int, photoID int) error

	UnlikePhoto(userID int, photoID int) error

	CommentPhoto(commentorID int, comment string, photoID int) (int, error)

	UncommentPhoto(commentID int) error

	UserExists(username string) (bool, error)
}

type appdbimpl struct {
	c *sql.DB
}

func New(db *sql.DB) (AppDatabase, error) {
	if db == nil {
		return nil, errors.New("A database is required when building a AppDatabase")
	}

	db.Exec("PRAGMA foreign_keys = ON")

	var tableName string

	err := db.QueryRow("SELECT name FROM sqlite_master WHERE type = table AND name = user;").Scan((&tableName))
	var errAny error = nil
	if errors.Is(err, sql.ErrNoRows) {
		var errors []error
		_, err1 := db.Exec(`CREATE TABLE IF NOT EXISTS User (
			UserID INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
			Username TEXT NOT NULL
			);`)
		errors = append(errors, err1)

		_, err2 := db.Exec(`CREATE TABLE IF NOT EXISTS Photo (
			file TEXT NOT NULL,
			photoID INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
			publisherID INTEGER NOT NULL,
			publicationDate TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
			FOREIGN KEY(publisherID) references User (userID)
			);`)
		errors = append(errors, err2)

		_, err3 := db.Exec(`CREATE TABLE IF NOT EXISTS Banned (
			bannerID INTEGER NOT NULL PRIMARY KEY,
			bannedID INTEGER NOT NULL KEY,
			CHECK bannerID != bannedID,
			FOREIGN KEY(bannerID) references User(userID),
			FOREIGN KEY(bannedID) references UseR(userID)
			);`)

		errors = append(errors, err3)

		_, err4 := db.Exec(`CREATE TABLE IF NOT EXISTS Following (
			followerID INTEGER NOT NULL PRIMARY KEY,
			followingID INTEGER NOT NULL KEY,
			CHECK followerID != followingID,
			FOREIGN KEY(followerID) references User(userID),
			FOREIGN KEY(followingID) references UseR(userID)
			);`)

		errors = append(errors, err4)

		_, err5 := db.Exec(`CREATE TABLE IF NOT EXISTS Comment (
			commentID INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
			comment TEXT NOT NULL, 
			publisherID INTEGER NOT NULL,
			photoID INTEGER NOT NULL KEY,
			FOREIGN KEY(publisherID) references User(userID),
			FOREIGN KEY(photoID) references Photo(photoID)
			);`)

		errors = append(errors, err5)

		_, err6 := db.Exec(`CREATE TABLE IF NOT EXISTS Like (
			likedPhotoID INTEGER NOT NULL,
			likerUserID INTEGER NOT NULL,
			FOREIGN KEY(likePhotoID) references Photo(photoID)
			FOREIGN KEY(likerUserID) references User(userID)
			);`)
		errors = append(errors, err6)

		for i := 1; i <= len(errors); i++ {
			if errors[i] != nil {
				errAny = errors[i]
				break
			}
		}
	}
	return &appdbimpl{
		c: db,
	}, errAny
}
