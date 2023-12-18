package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

// getHelloWorld is an example of HTTP endpoint that returns "Hello world!" as a plain text
func (rt *_router) commentPhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params) Comment {
	w.Header().Set("content-type", "text/plain")

	//Check utente
	userID, errConv := strconv.Atoi(ps.ByName("userID"))
	photoID, errConv2 := strconv.Atoi(ps.ByName("photoID"))
	if errConv != nil || errConv2 != nil {
		w.WriteHeader(http.StatusBadRequest)
		return nil
	}

	publisherID, errFetch := rt.db.getPhotoPublisher(photoID)
	if errFetch != nil || publisherID == -1 {
		http.Error(w, "An error has occurred while fetching the user who published the photo", 404)
		return nil
	}

	banned, _ := rt.db.checkBanned(publisherID, userID)
	if banned {
		http.Error(w, "You have been banned by the user who published the photo", 403)
		return nil
	}
	var comm string
	err := json.NewDecoder(r.Body).Decode(&comm)
	if err != nil {
		http.Error(w, "An error has occurred while decoding the Request Body", 400)
		return nil
	}

	commentID, errComm := rt.db.commentPhoto(photoID, comm, userID)
	if errComm != nil {
		http.Error(w, "An error has occurred while publishing the comment on the photo", 400)
		return nil
	}

	comment := Comment{CommentID: commentID, Comment: comm, PublisherID: userID, PhotoID: photoID}
	w.WriteHeader(http.StatusCreated)
	return comment
}
