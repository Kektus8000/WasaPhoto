package database

func (db *appdbimpl) GetFollowings(userID int) ([]User, error) {

	var followings []User
	rows, err := db.c.Query(`SELECT fl.followingID, ut.username 
	FROM Following fl, User ut
	WHERE fl.followerID = ?
	AND ut.userID = fl.followingID`, userID)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var followed User
		errScan := rows.Scan(&followed.UserID, followed.Username)
		if errScan != nil {
			return nil, errScan
		}
		followings = append(followings, followed)
	}

	if rows.Err() != nil {
		return nil, rows.Err()
	}
	return followings, nil
}
