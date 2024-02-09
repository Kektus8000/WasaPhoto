package api

import (
	"net/http"
	"strconv"

	"github.com/Kektus8000/WasaPhoto/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) UnfollowUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	w.Header().Set("content-type", "application/json")

	// Check ID dell'Utente
	userID, errConv := strconv.Atoi(ps.ByName("userID"))
	if errConv != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Check ID dell'Utente interessato
	followerID, errConv2 := strconv.Atoi(ps.ByName("followerID"))
	if errConv2 != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if !Authenticate(userID, r.Header.Get("Authorization")) {
		http.Error(w, "Authentification went wrong", 401)
		return
	}

	errUpdate := rt.db.UnFollowUser(userID, followerID)
	if errUpdate != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusNoContent)
	return
}
