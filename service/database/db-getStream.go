package database

import "errors"

func (db *appdbimpl) GetStream(userID int) ([]int, []int, error) {

	var photos []int
	var IDs []int
	rows, errQuery := db.c.Query(` SELECT pht.publisherID, pht.photoID FROM Photo pht, User us, Following fl
	WHERE pht.publisherID = fl.followerID
	AND fl.followingID = us.userID
	AND us.userID = ?
    ORDER BY pht.publicationDate;`, userID)
	if errQuery != nil {
		return nil, nil, errQuery
	}

	for rows.Next() {
		var photoID int
		var ID int
		errScan := rows.Scan(&ID, &photoID)
		if errScan != nil {
			return nil, nil, errScan
		}
		photos = append(photos, photoID)
		IDs = append(IDs, ID)
	}

	if rows.Err() != nil {
		return nil, nil, errors.New("An error has occurred while reading the rows of the query")
	}
	return photos, IDs, nil
}
