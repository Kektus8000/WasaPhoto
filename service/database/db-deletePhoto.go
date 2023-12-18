package database

func (db *appdbimpl) deletePhoto(photoID int) (bool, error) {
	flag := true
	_, err := db.c.Exec("DELETE FROM Photo WHERE PhotoID = ?);", photoID)
	if err != nil {
		flag = false
	}
	return flag, err
}
