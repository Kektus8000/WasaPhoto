package database

import(	
	"fmt"
	"database/sql"
	"errors"
)

func (db *appdbimpl) setMyUsername(userID int, newUsername string) (bool,error) {
	
	_, err := db.c.Exec("UPDATE table User
	SET username = ?
	WHERE userID == ?
	);", userID, newUsername)

	if err != nil{
		return false, err
	}
	return true, err
}