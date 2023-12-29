package database

func (db *appdbimpl) getPublishedPhotos(userID int) ([]Photo, error) {

	var photos []Photo
	var errors error
	errors = nil
	_, err := db.c.Query(`SELECT * FROM Photo WHERE publisherID = ?);`, userID).Scan(&photos)

	errors = err
	return photos, errors