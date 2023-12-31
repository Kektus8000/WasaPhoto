package database

func (db *appdbimpl) LikePhoto(userID int, photoID int) error {

	_, err := db.c.Exec(`INSERT INTO Like(likedPhotoID, likerUserID) VALUES(?,?);`, photoID, userID)
	return err
}
