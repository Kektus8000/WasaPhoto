package database

func (db *appdbimpl) GetPublishedPhotos(userID int) ([]int, error) {

	var photos []int
	rows, err := db.c.Query(`SELECT photoID FROM Photo WHERE publisherID = ?`, userID)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var photo int
		errScan := rows.Scan(&photo)
		if errScan != nil {
			return nil, errScan
		}
		photos = append(photos, photo)
	}

	if rows.Err() != nil {
		return nil, rows.Err()
	}
	return photos, nil
}
