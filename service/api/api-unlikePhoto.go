package api

import (
	"net/http"
	"strconv"

	"github.com/Kektus8000/WasaPhoto/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) UnlikePhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	w.Header().Set("content-type", "application/json")

	// Check dell'ID dell'Utente
	userID := Authenticate(r.Header.Get("Authorization"))
	if userID == -1 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Check dell'ID della Foto
	photoID, errConv2 := strconv.Atoi(ps.ByName("photoID"))
	if errConv2 != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	errUpdate := rt.db.UnlikePhoto(userID, photoID)
	if errUpdate != nil {
		http.Error(w, "You didn't like the photo, so you can't unlike it", 403)
		return
	}

	w.WriteHeader(http.StatusNoContent)

}
