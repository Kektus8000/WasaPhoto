package database

func (db *appdbimpl) CheckBanned(bannerID int, bannedID int) (bool, error) {

	flag := false
	banned, err := db.c.Query("COUNT (*) FROM Banned WHERE bannerID = ? AND bannedID = ?);", bannerID, bannedID)
	if banned != nil {
		flag = true
	} else if err != nil {
		return flag, err
	}
	return flag, nil
}
