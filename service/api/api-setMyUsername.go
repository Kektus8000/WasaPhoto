package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Kektus8000/WasaPhoto/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) SetMyUsername(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	w.Header().Set("content-type", "text/plain")

	// Check ID dell'Utente
	userID, errConv := strconv.Atoi(ps.ByName("userID"))
	if errConv != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if !Authenticate(userID, r.Header.Get("Authorization")) {
		http.Error(w, "Authentification went wrong", 401)
		return
	}

	var newUsername string

	err := json.NewDecoder(r.Body).Decode(&newUsername)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}

	errUpdate := rt.db.SetMyUsername(userID, newUsername)
	if errUpdate != nil {
		w.WriteHeader(http.StatusBadGateway)
	}

	w.WriteHeader(http.StatusNoContent)
}
