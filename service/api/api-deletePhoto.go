package api

import (
	"net/http"
	"os"
	"strconv"

	"github.com/Kektus8000/WasaPhoto/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) DeletePhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	w.Header().Set("content-type", "application/json")

	// Recupera l'ID dell'utente che effettua la richiesta
	userID := Authenticate(r.Header.Get("Authorization"))
	if userID == -1 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Recupera l'ID della foto da cancellare, ottenendo errore 400 se non è possibile recuperarlo
	photoID, errConv2 := strconv.Atoi(ps.ByName("photoID"))
	if errConv2 != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Ottiene l'ID del pubblicatore della foto, ottenendo errore 500 se vi sono errori nella query
	publisherID, errFetch := rt.db.GetPhotoPublisher(photoID)
	if errFetch != nil || publisherID == 0 {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Se l'ID del richiedente e del pubblicante non sono lo stesso, viene ritornato errore 403 ed impedita l'operazione
	if userID != publisherID {
		w.WriteHeader(http.StatusForbidden)
		return
	}

	// Cancella la foto dal database, ottenendo errore 500 se vi sono errori nella query
	_, errQuery := rt.db.DeletePhoto(photoID)
	if errQuery != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Cancella la foto dal server, ottenendo errore 500 se vi sono errori nella query
	errOS := os.Remove(PHOTOFOLDER + strconv.Itoa(userID) + PUBLISHEDFOLDER + strconv.Itoa(photoID))
	if errOS != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
