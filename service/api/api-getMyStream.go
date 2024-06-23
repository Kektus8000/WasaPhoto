package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Kektus8000/WasaPhoto/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) GetMyStream(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	w.Header().Set("content-type", "application/json")

	// Check ID dell'Utente
	userID := Authenticate(r.Header.Get("Authorization"))
	if userID == -1 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	photoIDs, publisherIDs, errFoll := rt.db.GetStream(userID)
	if errFoll != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var photos []string
	for i := 0; i < len(photoIDs); i++ {
		photo := photoIDs[i]
		publisher := publisherIDs[i]
		file := "./userProfile/" + strconv.Itoa(publisher) + "/publishedPhotos/" + strconv.Itoa(photo)
		photos = append(photos, file)
	}

	errEncode := json.NewEncoder(w).Encode(photos)
	if errEncode != nil {
		http.Error(w, "An error has occurred while encoding photo's infos", 400)
		return
	}
}
