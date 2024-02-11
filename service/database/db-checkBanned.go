package database

func (db *appdbimpl) CheckBanned(bannerID int, bannedID int) (bool, error) {

	var count int
	err := db.c.QueryRow("SELECT COUNT (*) FROM Banned WHERE bannerID = ? AND bannedID = ?;", bannerID, bannedID).Scan(&count)
	if err != nil {
		return false, err
	}
	return (count == 1), nil
}
