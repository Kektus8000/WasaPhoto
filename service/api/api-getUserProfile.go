package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Kektus8000/WasaPhoto/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) GetUserProfile(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	w.Header().Set("content-type", "application/json")
	// Check ID dell'Utente interessato
	checkID, errConv := strconv.Atoi(ps.ByName("userID"))
	
	if errConv != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	userID := Authenticate(r.Header.Get("Authorization"))
	if userID == -1 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Controlla se l'utente ricercato ha bannato l'utente richiedente
	banned, errBan := rt.db.CheckBanned(checkID, userID)
	if errBan != nil {
		http.Error(w, "An error has occurred while fetching the request", 400)
		return
	} else if banned == true {
		http.Error(w, "You can't access the user profile because you'have been banned by them", 403)
		return
	}

	user, errUser := rt.db.GetUserByID(checkID)
	if errUser != nil {
		http.Error(w, "An error has occurred during the query", 400)
		return
	}

	// Ritorna gli IDs di tutti gli utenti che seguono l'utente
	followers, errFollowers := rt.db.GetFollowers(checkID)
	if errFollowers != nil {
		http.Error(w, "An error has occurred during a query in the database", 400)
		return
	}

	// Ritorna gli IDs di tutti gli utenti seguiti dall'utente
	followings, errFollowings := rt.db.GetFollowings(checkID)
	if errFollowings != nil {
		http.Error(w, "An error has occurred during a query in the database", 400)
		return
	}

	// Ritorna gli IDs di tutti gli utenti bloccati dall'utente
	banneds, errBanneds := rt.db.GetBanList(checkID)
	if errBanneds != nil {
		http.Error(w, "An error has occurred during a query in the database", 400)
		return
	}


	// Ritorna gli IDs di tutte le foto pubblicate da tale utente
	photos, errPhotos := rt.db.GetPublishedPhotos(checkID)
	if errPhotos != nil {
		http.Error(w, "An error has occurred during a query in the database", 400)
		return
	}

	// Viene creato un nuovo Encoder per trasformare i dati delle query in Json
	encoder := json.NewEncoder(w)
	errEncode1 := encoder.Encode(user.UserID)
	if errEncode1 != nil {
		http.Error(w, "An error has occurred while encoding the user infos", 400)
		return
	}

	errEncode2 := encoder.Encode(user.Username)
	if errEncode2 != nil {
		http.Error(w, "An error has occurred while encoding the user infos", 400)
		return
	}
	errEncode3 := encoder.Encode(followers)
	if errEncode3 != nil {
		http.Error(w, "An error has occurred while encoding the follower list", 400)
		return
	}

	errEncode4 := encoder.Encode(followings)
	if errEncode4 != nil {
		http.Error(w, "An error has occurred while encoding the following list", 400)
		return
	}

	errEncode5 := encoder.Encode(banneds)
	if errEncode5 != nil {
		http.Error(w, "An error has occurred while encoding the banList", 400)
		return
	}

	errEncode6 := encoder.Encode(photos)
	if errEncode6 != nil {
		http.Error(w, "An error has occurred while encoding the photos", 400)
		return
	}
}
