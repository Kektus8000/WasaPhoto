package api

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"

	"github.com/Kektus8000/WasaPhoto/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
)

// CommentPhoto si occupa del commento di una foto da parte di un utente
func (rt *_router) CommentPhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	w.Header().Set("content-type", "text/plain")

	userID, errConv := strconv.Atoi(ps.ByName("userID"))
	photoID, errConv2 := strconv.Atoi(ps.ByName("photoID"))
	if errConv != nil || errConv2 != nil {
		w.WriteHeader(http.StatusBadRequest)
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

	commentID, errComm := rt.db.CommentPhoto(photoID, comm, userID)
	if errComm != nil || commentID == -1 {
		http.Error(w, "An error has occurred while publishing the comment on the photo", 400)
		return
	}

	var comment Comment = Comment{CommentID: commentID, Comment: comm, PublisherID: userID, PhotoID: photoID}
	json.NewEncoder(w).Encode(comment)
	w.WriteHeader(http.StatusCreated)
	return
}
