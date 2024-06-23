package api

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/Kektus8000/WasaPhoto/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) SearchUsers(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	w.Header().Set("content-type", "application/json")

	// Check ID dell'Utente
	userID := Authenticate(r.Header.Get("Authorization"))
	if userID == -1 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	word, errRead := io.ReadAll(r.Body)
	if errRead != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	risultato, errQuery := rt.db.SearchUsers(userID, string(word[:]))
	if errQuery != nil {
		w.WriteHeader(http.StatusBadGateway)
		return
	}

	err := json.NewDecoder(r.Body).Decode(&risultato)
	if err != nil {
		w.WriteHeader(http.StatusBadGateway)
		return
	}
}
