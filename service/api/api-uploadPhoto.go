package api

import (
	"io"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

// getHelloWorld is an example of HTTP endpoint that returns "Hello world!" as a plain text
func (rt *_router) uploadPhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Set("content-type", "text/plain")

	//Check utente
	userID, errConv := strconv.Atoi(ps.ByName("userID"))
	photoID, errConv2 := strconv.Atoi(ps.ByName("photoID"))
	if errConv != nil || errConv2 != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	publisherID, errFetch := rt.db.getPhotoPublisher(photoID)
	if errFetch != nil || publisherID == -1 {
		http.Error(w, "An error has occurred while fetching the user who published the photo", 404)
		return
	}

	banned, _ := rt.db.checkBanned(publisherID, userID)
	if banned {
		http.Error(w, "You have been banned by the user who published the photo", 403)
		return
	}
	var comm string
	photo, errPhoto := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "An error has occurred while decoding the Request Body", 400)
		return
	}

	commentID, errComm := rt.db.commentPhoto(photoID, comm, userID)
	if errComm != nil {
		http.Error(w, "An error has occurred while publishing the comment on the photo", 400)
		return
	}

	comment := Comment{CommentID: commentID, Comment: comm, PublisherID: userID, PhotoID: photoID}
	w.WriteHeader(http.StatusCreated)
	return
}
