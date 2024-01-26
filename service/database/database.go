package database

import (
	"database/sql"
	"errors"
)

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

	GetStream(userID int) ([]int, error)

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

	_, errEx := db.Exec("PRAGMA foreign_keys = ON;")
	if errEx != nil {
		return nil, errors.New("An error has occurred while turning the Foreign Keys ON")
	}

	var tableName string

	err := db.QueryRow("SELECT name FROM sqlite_master WHERE type = table AND name = user;").Scan((&tableName))

	if errors.Is(err, sql.ErrNoRows) {
		_, err1 := db.Exec(`CREATE TABLE IF NOT EXISTS User (
			UserID INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
			Username TEXT NOT NULL
			);`)
		if err1 != nil {
			return nil, errors.New("An error has occurred while building User table")
		}

		_, err2 := db.Exec(`CREATE TABLE IF NOT EXISTS Photo (
			file TEXT NOT NULL,
			photoID INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
			publisherID INTEGER NOT NULL,
			publicationDate TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
			FOREIGN KEY(publisherID) references User (userID)
			);`)

		if err2 != nil {
			return nil, errors.New("An error has occurred while building the Photo table")
		}

		_, err3 := db.Exec(`CREATE TABLE IF NOT EXISTS Banned (
			bannerID INTEGER NOT NULL PRIMARY KEY,
			bannedID INTEGER NOT NULL KEY,
			CHECK bannerID != bannedID,
			FOREIGN KEY(bannerID) references User(userID),
			FOREIGN KEY(bannedID) references UseR(userID)
			);`)

		if err3 != nil {
			return nil, errors.New("An error has occurred while building the Banned table")
		}

		_, err4 := db.Exec(`CREATE TABLE IF NOT EXISTS Following (
			followerID INTEGER NOT NULL PRIMARY KEY,
			followingID INTEGER NOT NULL KEY,
			CHECK followerID != followingID,
			FOREIGN KEY(followerID) references User(userID),
			FOREIGN KEY(followingID) references UseR(userID)
			);`)

		if err4 != nil {
			return nil, errors.New("An error has occurred while building the Following table")
		}
		_, err5 := db.Exec(`CREATE TABLE IF NOT EXISTS Comment (
			commentID INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
			comment TEXT NOT NULL, 
			publisherID INTEGER NOT NULL,
			photoID INTEGER NOT NULL KEY,
			FOREIGN KEY(publisherID) references User(userID),
			FOREIGN KEY(photoID) references Photo(photoID)
			);`)

		if err5 != nil {
			return nil, errors.New("An error has occurred while building the Comment table")
		}
		_, err6 := db.Exec(`CREATE TABLE IF NOT EXISTS Like (
			likedPhotoID INTEGER NOT NULL,
			likerUserID INTEGER NOT NULL,
			FOREIGN KEY(likePhotoID) references Photo(photoID)
			FOREIGN KEY(likerUserID) references User(userID)
			);`)
		if err6 != nil {
			return nil, errors.New("An error has occurred while building the Like table")
		}
	}

	return &appdbimpl{
		c: db,
	}, nil
}
