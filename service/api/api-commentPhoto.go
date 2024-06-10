package api

import (
	"io"
	"net/http"
	"strconv"

	"github.com/Kektus8000/WasaPhoto/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
)

// CommentPhoto si occupa di pubblicare un commento sotto la foto da parte di un utente
func (rt *_router) CommentPhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	w.Header().Set("content-type", "application/json")

	// Check dell'ID dell'Utente
	userID, errConv1 := strconv.Atoi(ps.ByName("userID"))
	if errConv1 != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Check dell'ID della Foto
	photoID, errConv2 := strconv.Atoi(ps.ByName("photoID"))
	if errConv2 != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if !Authenticate(userID, r.Header.Get("Authorization")) {
		http.Error(w, "Authentification went wrong", 401)
		return
	}

	publisherID, errFetch := rt.db.GetPhotoPublisher(photoID)
	if errFetch != nil || publisherID == -1 {
		http.Error(w, "An error has occurred while fetching the user who published the photo", 404)
		return
	}

	banned, errQuery2 := rt.db.CheckBanned(publisherID, userID)
	if errQuery2 != nil {
		w.WriteHeader(http.StatusForbidden)
		return
	} else if banned == true {
		http.Error(w, "You have been banned by the user who published the photo", 403)
		return
	}

	words, errRead := io.ReadAll(r.Body)
	if errRead != nil {
		http.Error(w, "An error has occured while reading the comment from the requestbody", 400)
		return
	}
	comm := string(words[:])

	if len(comm) < 6 || len(comm) > 160 {
		http.Error(w, "The comment doesn't follow the rules about comment's lenght, so it cannot be published", 400)
		return
	}

	_, errComm := rt.db.CommentPhoto(photoID, comm, userID)
	if errComm != nil {
		http.Error(w, "An error has occurred while publishing the comment on the photo", 400)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
