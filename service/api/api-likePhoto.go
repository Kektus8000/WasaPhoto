package api

import (
	"net/http"
	"strconv"

	"github.com/Kektus8000/WasaPhoto/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) LikePhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	w.Header().Set("content-type", "application/json")

	// Check dell'ID dell'Utente
	userID := Authenticate(r.Header.Get("Authorization"))
	if userID == -1 {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	// Check dell'ID della Foto
	photoID, errConv2 := strconv.Atoi(ps.ByName("photoID"))
	if errConv2 != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	publisherID, errQuery := rt.db.GetPhotoPublisher(photoID)
	if errQuery != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	banned, errQuery2 := rt.db.CheckBanned(publisherID, userID)
	if errQuery2 != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	} else if banned == true {
		w.WriteHeader(http.StatusForbidden)
		return
	}

	errUpdate := rt.db.LikePhoto(userID, photoID)
	if errUpdate != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
