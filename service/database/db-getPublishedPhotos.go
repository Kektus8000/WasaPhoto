package database

func (db *appdbimpl) getPublishedPhotos(userID int) ([]Photo, error) {

	var photos []Photo
	_, err := db.c.Query("SELECT * FROM Photo WHERE publisherID = ?);", userID).Scan(&photos)

	if err != nil{
		return nil, err
	}
	return photos, nil