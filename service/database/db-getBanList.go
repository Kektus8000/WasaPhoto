package database

func (db *appdbimpl) GetBanList(userID int) ([]User, error) {

	var banneds []User
	rows, err := db.c.Query(`SELECT bd.bannedID, ut.username 
	FROM Banned bd, User ut 
	WHERE bd.bannerID = ?
	AND bd.banned = ut.userID`, userID)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var ban User
		errScan := rows.Scan(&ban.UserID, &ban.Username)
		if errScan != nil {
			return nil, errScan
		}
		banneds = append(banneds, ban)
	}

	if rows.Err() != nil {
		return nil, rows.Err()
	}
	return banneds, nil
}
