package api

import (
	"encoding/json"
	"net/http"
	"os"
	"strconv"

	"github.com/Kektus8000/WasaPhoto/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) DoLogin(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	w.Header().Set("content-type", "text/plain")

	var username string
	errDecode := json.NewDecoder(r.Body).Decode(&username)
	if errDecode != nil {
		http.Error(w, "An error has occurred while decoding the Request Body", 400)
		return
	}

	if len(username) < 3 || len(username) > 16 {
		http.Error(w, "Invalid length for the username, use another one", 403)
		return
	}
	exist, errQuery := rt.db.UserExists(username)
	encoder := json.NewEncoder(w)
	if errQuery != nil {
		http.Error(w, "An error has occurred during a query in the database", 400)
		return
	} else if !exist {
		ID, errQuery := rt.db.AddUser(username)
		if ID == -1 || errQuery != nil {
			http.Error(w, "An error has occurred during a query in the database", 400)
			return
		}
		_, errOS := os.Create("/user/" + username + strconv.Itoa(ID))
		if errOS != nil {
			http.Error(w, "An error has occurred while adding the user to the user list", 400)
			return
		}

		var newUser User = User{UserID: ID, Username: username}
		errEncode := encoder.Encode(newUser)
		if errEncode != nil {
			http.Error(w, "An error has occurred while encoding the user infos", 400)
			return
		}

		w.WriteHeader(http.StatusOK)
		return

	} else if exist {
		user, errQuery := rt.db.GetUserByUsername(username)
		if errQuery != nil {
			http.Error(w, "An error has occurred during a query in the database", 400)
			return
		}
		errEncode := encoder.Encode(user)
		if errEncode != nil {
			http.Error(w, "An error has occurred while encoding the user infos", 400)
			return
		}
		w.WriteHeader(http.StatusCreated)
	}
}
