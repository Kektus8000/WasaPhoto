package api

import (
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

func (rt *_router) deletePhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Set("content-type", "text/plain")

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

	if userID != publisherID {
		http.Error(w, "You can't unlike the photo because it isn't your photo", 403)
		return
	}

	result, errQuery := rt.db.deletePhoto(photoID)
	if errQuery != nil || result == false {
		http.Error(w, "There isn't a photo with that Id, so it can't be removed", 404)
		return
	}

	w.WriteHeader(http.StatusNoContent)
	return

}
