package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Kektus8000/WasaPhoto/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) GetMyStream(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	w.Header().Set("content-type", "application/json")

	// Recupera l'ID dell'utente che effettua la richiesta
	userID := Authenticate(r.Header.Get("Authorization"))
	if userID == -1 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Recupera lo stream di foto dell'utente, ottenendo errore 500 se vi sono problemi nella query
	photos, errFoll := rt.db.GetStream(userID)
	if errFoll != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Viene preparato lo stream dell'utente, ritornando errore 500 se vi sono problemi nella query dei commenti
	var stream []Photo
	for i := 0; i < len(photos); i++ {
		temp := photos[i]
		var photo Photo
		photo.PhotoID = temp.PhotoID
		photo.PublisherID = temp.PublisherID
		photo.PublicationDate = temp.PublicationDate
		photo.File = PHOTOFOLDER + strconv.Itoa(photo.PublisherID) + PUBLISHEDFOLDER + strconv.Itoa(photo.PhotoID)
		result, errComm := rt.db.GetComments(photo.PhotoID)

		if errComm != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		for j := 0; j < len(result); j++ {
			temp2 := result[i]
			var comm Comment
			comm.CommentID = temp2.CommentID
			comm.Comment = temp2.Comment
			comm.PhotoID = temp2.PhotoID
			comm.PublisherID = temp2.PublisherID
			photo.Comments = append(photo.Comments, comm)
		}

		likes, errComm := rt.db.GetLikes(photo.PhotoID)
		for k := 0; k < len(likes); k++ {
			temp3 := likes[k]
			var lover User
			lover.UserID = temp3.UserID
			lover.Username = temp3.Username
			photo.Likes = append(photo.Likes, lover)
		}
		stream = append(stream, photo)
	}

	// Viene codificato il risultato, ritornando errore 500 se vi sono problemi
	errEncode := json.NewEncoder(w).Encode(stream)
	if errEncode != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
