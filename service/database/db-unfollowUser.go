package database

func (db *appdbimpl) UnFollowUser(followingID int, followerID int) error {
	_, err := db.c.Exec("DELETE FROM Follower WHERE followingID = ? AND followerID = ?);", followingID, followerID)
	return err
}
