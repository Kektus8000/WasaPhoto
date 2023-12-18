package database

func (db *appdbimpl) followUser(followerID int, followingID int) (bool, error) {

	flag := true
	_, err := db.c.Exec("INSERT INTO Follower(followerID, followingID) VALUE (?,?)", followerID, followingID)
	if err != nil {
		flag = false
	}
	return flag, err
}
