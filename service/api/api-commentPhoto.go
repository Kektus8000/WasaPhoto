package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Kektus8000/WasaPhoto/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
)

// CommentPhoto si occupa di pubblicare un commento sotto la foto da parte di un utente
func (rt *_router) CommentPhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	w.Header().Set("content-type", "application/json")

	// Recupera l'ID dell'utente che effettua la richiesta
	userID := Authenticate(r.Header.Get("Authorization"))
	if userID == -1 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Check dell'ID della Foto
	photoID, errConv := strconv.Atoi(ps.ByName("photoID"))
	if errConv != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Cerca la foto con dato ID, ritornando errore 404 se non trovata o errore 500 per problemi di query
	publisherID, errFetch := rt.db.GetPhotoPublisher(photoID)
	if errFetch != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	} else if publisherID == 0 {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	// Controlla se l'utente è stato bannato dall'utente pubblicante, ritornando errore 403 se risulta bannato o errore 500 per problemi di query
	banned, errQuery := rt.db.CheckBanned(publisherID, userID)
	if errQuery != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	} else if banned == true {
		w.WriteHeader(http.StatusForbidden)
		return
	}

	// Legge il request body, ritornando errore 500 se vi sono problemi con la query

	var word Comment
	err := json.NewDecoder(r.Body).Decode(&word)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if len(word.Text) < 6 || len(word.Text) > 160 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Aggiunge il commento nella tabella del database, ritornando errore 500 se non possibile
	errComm := rt.db.CommentPhoto(photoID, word.Text, userID)
	if errComm != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
