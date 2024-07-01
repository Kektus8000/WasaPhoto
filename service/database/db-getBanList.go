package database

func (db *appdbimpl) GetBanList(userID int) ([]int, error) {

	var bannedIDs []int
	rows, err := db.c.Query(`SELECT bannedID FROM Banned WHERE bannerID = ?`, userID)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var ID int
		errScan := rows.Scan(&ID)
		if errScan != nil {
			return nil, errScan
		}
		bannedIDs = append(bannedIDs, ID)
	}

	if rows.Err() != nil {
		return nil, rows.Err()
	}
	return bannedIDs, nil
}
