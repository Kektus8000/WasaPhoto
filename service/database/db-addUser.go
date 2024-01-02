package database

func (db *appdbimpl) AddUser(username string) (int, error) {

	update, errQuery := db.c.Exec("INSERT INTO User(Username) VALUE (?)", username)
	if errQuery != nil {
		return -1, errQuery
	}

	ID, errFetch := update.LastInsertId()
	if errFetch != nil {
		return -1, errFetch
	}
	return int(ID), nil
}
