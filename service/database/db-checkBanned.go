package database

func (db *appdbimpl) checkBanned(bannerID int, bannedID int) (bool, error) {

	flag := true
	count, err := db.c.Query("COUNT (*) FROM Banned WHERE bannerID = ? AND bannedID = ?);", bannerID, bannedID)

	if err != nil || count > 0 {
		flag = false
	}
	return flag, err
}
