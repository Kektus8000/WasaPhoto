package database

func (db *appdbimpl) CommentPhoto(commentorID int, photoID int, comment string) (int, error) {
	query, errExec := db.c.Exec(`INSERT INTO Comment(comment, publisherID, photoID) VALUES (?,?,?);`, comment, commentorID, photoID)
	if errExec != nil {
		return -1, errExec
	}
	ID, errID := query.LastInsertId()
	if errID != nil{
		return -1, errID
	}
	return int(ID), nil
}
