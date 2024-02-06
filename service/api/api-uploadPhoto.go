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
	w.Header().Set("content-type", "text/plain")

	// Check dell'ID dell'Utente
	userID, errConv := strconv.Atoi(ps.ByName("userID"))
	if errConv != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if !Authenticate(userID, r.Header.Get("Authorization")) {
		http.Error(w, "Authentification went wrong", 401)
		return
	}

	multipart, errMulti := r.MultipartReader()
	if errMulti != nil {
		http.Error(w, "An error has occurred while decoding the photo", 400)
		return
	}

	photoFile, errFile := multipart.NextPart()
	defer photoFile.Close()
	if errFile != nil {
		http.Error(w, "An error has occurred while decoding the photo", 400)
		return
	}
	photoID, errQuery := rt.db.UploadPhoto(photoFile.FileName(), userID)
	if errQuery != nil || photoID == -1 {
		http.Error(w, "An error has occurred while uploading the photo", 400)
		return
	}
	path, errOS := os.Create("main/userProfile/" + strconv.Itoa(userID) + "/publishedPhotos/" + strconv.Itoa(photoID))
	if errOS != nil {
		http.Error(w, "An error has occurred while uploading the photo", 400)
		return
	}

	// Aggiunge la foto nella cartella appena creata, sotto l'ID richiesto
	_, errIO := io.Copy(path, photoFile)
	if errIO != nil {
		http.Error(w, "An error has occurred while uploading the photo", 400)
		return
	}
	w.WriteHeader(http.StatusCreated)
	return
}
