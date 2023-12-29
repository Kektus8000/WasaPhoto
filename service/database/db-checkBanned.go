package database

func (db *appdbimpl) checkBanned(bannerID int, bannedID int) (bool, error) {

	flag := false
	banned, err := db.c.Query("COUNT (*) FROM Banned WHERE bannerID = ? AND bannedID = ?);", bannerID, bannedID)
	if banned != nil {
		flag = true
	}
	return flag, err
}
