package database

func (db *appdbimpl) UploadPhoto(filename string, userID int) (int, error) {
	query, errInsert := db.c.Exec("INSERT INTO Photo(file, publicationDate, publisherID) VALUE (?,NOW,?));",
		filename, userID)
	if errInsert != nil {
		return -1, errInsert
	}

	ID, errID := query.LastInsertId()
	if errID != nil {
		return -1, errID
	}
	return int(ID), nil
}
