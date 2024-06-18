package api

import (
	"encoding/json"
	"io/fs"
	"net/http"
	"os"
	"strconv"

	"github.com/Kektus8000/WasaPhoto/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) DoLogin(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	w.Header().Set("content-type", "application/json")

	var user User
	errDecode := json.NewDecoder(r.Body).Decode(&user)
	if errDecode != nil {
		http.Error(w, "An error has occurred while decoding the Request Body", 500)
		return
	}

	var username string = user.Username
	encoder := json.NewEncoder(w)
	exist, errQuery := rt.db.UserExists(username)
	if errQuery != nil {
		http.Error(w, "An error has occurred during a query in the database", 500)
		return
	}

	if !exist {
		ID, errQuery := rt.db.AddUser(username)
		if ID == -1 || errQuery != nil {
			http.Error(w, "An error has occurred during a query in the database", 500)
			return
		}

		errDir := os.MkdirAll("/main/user/", fs.FileMode(os.O_RDWR))
		if errDir != nil {
			http.Error(w, "An error has occurred while adding the user to the user list", 500)
			return
		}

		errOS := os.WriteFile("/main/user/"+username+strconv.Itoa(ID), nil, fs.FileMode(os.O_RDWR))
		if errOS != nil {
			http.Error(w, "An error has occurred while adding the user to the user list", 500)
			return
		}
		user.UserID = ID
		w.WriteHeader(http.StatusCreated)
	} else if exist {
		found, errQuery := rt.db.GetUserByUsername(username)
		if errQuery != nil {
			http.Error(w, "An error has occurred during a query in the database", 500)
			return
		}
		user.UserID = found.UserID
		user.Username = found.Username
	}

	errEncode := encoder.Encode(user)
	if errEncode != nil {
		http.Error(w, "An error has occurred while encoding the user infos", 500)
		return
	}
	w.WriteHeader(http.StatusOK)
}
