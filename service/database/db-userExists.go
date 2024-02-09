package database

func (db *appdbimpl) UserExists(username string) (bool, error) {
	var count int
	errQuery := db.c.QueryRow("SELECT COUNT(*) FROM User WHERE Username = ?;", username).Scan(&count)
	if errQuery != nil {
		return false, errQuery
	}
	return (count == 1), nil
}
