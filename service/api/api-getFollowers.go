package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Kektus8000/WasaPhoto/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) GetFollowers(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	w.Header().Set("content-type", "application/json")

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

	followers, errFoll := rt.db.GetFollowers(userID)
	if errFoll != nil {
		http.Error(w, "An error has occurred during the query from the database", 400)
		return
	}

	errEncode := json.NewEncoder(w).Encode(followers)
	if errEncode != nil {
		http.Error(w, "An error has occurred while encoding photo's infos", 400)
		return
	}
	w.WriteHeader(http.StatusFound)
	return
}
