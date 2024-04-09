package api

import (
	"net/http"
	"os"
	"strconv"

	"github.com/Kektus8000/WasaPhoto/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) DeletePhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
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

	if userID != publisherID {
		http.Error(w, "You can't delete the photo because it isn't your photo", 403)
		return
	}

	_, errQuery := rt.db.DeletePhoto(photoID)
	if errQuery != nil {
		http.Error(w, "There isn't a photo with that Id, so it can't be removed", 404)
		return
	}

	errOS := os.Remove("main/userProfile/" + strconv.Itoa(userID) + "/publishedPhotos/" + strconv.Itoa(photoID))
	if errOS != nil {
		http.Error(w, "An error has occurred while uploading the photo", 400)
		return
	}
	w.WriteHeader(http.StatusNoContent)
	return

}
