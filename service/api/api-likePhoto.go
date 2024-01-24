package api

import (
	"net/http"
	"strconv"

	"github.com/Kektus8000/WasaPhoto/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
)

// getHelloWorld is an example of HTTP endpoint that returns "Hello world!" as a plain text
func (rt *_router) LikePhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	w.Header().Set("content-type", "text/plain")

	//Check Utente
	userID, errConv1 := strconv.Atoi(ps.ByName("userID"))
	if errConv1 != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if !Authenticate(userID, r.Header.Get("Authorization")) {
		http.Error(w, "Authentification went wrong", 401)
		return
	}
	photoID, errConv2 := strconv.Atoi(ps.ByName("photoID"))
	if errConv2 != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	publisherID, errQuery := rt.db.GetPhotoPublisher(photoID)
	if errQuery != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	banned, errQuery2 := rt.db.CheckBanned(publisherID, userID)
	if errQuery2 != nil {
		w.WriteHeader(http.StatusForbidden)
		return
	} else if banned == true {
		http.Error(w, "You have been banned by the user who published the photo", 403)
		return
	}

	errUpdate := rt.db.LikePhoto(userID, photoID)
	if errUpdate != nil {
		w.WriteHeader(http.StatusBadGateway)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
