package api

import (
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

// getHelloWorld is an example of HTTP endpoint that returns "Hello world!" as a plain text
func (rt *_router) followUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Set("content-type", "text/plain")

	//Check utente
	userID, errConv := strconv.Atoi(ps.ByName("userID"))
	toFollowID, errConv2 := strconv.Atoi(ps.ByName("toFollowID"))
	if errConv != nil || errConv2 != nil || userID == toFollowID {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	banned, errQuery := rt.db.checkBanned(toFollowID, userID)
	if errQuery != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	} else if banned == true {
		http.Error(w, "You can't add the user in you following list because you are in its banList", 403)
		return
	}

	errUpdate := rt.db.followUser(userID, toFollowID)
	if errUpdate != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusNoContent)
	return
}
