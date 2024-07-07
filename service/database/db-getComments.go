package database

func (db *appdbimpl) GetComments(photoID int) ([]Comment, error) {

	var comments []Comment
	rows, err := db.c.Query(`SELECT cm.commentID, cm.comment, cm.publisherID, ut.username, cm.photoID 
	FROM Comment cm, User ut
	WHERE photoID = ?
	AND cm.publisherID = ut.username`, photoID)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var comm Comment
		errScan := rows.Scan(&comm.CommentID, &comm.Text, &comm.PublisherID, &comm.PublisherName, &comm.PhotoID)
		if errScan != nil {
			return nil, errScan
		}
		comments = append(comments, comm)
	}

	if rows.Err() != nil {
		return nil, rows.Err()
	}
	return comments, nil
}
