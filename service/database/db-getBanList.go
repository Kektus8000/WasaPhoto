package database

func (db *appdbimpl) GetBanList(userID int) ([]int, error) {

	var bannedIDs []int
	rows, err := db.c.Query(`SELECT * FROM Banned WHERE bannerID = ?`, userID)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var ID int
		rows.Scan(&ID)
		bannedIDs = append(bannedIDs, ID)
	}

	if rows.Err() != nil {
		return nil, rows.Err()
	}
	return bannedIDs, nil
}
