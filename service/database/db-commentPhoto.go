package database

func (db *appdbimpl) CommentPhoto(commentorID int, comment string, photoID int) (int, error) {
	update, errExec := db.c.Exec(`INSERT INTO Comment(comment, publisherID, photoID) VALUES (?,?,?));`, comment, commentorID, photoID)
	if errExec != nil {
		return -1, errExec
	}

	ID, errFetch := update.LastInsertId()
	if errFetch != nil {
		return -1, errFetch
	}
	return int(ID), errExec
}
