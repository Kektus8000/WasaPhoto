package api

import (
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

// getHelloWorld is an example of HTTP endpoint that returns "Hello world!" as a plain text
func (rt *_router) likePhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Set("content-type", "text/plain")

	//Check Utente
	userID, errConv1 := strconv.Atoi(ps.ByName("userID"))
	if errConv1 != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	photoID, errConv2 := strconv.Atoi(ps.ByName("photoID"))
	if errConv2 != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	publisherID, errQuery := rt.db.getPhotoPublisher(photoID)
	if errQuery != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if rt.db.checkBanned(publisherID, userID)[0] {
		w.WriteHeader(http.StatusForbidden)
		return
	}

	errUpdate := rt.db.likePhoto(userID, photoID)
	if errUpdate != nil {
		w.WriteHeader(http.StatusBadGateway)
		return
	}

	w.WriteHeader(http.StatusNoContent)
	return

}
