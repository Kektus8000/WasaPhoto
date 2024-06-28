package database

import (
	"database/sql"
	"errors"
	"time"
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

type Photo struct {
	File            string
	PhotoID         int
	PublisherID     int
	PublicationDate time.Time
	Comments        []Comment
}

type AppDatabase interface {
	// Metodi GET
	GetUserByID(userID int) (User, error)
	GetUserByUsername(username string) (User, error)
	GetFollowers(userID int) ([]int, error)
	GetFollowings(userID int) ([]int, error)
	GetBanList(userID int) ([]int, error)
	GetPublishedPhotos(userID int) ([]Photo, error)
	GetPhotoPublisher(photoID int) (int, error)
	CheckBanned(bannerID int, bannedID int) (bool, error)
	GetStream(userID int) ([]Photo, error)
	GetComments(photoID int) ([]Comment, error)
	UserExists(username string) (bool, error)

	// Metodi POST
	AddUser(username string) (int, error)
	SetMyUsername(userID int, newUsername string) error
	FollowUser(followerID int, tofollowID int) error
	BanUser(bannerID int, bannedId int) error
	LikePhoto(userID int, photoID int) error

	// Metodi PUT
	SearchUsers(searcherID int, searchedName string) ([]User, error)
	UploadPhoto(userID int) (int, error)
	CommentPhoto(commentorID int, comment string, photoID int) (int, error)

	// Metodi DELETE
	UnFollowUser(followerID int, followingID int) error
	UnbanUser(bannerID int, bannedID int) error
	DeletePhoto(photoID int) (string, error)
	UnlikePhoto(userID int, photoID int) error
	UncommentPhoto(commentID int) error
}

type appdbimpl struct {
	c *sql.DB
}

func New(db *sql.DB) (AppDatabase, error) {
	if db == nil {
		return nil, errors.New("A database is required when building a AppDatabase")
	}

	_, errEx := db.Exec("PRAGMA foreign_keys = ON")
	if errEx != nil {
		return nil, errors.New("An error has occurred while turning the Foreign Keys ON")
	}

	var tableName string

	err := db.QueryRow("SELECT name FROM sqlite_master WHERE type = 'table' AND name = 'User' ").Scan(&tableName)

	if errors.Is(err, sql.ErrNoRows) {
		_, err1 := db.Exec(`CREATE TABLE IF NOT EXISTS User (
			userID INTEGER PRIMARY KEY AUTOINCREMENT,
			username TEXT NOT NULL
			);`)
		if err1 != nil {
			return nil, errors.New("An error has occurred while building User table")
		}

		_, err2 := db.Exec(`CREATE TABLE IF NOT EXISTS Photo (
			file TEXT,
			photoID INTEGER PRIMARY KEY AUTOINCREMENT,
			publisherID INTEGER NOT NULL,
			publicationDate TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			FOREIGN KEY(publisherID) references User (userID)
			);`)

		if err2 != nil {
			return nil, errors.New("An error has occurred while building the Photo table")
		}

		_, err3 := db.Exec(`CREATE TABLE IF NOT EXISTS Banned (
			bannerID INTEGER NOT NULL,
			bannedID INTEGER NOT NULL,
			PRIMARY KEY (bannerID, bannedID),
			FOREIGN KEY(bannerID) references User(userID),
			FOREIGN KEY(bannedID) references User(userID),
			CHECK (bannerID != bannedID)
			);`)

		if err3 != nil {
			return nil, errors.New("An error has occurred while building the Banned table")
		}

		_, err4 := db.Exec(`CREATE TABLE IF NOT EXISTS Following (
			followerID INTEGER NOT NULL,
			followingID INTEGER NOT NULL,
			PRIMARY KEY (followerID, followingID),
			FOREIGN KEY(followerID) references User(userID),
			FOREIGN KEY(followingID) references User(userID),
			CHECK (followerID != followingID)
			);`)

		if err4 != nil {
			return nil, errors.New("An error has occurred while building the Following table")
		}
		_, err5 := db.Exec(`CREATE TABLE IF NOT EXISTS Comment (
			commentID INTEGER PRIMARY KEY AUTOINCREMENT,
			comment TEXT NOT NULL, 
			publisherID INTEGER NOT NULL,
			photoID INTEGER NOT NULL,
			FOREIGN KEY(publisherID) references User(userID),
			FOREIGN KEY(photoID) references Photo(photoID)
			);`)

		if err5 != nil {
			return nil, errors.New("An error has occurred while building the Comment table")
		}
		_, err6 := db.Exec(`CREATE TABLE IF NOT EXISTS Like (
			likedPhotoID INTEGER NOT NULL,
			likerUserID INTEGER NOT NULL,
			PRIMARY KEY (likedPhotoID, likerUserID),
			FOREIGN KEY(likedPhotoID) references Photo(photoID),
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

// Ping checks the connection to the database.
func (db *appdbimpl) Ping() error {
	return db.c.Ping()
}
