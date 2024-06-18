package database

func (db *appdbimpl) DeletePhoto(photoID int) (string, error) {
	var filename string
	errQuery := db.c.QueryRow("SELECT file FROM Photo WHERE photoID = ?", photoID).Scan(&filename)
	if errQuery != nil {
		return "", errQuery
	}
	_, errDelete := db.c.Exec("DELETE FROM Photo WHERE photoID = ?);", photoID)
	if errDelete != nil {
		return "", errDelete
	}
	return filename, nil
}
