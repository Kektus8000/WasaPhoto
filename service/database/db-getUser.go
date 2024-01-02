package database

func (db *appdbimpl) GetUserByID(userID int) (User, error) {

	var utente User
	err := db.c.QueryRow(`SELECT UserID, Username FROM User WHERE UserID = ?`, userID).Scan(&utente.UserID, &utente.Username)
	return utente, err
}

func (db *appdbimpl) GetUserByUsername(username string) (User, error) {
	var utente User
	err := db.c.QueryRow(`SELECT UserID, Username FROM User WHERE Username = ?`, username).Scan(&utente.UserID, &utente.Username)
	return utente, err
}
