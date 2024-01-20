package database

func (db *appdbimpl) GetFollowings(userID int) ([]int, error) {

	var followingIDs []int
	rows, err := db.c.Query(`SELECT * FROM Following WHERE FollowerID = ?`, userID)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var ID int
		errScan := rows.Scan(&ID)
		if errScan != nil {
			return nil, errScan
		}
		followingIDs = append(followingIDs, ID)
	}

	if rows.Err() != nil {
		return nil, rows.Err()
	}
	return followingIDs, nil
}
