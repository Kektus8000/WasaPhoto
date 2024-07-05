package database

func (db *appdbimpl) RemoveAllLikes(bannerID int, userID int) error {
	_, err := db.c.Exec(`DELETE FROM Like
	WHERE likerUserID = ? 
	AND likedPhotoID IN (
		SELECT photoID
		FROM Photo ph
		WHERE ph.PublisherID = ? 
	);`, userID, bannerID)
	if err != nil {
		return err
	}
	return nil
}
