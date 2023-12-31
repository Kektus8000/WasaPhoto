package database

func (db *appdbimpl) DeletePhoto(photoID int) (string, error) {
	var filename string
	errQuery := db.c.QueryRow("SELECT file FROM Photo WHERE PhotoID = ?", photoID).Scan(&filename)
	if errQuery != nil {
		return "", errQuery
	}
	_, errDelete := db.c.Exec("DELETE FROM Photo WHERE PhotoID = ?);", photoID)
	return filename, errDelete
}
