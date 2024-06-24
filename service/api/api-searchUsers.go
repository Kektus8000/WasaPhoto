package api

import (
	"encoding/json"
	"net/http"

	"github.com/Kektus8000/WasaPhoto/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
)

type Ricerca struct {
	Ricercato string
}


func (rt *_router) SearchUsers(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	w.Header().Set("content-type", "plaintext")

	// Check ID dell'Utente
	userID := Authenticate(r.Header.Get("Authorization"))
	if userID == -1 {
		http.Error(w, "Errore all'autenticazione", 400)
		return
	}

	var word Ricerca
	err := json.NewDecoder(r.Body).Decode(&word)
	if err != nil {
		http.Error(w, "Errore nel decode", 401)
		return
	}

	risultato, errQuery := rt.db.SearchUsers(userID, word.Ricercato)
	if errQuery != nil {
		http.Error(w, "Errore nella query", 402)
		return
	}

	errEncode := json.NewEncoder(w).Encode(risultato)
	if errEncode != nil {
		http.Error(w, "Errore nell'encode", 403)
		return
	}
}
