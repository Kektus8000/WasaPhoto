package database

func (db *appdbimpl) GetStream(userID int) ([]Photo, error) {

	var photos []Photo
	rows, errQuery := db.c.Query(`SELECT pht.file, pht.photoID, pht.publisherID, pht.publicationDate
	FROM Photo pht, User us, Following fl
	WHERE pht.publisherID = fl.followingID
	AND fl.followerID = us.userID
	AND us.userID = ?
    ORDER BY pht.publicationDate;`, userID)
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
