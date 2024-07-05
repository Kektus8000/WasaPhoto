package database

func (db *appdbimpl) RemoveAllComments(bannerID int, userID int) error {
	_, err := db.c.Exec(`DELETE FROM Comment
	WHERE publisherID = ? 
	AND photoID IN (
		SELECT photoID
		FROM Photo ph
		WHERE ph.PublisherID = ? 
	);`, userID, bannerID)
	if err != nil {
		return err
	}
	return nil
}
