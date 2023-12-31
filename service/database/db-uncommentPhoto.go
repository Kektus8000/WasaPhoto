package database

func (db *appdbimpl) UncommentPhoto(commentID int) error {

	_, err := db.c.Exec("DELETE FROM table Comment WHERE commentID == ?", commentID)
	return err
}
