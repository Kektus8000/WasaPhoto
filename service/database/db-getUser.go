package database

func (db *appdbimpl) GetUser(userID int) (User, error) {

	var utente User
	err := db.c.QueryRow(`SELECT UserID, Username FROM User WHERE UserID = ?`, userID).Scan(&utente.UserID, &utente.Username)
	return utente, err
}
