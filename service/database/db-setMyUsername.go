package database

func (db *appdbimpl) setMyUsername(userID int, newUsername string) (bool, error) {

	_, err := db.c.Query(`UPDATE table User
	SET username = ?
	WHERE userID = ?
	);`, userID, newUsername)

	if err != nil {
		return false, err
	}
	return true, err
}
