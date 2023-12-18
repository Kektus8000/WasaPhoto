package database

func (db *appdbimpl) getMyStream(userID int) ([]Photo, error) {

	var photos []Photo
	_, err := db.c.Query("SELECT * FROM Photo
	WHERE publisherID IN (SELECT followingID FROM Following WHERE followerID = ?)
	ORDER BY publicationDate DEC);", userID).Scan(&photos)

	if err != nil{
		return nil, err
	}
	return photos, nil
}
