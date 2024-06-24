package database

func (db *appdbimpl) SearchUsers(searcherID int, searchedName string) ([]string, error) {

	var risultato []string
	newString := "%"
	for i := 0; i < len(searchedName); i++ {
		newString += string(searchedName[i]) + "%"
	}

	rows, errQuery := db.c.Query(`SELECT ut.username FROM User ut, Banned bd
		WHERE ut.username LIKE ? 
		AND NOT (bd.bannedID == ?
			AND bd.bannerID == ut.userID
			AND ut.userID == ?)`, newString, searcherID, searcherID)

	if errQuery != nil {
		return risultato, errQuery
	}
	for rows.Next() {
		var username string
		errRow := rows.Scan(&username)
		if errRow != nil {
			return nil, errRow
		}
		risultato = append(risultato, username)
	}
	if rows.Err() != nil {
		return risultato, rows.Err()
	}
	return risultato, nil
}
