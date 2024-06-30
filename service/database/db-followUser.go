package database

func (db *appdbimpl) FollowUser(followerID int, followingID int) error {

	_, err := db.c.Exec(`INSERT INTO Following(followerID, followingID) VALUES (?,?)`, followerID, followingID)
	if err != nil {
		return err
	}
	return nil
}
