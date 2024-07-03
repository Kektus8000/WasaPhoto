package database

func (db *appdbimpl) GetPublishedPhotos(userID int) ([]Photo, error) {

	var photos []Photo
	rows, errQuery := db.c.Query(`SELECT file, photoID, publisherID, publicationDate
	FROM Photo WHERE publisherID = ?
	ORDER BY publicationDate DESC`, userID)
	if errQuery != nil {
		return nil, errQuery
	}

	for rows.Next() {
		var photo Photo
		errScan := rows.Scan(&photo.File, &photo.PhotoID, &photo.PublisherID, &photo.PublicationDate)
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
