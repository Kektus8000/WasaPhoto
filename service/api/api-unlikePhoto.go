package api

import (
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

// getHelloWorld is an example of HTTP endpoint that returns "Hello world!" as a plain text
func (rt *_router) unlikePhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Set("content-type", "text/plain")

	//Check Utente
	userID, errConv1 := strconv.Atoi(ps.ByName("userID"))
	photoID, errConv2 := strconv.Atoi(ps.ByName("photoID"))
	if errConv1 != nil || errConv2 != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	result, errUpdate := rt.db.unlikePhoto(userID, photoID)
	if errUpdate != nil || result == false {
		http.Error(w, "You didn't like the photo, so you can't unlike it", 403)
		return
	}

	w.WriteHeader(http.StatusNoContent)
	return

}
