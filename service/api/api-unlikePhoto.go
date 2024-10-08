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
		w.WriteHeader(http.StatusNotFound)
		return
	}

	// Check dell'ID della Foto
	photoID, errConv := strconv.Atoi(ps.ByName("photoID"))
	if errConv != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	errUpdate := rt.db.UnlikePhoto(userID, photoID)
	if errUpdate != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
