package database

import (
	"fmt"
	"database/sql"
	"errors"
)

type AppDatabase interface{
	doLogin(username string) (User, error)

	getUserProfile(username string) (User, error)

	setMyUserName(oldUsername string, newUsername string) error

	followUser(followerId int) error

	unFollowUser(followerId int) error

	getMyStream() (Photo, error)

	banUser(bannedId int) error

	unBanUser(bannedId int) error

	addPhoto(photo *Photo) error

	removePhoto(photo *Photo) error

	likePhoto(photoId int) error

	unlikePhoto(photoId int) error

	commentPhoto(comment Comment, photoId int) error

	uncommentPhoto(commentId int, photoId int) error
}

type appdbimpl struct{
	c *sql.DB
}

func New(db *sql.DB) (AppDatabase, error) {
	if db == nil{
		return nil, errors.New("A database is required when building a AppDatabase")
	}

	db.Exec("PRAGMA foreign_keys = ON")

	var tableName string

	err := db.QueryRow("SELECT name FROM sqlite_master WHERE type = 'table' AND name = 'user';").Scan((&tableName))
	if errors.Is(err, sql.ErrNoRows) {
		var errors []error
		_, err1 := db.Exec("CREATE TABLE IF NOT EXISTS User (
			userID INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
			username TEXT NOT NULL
			);")
		errors = append(errors, err1)
		
		_, err2 := db.Exec("CREATE TABLE IF NOT EXISTS Photo (
			file TEXT NOT NULL,
			photoID INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
			publisherID INTEGER NOT NULL,
			publicationDate DATE NOT NULL,
			FOREIGN KEY(publisherID) references User (userID)
			);")
		errors = append(errors, err2)

		_,err3 := db.Exec("CREATE TABLE IF NOT EXISTS Banned (
			bannerID INTEGER NOT NULL PRIMARY KEY,
			bannedID INTEGER NOT NULL KEY,
			CHECK bannerID != bannedID,
			FOREIGN KEY(bannerID) references User(userID),
			FOREIGN KEY(bannedID) references UseR(userID)
			);")

		errors = append(errors, err3)

		_,err4 := db.Exec("CREATE TABLE IF NOT EXISTS Following (
			followerID INTEGER NOT NULL PRIMARY KEY,
			followingID INTEGER NOT NULL KEY,
			CHECK followerID != followingID,
			FOREIGN KEY(followerID) references User(userID),
			FOREIGN KEY(followingID) references UseR(userID)
			);")

		errors = append(errors, err4)

		_, err5 = db.Exec("CREATE TABLE IF NOT EXISTS Comment (
			commentID INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
			comment TEXT NOT NULL, 
			publisherID INTEGER NOT NULL,
			photoID INTEGER NOT NULL KEY,
			FOREIGN KEY(publisherID) references User(userID),
			FOREIGN KEY(photoID) references Photo(photoID)
			);")
		
		errors = append(errors, err5)

		_, err6 = db.Exec("CREATE TABLE IF NOT EXISTS Like (
			likedPhotoID INTEGER NOT NULL,
			likerUserID INTEGER NOT NULL,
			FOREIGN KEY(likePhotoID) references Photo(photoID)
			FOREIGN KEY(likerUserID) references User(userID)
			);")
		errors = append(errors, err6)

		for i:=1; i <= len(errors); i++ {
			if errors[i] != nil {
				return err
			}
		}
	}
	return &appdbimpl{
		c: db,
	}, nil
}