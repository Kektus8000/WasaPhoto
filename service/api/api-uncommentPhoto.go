package api

import (
	"net/http"
	"strconv"

	"github.com/Kektus8000/WasaPhoto/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) UncommentPhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	w.Header().Set("content-type", "application/json")

	// Recupera l'ID dell'utente che effettua la richiesta
	userID := Authenticate(r.Header.Get("Authorization"))
	if userID == -1 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Conversione in int dell'ID del commento da rimuovere, ritornando errore 500 se vi sono errori nella query
	commentID, errConv2 := strconv.Atoi(ps.ByName("commentID"))
	if errConv2 != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Rimuove il commento dal database, ritornando errore 500 se vi sono errori nella query
	errQuery := rt.db.UncommentPhoto(commentID)
	if errQuery != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
