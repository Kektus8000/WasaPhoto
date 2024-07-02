package database

func (db *appdbimpl) GetFollowers(userID int) ([]User, error) {

	var followers []User
	rows, err := db.c.Query(`SELECT fl.followerID, ut.username 
	FROM Following fl, User ut 
	WHERE fl.followingID = ?
	AND ut.userID = fl.followerID`, userID)

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var utente User
		errScan := rows.Scan(&utente.UserID, &utente.Username)
		if errScan != nil {
			return nil, errScan
		}
		followers = append(followers, utente)
	}

	if rows.Err() != nil {
		return nil, rows.Err()
	}

	return followers, nil
}
