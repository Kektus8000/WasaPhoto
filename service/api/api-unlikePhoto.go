package api

import (
	"net/http"
	"strconv"

	"github.com/Kektus8000/WasaPhoto/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
)

// getHelloWorld is an example of HTTP endpoint that returns "Hello world!" as a plain text
func (rt *_router) UnlikePhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	w.Header().Set("content-type", "text/plain")

	//Check Utente
	userID, errConv1 := strconv.Atoi(ps.ByName("userID"))
	photoID, errConv2 := strconv.Atoi(ps.ByName("photoID"))
	if errConv1 != nil || errConv2 != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if !DoLogin(userID, r.Header.Get("Authorization")) {
		http.Error(w, "Authentification went wrong", 401)
		return
	}
	errUpdate := rt.db.UnlikePhoto(userID, photoID)
	if errUpdate != nil {
		http.Error(w, "You didn't like the photo, so you can't unlike it", 403)
		return
	}

	w.WriteHeader(http.StatusNoContent)
	return

}
