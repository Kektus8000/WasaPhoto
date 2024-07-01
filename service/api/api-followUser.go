package api

import (
	"net/http"
	"strconv"
	"fmt"

	"github.com/Kektus8000/WasaPhoto/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
)

// FollowUser si occupa di effettuare il follow di un utente ad un altro
// Come parametri, prende l'ID dell'utente e quello dell'utente che si intende seguire
func (rt *_router) FollowUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	w.Header().Set("content-type", "application/json")

	// Check ID dell'Utente
	fmt.Println("FOLLOW")
	fmt.Println(r.Header.Get("Authorization"))
	userID := Authenticate(r.Header.Get("Authorization"))
	if userID == -1 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	tofollow, errConv := strconv.Atoi(ps.ByName("followerID"))
	if errConv != nil {
		fmt.Println(errConv)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	banned, errQuery := rt.db.CheckBanned(tofollow, userID)
	if errQuery != nil {
		fmt.Println(errQuery)
		w.WriteHeader(http.StatusInternalServerError)
		return
	} else if banned == true {
		w.WriteHeader(http.StatusForbidden)
		return
	}
	errUpdate := rt.db.FollowUser(userID, tofollow)
	if errUpdate != nil {
		fmt.Println(errUpdate)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

}
