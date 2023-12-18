package database

func (db *appdbimpl) unFollowUser(followingID int, followerID int) (bool, error) {

	flag := true
	_, err := db.c.Exec("DELETE FROM Follower WHERE followingID = ? AND followerID = ?);", followingID, followerID)
	if err != nil {
		flag = false
	}
	return flag, err
}
