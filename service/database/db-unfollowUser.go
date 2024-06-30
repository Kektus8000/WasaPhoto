package database

func (db *appdbimpl) UnFollowUser(followerID int, tofollowID int) error {
	_, err := db.c.Exec("DELETE FROM Following WHERE followingID = ? AND followerID = ?;", tofollowID, followerID)
	if err != nil {
		return err
	}
	return nil
}
