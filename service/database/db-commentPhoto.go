package database

func (db *appdbimpl) commentPhoto(photoID int, comment string, publisherID int) (int, error) {
	query, errExec := db.c.Exec("INSERT INTO Comment(comment, publisherID, photoID) VALUES (?,?,?));", comment, publisherID, photoID)
	if errExec != nil {
		return -1, errExec
	}
	commentID, errQuery := query.LastInsertId()

	if errQuery != nil {
		return -1, errQuery
	}
	return commentID, nil

}
