package database

func (db *appdbimpl) GetLikes(photoID int) ([]User, error) {

	var likes []User
	rows, err := db.c.Query(`SELECT lk.likerUserID, ut.username 
	FROM Like lk, User ut 
	WHERE lk.likedPhotoID = ?
	AND lk.likerUserID = ut.userID`, photoID)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var lover User
		errScan := rows.Scan(&lover.UserID, &lover.Username)
		if errScan != nil {
			return nil, errScan
		}
		likes = append(likes, lover)
	}

	if rows.Err() != nil {
		return nil, rows.Err()
	}

	return likes, nil
}
