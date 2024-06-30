package api

import (
	"net/http"
	"strconv"

	"github.com/Kektus8000/WasaPhoto/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) RecoverPhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	w.Header().Set("content-type", "application/json")

	// Check ID dell'Utente
	userID, errConv1 := strconv.Atoi(ps.ByName("userID"))
	if errConv1 != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	photoID, errConv2 := strconv.Atoi(ps.ByName("photoID"))
	if errConv2 != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	file := "/tmp/userProfile/" + strconv.Itoa(userID) + "/publishedPhotos/" + strconv.Itoa(photoID)

	http.ServeFile(w, r, file)
}
