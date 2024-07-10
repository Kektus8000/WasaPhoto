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
		w.WriteHeader(http.StatusNotFound)
		return
	}

	// Conversione in int dell'ID del commento da rimuovere, ritornando errore 500 se vi sono errori nella query
	commentID, errConv := strconv.Atoi(ps.ByName("commentID"))
	if errConv != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	comment, errQuery := rt.db.GetComment(commentID)
	if errQuery != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	publisherID, errQuery2 := rt.db.GetPhotoPublisher(comment.PhotoID)
	if errQuery2 != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if userID == publisherID || userID == comment.PublisherID {
		errQuery := rt.db.UncommentPhoto(commentID)
		if errQuery != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	} else {
		w.WriteHeader(http.StatusForbidden)
		return
	}
	// Rimuove il commento dal database, ritornando errore 500 se vi sono errori nella query

}
