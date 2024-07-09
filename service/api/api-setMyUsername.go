package api

import (
	"encoding/json"
	"net/http"

	"github.com/Kektus8000/WasaPhoto/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) SetMyUsername(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	w.Header().Set("content-type", "application/json")

	// Check ID dell'Utente
	userID := Authenticate(r.Header.Get("Authorization"))
	if userID == -1 {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	var newUser User
	err := json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	result, errQuery := rt.db.AlreadyInUse(userID, newUser.Username)
	if errQuery != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	} else if result {
		w.WriteHeader(http.StatusForbidden)
		return
	}

	errUpdate := rt.db.SetMyUsername(userID, newUser.Username)
	if errUpdate != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}
