package api

import (
	"net/http"
	"strconv"

	"github.com/Kektus8000/WasaPhoto/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) BanUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	w.Header().Set("content-type", "application/json")

	userID := Authenticate(r.Header.Get("Authorization"))
	if userID == -1 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	bannedID, errConv2 := strconv.Atoi(ps.ByName("bannedID"))
	if errConv2 != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	errUpdate := rt.db.BanUser(userID, bannedID)
	if errUpdate != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	errUnfollow := rt.db.UnFollowUser(userID, bannedID)
	if errUnfollow != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	errUnfollow2 := rt.db.UnFollowUser(bannedID, userID)
	if errUnfollow2 != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
