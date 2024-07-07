package database

func (db *appdbimpl) CommentPhoto(commentorID int, photoID int, comment string) error {
	_, errExec := db.c.Exec(`INSERT INTO Comment(comment, publisherID, photoID) VALUES (?,?,?);`, comment, commentorID, photoID)
	if errExec != nil {
		return errExec
	}
	return nil
}
