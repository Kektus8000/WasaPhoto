package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Kektus8000/WasaPhoto/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) GetUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	w.Header().Set("content-type", "application/json")

	// Recupera l'ID dell'utente di cui recuperare l'informazione, ottenendo errore 500 se vi sono problemi nella query
	userID, errConv := strconv.Atoi(ps.ByName("userID"))
	if errConv != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Viene recuperato l'utente, ritornando errore 500 se vi sono problemi nella query
	user, errQuery := rt.db.GetUserByID(userID)
	if errQuery != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Viene codificato il risultato, ritornando errore 500 se vi sono problemi
	errEncode := json.NewEncoder(w).Encode(user)
	if errEncode != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
