package database

func (db *appdbimpl) banUser(bannerID int, bannedID int) (bool, error) {

	flag := true
	_, err := db.c.Exec("INSERT INTO Banned(bannerID, bannedID) VALUE (?,?)", bannerID, bannedID)
	if err != nil {
		flag = false
	}
	return flag, err
}
