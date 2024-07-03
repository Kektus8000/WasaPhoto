package database

func (db *appdbimpl) RemoveAllComments(bannerID int, userID int) error {
	_, err := db.c.Exec(`DELETE FROM Comment cm
	WHERE cm.publisherID = ? 
	AND cm.photoID IN (
		SELECT photoID
		FROM Photo ph
		WHERE ph.PublisherID = ? 
	);`, userID, bannerID)
	if err != nil {
		return err
	}
	return nil
}
