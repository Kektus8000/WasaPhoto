package database

func (db *appdbimpl) unlikePhoto(userID int, photoID int) (bool, error) {
	flag := true
	_, err := db.c.Exec("DELETE FROM Like WHERE likedPhotoID = ? AND likerUserID = ?;", photoID, userID)

	if err != nil {
		flag = false
	}
	return flag, err
}
