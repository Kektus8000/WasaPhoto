package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Kektus8000/WasaPhoto/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
)

// getHelloWorld is an example of HTTP endpoint that returns "Hello world!" as a plain text
func (rt *_router) GetMyStream(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	w.Header().Set("content-type", "text/plain")

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

	var photoIDs []int
	for i := 0; i < len(followers); i++ {
		follower := followers[i]
		photos, errQuery := rt.db.GetPublishedPhotos(follower)
		if errQuery == nil {
			for j := 0; j < len(photos); j++ {
				photoIDs = append(photoIDs, photos[i])
			}
		}
	}
	errEncode := json.NewEncoder(w).Encode(photoIDs)
	if errEncode != nil {
		http.Error(w, "An error has occurred while encoding photo's infos", 400)
		return
	}
	w.WriteHeader(http.StatusFound)
	return
}
