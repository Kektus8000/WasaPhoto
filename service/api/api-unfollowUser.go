package api

import (
	"net/http"
	"strconv"
	"fmt"

	"github.com/Kektus8000/WasaPhoto/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) UnfollowUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	w.Header().Set("content-type", "application/json")

	// Check ID dell'Utente
	fmt.Println("UNFOLLOW")
	userID := Authenticate(r.Header.Get("Authorization"))
	if userID == -1 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Check ID dell'Utente interessato
	followerID, errConv2 := strconv.Atoi(ps.ByName("followerID"))
	if errConv2 != nil {
		fmt.Println(errConv2)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	errUpdate := rt.db.UnFollowUser(userID, followerID)
	if errUpdate != nil {
		fmt.Println(errUpdate)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	
}
