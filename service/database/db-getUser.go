package database

func (db *appdbimpl) GetUserByID(userID int) (User, error) {
	var utente User
	err := db.c.QueryRow(`SELECT userID, username FROM User WHERE userID = ?`, userID).Scan(&utente.UserID, &utente.Username)
	if err != nil {
		return utente, err
	}
	return utente, nil
}

func (db *appdbimpl) GetUserByUsername(username string) (User, error) {
	var utente User
	err := db.c.QueryRow(`SELECT userID, username FROM User WHERE username = ?`, username).Scan(&utente.UserID, &utente.Username)
	if err != nil {
		return utente, err
	}
	return utente, nil
}
