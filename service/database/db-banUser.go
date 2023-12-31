package database

func (db *appdbimpl) BanUser(bannerID int, bannedID int) error {

	_, err := db.c.Exec("INSERT INTO Banned(bannerID, bannedID) VALUE (?,?)", bannerID, bannedID)
	return err
}
