package api

import (
	"encoding/json"
	"net/http"

	"github.com/Kektus8000/WasaPhoto/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) DoLogin(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	w.Header().Set("content-type", "application/json")

	var user User
	// Recupera le informazioni dell'utente, ritornano errore 400 se impossibile
	errDecode := json.NewDecoder(r.Body).Decode(&user)
	if errDecode != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Controlla se l'utente è nel database, ritornando errore 500 se vi sono errori nella query
	exist, errQuery := rt.db.UserExists(user.Username)
	if errQuery != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Se l'utente non esiste, viene aggiunto nel database
	if !exist {
		// Si aggiunge l'utente al database, ritornando errore 500 se vi sono errori nella query
		ID, errQuery := rt.db.AddUser(user.Username)
		if errQuery != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		user.UserID = ID
		w.WriteHeader(http.StatusCreated)
	} else if exist { // Se invece l'utente esiste, le sue informazioni vengono recuperate dal database
		// Vengono recuperate le informazioni dell'utente, ritornando errore 500 se vi sono errori nella query
		found, errQuery := rt.db.GetUserByUsername(user.Username)
		if errQuery != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		user.UserID = found.UserID
		user.Username = found.Username
	}

	// Le informazioni sull'utente vengono ritornate, ritornando errore 500 se vi sono errore nell'encode
	errEncode := json.NewEncoder(w).Encode(user)
	if errEncode != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
