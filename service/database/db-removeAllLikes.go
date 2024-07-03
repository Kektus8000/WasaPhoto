package database

func (db *appdbimpl) RemoveAllLikes(bannerID int, userID int) error {
	_, err := db.c.Exec(`DELETE FROM Like lk
	WHERE cm.likerUserID = ? 
	AND cm.likedPhotoID IN (
		SELECT photoID
		FROM Photo ph
		WHERE ph.PublisherID = ? 
	);`, userID, bannerID)
	if err != nil {
		return err
	}
	return nil
}
