package api

import (
	"net/http"
	"strconv"

	"github.com/Kektus8000/WasaPhoto/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
)

// getHelloWorld is an example of HTTP endpoint that returns "Hello world!" as a plain text
func (rt *_router) UnBanUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	w.Header().Set("content-type", "application/json")

	// Check ID dell'Utente
	userID := Authenticate(r.Header.Get("Authorization"))
	if userID == -1 {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	// Check ID dell'Utente interessato
	bannedID, errConv := strconv.Atoi(ps.ByName("bannedID"))
	if errConv != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	errUpdate := rt.db.UnbanUser(userID, bannedID)
	if errUpdate != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
