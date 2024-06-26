package database

func (db *appdbimpl) GetComments(photoID int) ([]Comment, error) {

	var comments []Comment
	rows, err := db.c.Query(`SELECT * FROM Comment WHERE photoID = ?`, photoID)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var comm Comment
		errScan := rows.Scan(&comm)
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
