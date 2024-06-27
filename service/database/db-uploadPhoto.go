package database

import (
	"strconv"
	"fmt"
)

func (db *appdbimpl) UploadPhoto(userID int) (int, error) {
	query, errInsert := db.c.Exec("INSERT INTO Photo(publisherID) VALUES (?);",
		userID)
	if errInsert != nil {
		fmt.Println(errInsert)
		return -1, errInsert
	}

	ID, errID := query.LastInsertId()
	if errID != nil {
		fmt.Println(errID)
		return -1, errID
	}

	path := "./userProfile/" + strconv.Itoa(userID) + "/publishedPhotos/" + strconv.Itoa(int(ID))
	_, errUpdate := db.c.Exec(`UPDATE Photo
	SET file = ?
	WHERE photoID = ?`, path, ID)
	if errUpdate != nil {
		fmt.Println(errUpdate)
		return -1, errUpdate
	}

	return int(ID), nil
}
