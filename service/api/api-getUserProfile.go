package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Kektus8000/WasaPhoto/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
)

// getHelloWorld is an example of HTTP endpoint that returns "Hello world!" as a plain text
func (rt *_router) GetUserProfile(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	w.Header().Set("content-type", "text/plain")

	userID, errConv := strconv.Atoi(ps.ByName("userID"))
	checkID, errConv2 := strconv.Atoi(ps.ByName("checkID"))
	if errConv != nil || errConv2 != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if !DoLogin(userID, r.Header.Get("Authorization")) {
		http.Error(w, "Authentification went wrong", 401)
		return
	}
	banned, errBan := rt.db.CheckBanned(checkID, userID)
	if errBan != nil {
		http.Error(w, "An error has occurred while fetching the request", 400)
		return
	} else if banned == true {
		http.Error(w, "You can't access the user profile because you'have been banned by them", 403)
		return
	}

	user, errUser := rt.db.GetUser(checkID)
	if errUser != nil {
		http.Error(w, "An error has occurred during the query", 400)
		return
	}

	followers, errFollowers := rt.db.GetFollowers(checkID)
	if errFollowers != nil {
		http.Error(w, "An error has occurred during a query in the database", 400)
		return
	}

	followings, errFollowings := rt.db.GetFollowings(checkID)
	if errFollowings != nil {
		http.Error(w, "An error has occurred during a query in the database", 400)
		return
	}

	banneds, errBanneds := rt.db.GetBanList(checkID)
	if errBanneds != nil {
		http.Error(w, "An error has occurred during a query in the database", 400)
		return
	}

	json.NewEncoder(w).Encode(user)
	json.NewEncoder(w).Encode(followers)
	json.NewEncoder(w).Encode(followings)
	json.NewEncoder(w).Encode(banneds)
}
