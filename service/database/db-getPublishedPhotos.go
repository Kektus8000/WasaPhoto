package database

import "errors"

func (db *appdbimpl) GetPublishedPhotos(userID int) ([]int, error) {

	var photos []int
	rows, errQuery := db.c.Query(`SELECT photoID FROM Photo WHERE publisherID = ?
	ORDER BY publicationDate`, userID)
	if errQuery != nil {
		return nil, errQuery
	}

	for rows.Next() {
		var photoID int
		errScan := rows.Scan(&photoID)
		if errScan != nil {
			return nil, errScan
		}
		photos = append(photos, photoID)
	}

	if rows.Err() != nil {
		return nil, errors.New("An error has occurred while reading the rows of the query")
	}
	return photos, nil
}
