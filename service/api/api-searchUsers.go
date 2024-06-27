package api

import (
	"encoding/json"
	"net/http"
	"fmt"

	"github.com/Kektus8000/WasaPhoto/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) SearchUsers(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	w.Header().Set("content-type", "plaintext")

	// Check ID dell'Utente
	userID := Authenticate(r.Header.Get("Authorization"))
	if userID == -1 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var word User
	err := json.NewDecoder(r.Body).Decode(&word)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	risultato, errQuery := rt.db.SearchUsers(userID, word.Username)
	if errQuery != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	fmt.Println(risultato)

	errEncode := json.NewEncoder(w).Encode(risultato)
	if errEncode != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
