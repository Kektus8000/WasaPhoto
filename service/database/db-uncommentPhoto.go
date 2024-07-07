package database

func (db *appdbimpl) UncommentPhoto(commentID int) error {

	_, err := db.c.Exec("DELETE FROM Comment WHERE commentID = ?", commentID)
	if err != nil {
		return err
	}
	return nil
}
