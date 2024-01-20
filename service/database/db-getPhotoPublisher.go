package database

func (db *appdbimpl) GetPhotoPublisher(photoID int) (int, error) {
	var userID int
	_, err := db.c.Query(`SELECT publisherID FROM Photo WHERE PhotoID = ?);`, photoID)

	if err != nil {
		return -1, err
	}
	return userID, err
}
