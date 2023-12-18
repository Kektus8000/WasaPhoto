package database

func (db *appdbimpl) getUserProfile(userID int) (User, error) {

	var wantedUser User
	err := db.c.QueryRow("SELECT * FROM User WHERE userID = ?(;", userID).Scan(&wantedUser)
	return wantedUser, err
}
