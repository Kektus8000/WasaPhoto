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
	userID, errConv := strconv.Atoi(ps.ByName("userID"))
	if errConv != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Check ID dell'Utente interessato
	toFollowID, errConv2 := strconv.Atoi(ps.ByName("toFollowID"))
	if errConv2 != nil || userID == toFollowID {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if !Authenticate(userID, r.Header.Get("Authorization")) {
		http.Error(w, "Authentification went wrong", 401)
		return
	}

	banned, errQuery := rt.db.CheckBanned(toFollowID, userID)
	if errQuery != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	} else if banned == true {
		http.Error(w, "You can't add the user in you following list because you are in its banList", 403)
		return
	}

	errUpdate := rt.db.FollowUser(userID, toFollowID)
	if errUpdate != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
