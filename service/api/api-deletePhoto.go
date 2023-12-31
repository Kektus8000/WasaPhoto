package api

import (
	"net/http"
	"os"
	"strconv"

	"github.com/Kektus8000/WasaPhoto/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) DeletePhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
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

	if userID != publisherID {
		http.Error(w, "You can't unlike the photo because it isn't your photo", 403)
		return
	}

	photoFile, errQuery := rt.db.DeletePhoto(photoID)
	if errQuery != nil {
		http.Error(w, "There isn't a photo with that Id, so it can't be removed", 404)
		return
	}

	errOS := os.Remove("/userProfile/" + strconv.Itoa(userID) + "/publishedPhotos/" + strconv.Itoa(photoID) + photoFile)
	if errOS != nil {
		http.Error(w, "An error has occurred while uploading the photo", 400)
		return
	}
	w.WriteHeader(http.StatusNoContent)
	return

}
