package database

func (db *appdbimpl) UnbanUser(bannerID int, bannedID int) error {
	_, err := db.c.Exec("DELETE FROM Banned WHERE bannerID = ? AND bannedID = ?", bannerID, bannedID)
	if err != nil {
		return err
	}
	return nil
}
