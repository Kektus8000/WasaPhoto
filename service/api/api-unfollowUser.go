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
	userID := Authenticate(r.Header.Get("Authorization"))
	if userID == -1 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Check ID dell'Utente interessato
	followerID, errConv := strconv.Atoi(ps.ByName("followerID"))
	if errConv != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	errUpdate := rt.db.UnFollowUser(userID, followerID)
	if errUpdate != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

}
