package database

func (db *appdbimpl) GetPhotoPublisher(photoID int) (int, error) {
	var userID int
	err := db.c.QueryRow(`SELECT publisherID FROM Photo WHERE PhotoID = ?;`, photoID).Scan(userID)
	if err != nil {
		return -1, err
	}
	return userID, err
}
