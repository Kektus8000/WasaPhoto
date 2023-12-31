package database

func (db *appdbimpl) GetMyStream(userID int) error {

	_, err := db.c.Query(`SELECT * FROM Photo
	WHERE publisherID IN (SELECT followingID FROM Following WHERE followerID = ?)
	ORDER BY publicationDate DEC);`, userID)

	return err
}
