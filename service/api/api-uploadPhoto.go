package api

import (
	"io"
	"net/http"
	"os"
	"strconv"

	"github.com/Kektus8000/WasaPhoto/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) UploadPhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	w.Header().Set("content-type", "application/json")

	// Check dell'ID dell'Utente
	userID := Authenticate(r.Header.Get("Authorization"))
	if userID == -1 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	photoID, errQuery := rt.db.UploadPhoto(userID)
	if errQuery != nil || photoID == -1 {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	path, errOS := os.Create("./userProfile/" + strconv.Itoa(userID) + "/publishedPhotos/" + strconv.Itoa(photoID))
	if errOS != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Aggiunge la foto nella cartella appena creata, sotto l'ID richiesto
	_, errIO := io.Copy(path, r.Body)
	if errIO != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

}
