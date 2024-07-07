package database

func (db *appdbimpl) AlreadyInUse(userID int, username string) (bool, error) {

	var count int
	err := db.c.QueryRow("SELECT COUNT (*) FROM User WHERE userID != ? AND username = ?;", userID, username).Scan(&count)
	if err != nil {
		return false, err
	}
	return (count >= 1), nil
}
