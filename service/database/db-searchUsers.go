package database

func (db *appdbimpl) SearchUsers(searcherID int, searchedName string) ([]User, error) {

	var risultato []User
	newString := "%"
	for i := 0; i < len(searchedName); i++ {
		newString += string(searchedName[i]) + "%"
	}
	rows, errQuery := db.c.Query(`SELECT username, userID FROM User
		WHERE username LIKE ? `, newString)

	if errQuery != nil {
		return risultato, errQuery
	}
	for rows.Next() {
		var user User
		errRow := rows.Scan(&user.Username, &user.UserID)
		if errRow != nil {
			return nil, errRow
		}
		risultato = append(risultato, user)
	}
	if rows.Err() != nil {
		return risultato, rows.Err()
	}
	return risultato, nil
}
