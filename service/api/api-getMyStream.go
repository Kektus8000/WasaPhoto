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

	// Recupera l'ID dell'utente che effettua la richiesta
	userID := Authenticate(r.Header.Get("Authorization"))
	if userID == -1 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Recupera lo stream degli ID delle foto e l'ID del loro pubblicatore, ottenendo errore 500 se vi sono problemi nella query
	photoIDs, publisherIDs, errFoll := rt.db.GetStream(userID)
	if errFoll != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Viene preparato l'insieme di path delle foto
	var photos []string
	for i := 0; i < len(photoIDs); i++ {
		file := "./userProfile/" + strconv.Itoa(publisherIDs[i]) + "/publishedPhotos/" + strconv.Itoa(photoIDs[i])
		photos = append(photos, file)
	}

	// Viene codificato il risultato, ritornando errore 500 se vi sono problemi
	errEncode := json.NewEncoder(w).Encode(photos)
	if errEncode != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
