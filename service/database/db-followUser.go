package database

func (db *appdbimpl) FollowUser(followerID int, followingID int) error {

	_, err := db.c.Exec(`INSERT INTO Follower(followerID, followingID) VALUE (?,?)`, followerID, followingID)
	return err
}
