package database

func (db *appdbimpl) GetComment(commentID int) (Comment, error) {

	var comment Comment
	err := db.c.QueryRow(`SELECT *
	FROM Comment 
	WHERE CommentID = ?`, commentID).Scan(&comment.CommentID, &comment.Text, &comment.PublisherID, &comment.PhotoID)
	if err != nil {
		return comment, err
	}
	return comment, nil
}
