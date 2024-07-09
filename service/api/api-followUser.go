package api

import (
	"net/http"
	"strconv"

	"github.com/Kektus8000/WasaPhoto/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
)

// FollowUser si occupa di effettuare il follow di un utente ad un altro
// Come parametri, prende l'ID dell'utente e quello dell'utente che si intende seguire
func (rt *_router) FollowUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	w.Header().Set("content-type", "application/json")

	// Check ID dell'Utente
	userID := Authenticate(r.Header.Get("Authorization"))
	if userID == -1 {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	tofollow, errConv := strconv.Atoi(ps.ByName("followerID"))
	if errConv != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	banned, errQuery := rt.db.CheckBanned(tofollow, userID)
	if errQuery != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	} else if banned == true {
		w.WriteHeader(http.StatusForbidden)
		return
	}
	errUpdate := rt.db.FollowUser(userID, tofollow)
	if errUpdate != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

}
