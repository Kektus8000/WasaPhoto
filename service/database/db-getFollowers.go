package database

func (db *appdbimpl) GetFollowers(userID int) ([]int, error) {

	var followerIDs []int
	rows, err := db.c.Query(`SELECT followerID FROM Following WHERE followingID = ?`, userID)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var ID int
		errScan := rows.Scan(&ID)
		if errScan != nil {
			return nil, errScan
		}
		followerIDs = append(followerIDs, ID)
	}

	if rows.Err() != nil {
		return nil, rows.Err()
	}
	return followerIDs, nil
}
