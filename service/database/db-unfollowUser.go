package database

func (db *appdbimpl) UnFollowUser(followingID int, followerID int) error {
	_, err := db.c.Exec("DELETE FROM Following WHERE followingID = ? AND followerID = ?;", followingID, followerID)
	if err != nil {
		return err
	}
	return nil
}
