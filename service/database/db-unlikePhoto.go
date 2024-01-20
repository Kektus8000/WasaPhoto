package database

func (db *appdbimpl) UnlikePhoto(userID int, photoID int) error {
	_, err := db.c.Exec(`DELETE FROM Like WHERE likedPhotoID = ? AND likerUserID = ?;`, photoID, userID)
	if err != nil {
		return err
	}
	return nil
}
