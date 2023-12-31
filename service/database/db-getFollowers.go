package database

func (db *appdbimpl) GetFollowers(userID int) ([]int, error) {

	var followerIDs []int
	rows, err := db.c.Query(`SELECT * FROM Following WHERE FollowingID = ?`, userID)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var ID int
		rows.Scan(&ID)
		followerIDs = append(followerIDs, ID)
	}

	if rows.Err() != nil {
		return nil, rows.Err()
	}
	return followerIDs, nil
}
