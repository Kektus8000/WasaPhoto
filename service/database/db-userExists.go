package database

func (db *appdbimpl) UserExists(username string) (bool, error) {
	var count int
	errQuery := db.c.QueryRow("SELECT COUNT(*) FROM User WHERE Username = ?);", username).Scan(&count)
	return (count == 1), errQuery
}
