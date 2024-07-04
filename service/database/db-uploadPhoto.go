package database

import (
	"strconv"
)

func (db *appdbimpl) UploadPhoto(userID int) (int, error) {
	query, errInsert := db.c.Exec("INSERT INTO Photo(publisherID) VALUES (?);",
		userID)
	if errInsert != nil {
		return -1, errInsert
	}

	ID, errID := query.LastInsertId()
	if errID != nil {
		return -1, errID
	}

	path := PHOTOFOLDER + strconv.Itoa(userID) + PUBLISHEDFOLDER + strconv.Itoa(int(ID))
	_, errUpdate := db.c.Exec(`UPDATE Photo
	SET file = ?
	WHERE photoID = ?`, path, ID)
	if errUpdate != nil {
		return -1, errUpdate
	}

	return int(ID), nil
}
