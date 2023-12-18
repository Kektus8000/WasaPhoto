package database

func (db *appdbimpl) uploadPhoto(filename string, userID int) (bool, error) {
	flag := true
	_, err := db.c.Exec("INSERT INTO Photo(file, publicationDate, publisherID) VALUE (?,NOW, ?));",
		filename, userID)
	if err != nil {
		flag = false
	}
	return flag, err
}
