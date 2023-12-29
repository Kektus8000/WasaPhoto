package database

func (db *appdbimpl) likePhoto(userID int, photoID int) (bool, error) {
	flag := true
	_, err := db.c.Exec(`INSERT INTO Like(likedPhotoID, likerUserID) VALUES(?,?);`, photoID, userID)

	if err != nil {
		flag = false
	}
	return flag, err
}
