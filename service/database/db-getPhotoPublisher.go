package database

func (db *appdbimpl) GetPhotoPublisher(photoID int) (int, error) {
	userID := 0
	err := db.c.QueryRow(`SELECT publisherID FROM Photo WHERE photoID = ?;`, photoID).Scan(&userID)
	if err != nil {
		return userID, err
	}
	return userID, err
}
