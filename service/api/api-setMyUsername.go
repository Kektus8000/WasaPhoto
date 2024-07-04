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
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var newUser User
	err := json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	user2, errQuery := rt.db.GetUserByUsername(newUser.Username)
	if errQuery != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	} else if userID != user2.UserID {
		w.WriteHeader(http.StatusForbidden)
		return
	}

	errUpdate := rt.db.SetMyUsername(userID, newUser.Username)
	if errUpdate != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}
