package database

func (db *appdbimpl) uncommentPhoto(commentID int) error {

	_, err := db.c.Exec("DELETE FROM table Comment WHERE commentID == ?", commentID)
	if err != nil {
		return err
	}
	return nil
}
