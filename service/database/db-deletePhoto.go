package database

func (db *appdbimpl) DeletePhoto(photoID int) error {
	_, errDelete := db.c.Exec("DELETE FROM Photo WHERE photoID = ?);", photoID)
	if errDelete != nil {
		return errDelete
	}
	return nil
}
