package database

func (db *appdbimpl) unbanUser(bannerID int, bannedID int) (bool, error) {

	flag := true
	_, err := db.c.Exec("DELETE FROM Banned WHERE bannerID = ? AND bannedID = ?", bannerID, bannedID)
	if err != nil {
		flag = false
	}
	return flag, err
}
