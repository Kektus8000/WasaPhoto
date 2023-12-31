package database

func (db *appdbimpl) GetPublishedPhotos(userID int) ([]int, error) {

	var photos []int
	rows, err := db.c.Query(`SELECT photoID FROM Photo WHERE publisherID = ?`, userID)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var photo int
		rows.Scan(&photo)
		photos = append(photos, photo)
	}

	if rows.Err() != nil {
		return nil, rows.Err()
	}
	return photos, nil
}
