package database

func (db *appdbimpl) SetMyUsername(userID int, newUsername string) error {

	_, err := db.c.Exec(`UPDATE User
	SET username = ?
	WHERE userID = ?
	;`, newUsername, userID)
	if err != nil {
		return err
	}
	return nil
}
