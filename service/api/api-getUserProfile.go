package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/Kektus8000/WasaPhoto/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) GetUserProfile(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	w.Header().Set("content-type", "application/json")

	// ID dell'utente a cui si vuole accedere al profilo
	checkID, errConv := strconv.Atoi(ps.ByName("userID"))
	if errConv != nil {
		fmt.Println(errConv)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Recupera l'ID dell'utente che effettua la richiesta
	userID := Authenticate(r.Header.Get("Authorization"))
	if userID == -1 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Controlla se l'utente ricercato ha bannato l'utente richiedente
	banned, errBan := rt.db.CheckBanned(checkID, userID)
	if errBan != nil {
		fmt.Println(errBan)
		w.WriteHeader(http.StatusInternalServerError)
		return
	} else if banned == true {
		w.WriteHeader(http.StatusForbidden)
		return
	}

	var profilo UserProfile

	user, errUser := rt.db.GetUserByID(checkID)
	if errUser != nil {
		fmt.Println(errUser)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Ritorna gli IDs di tutti gli utenti che seguono l'utente
	followers, errFollowers := rt.db.GetFollowers(checkID)
	if errFollowers != nil {
		fmt.Println(errFollowers)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Ritorna gli IDs di tutti gli utenti seguiti dall'utente
	followings, errFollowings := rt.db.GetFollowings(checkID)
	if errFollowings != nil {
		fmt.Println(errFollowings)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Ritorna gli IDs di tutti gli utenti bloccati dall'utente
	banneds, errBanneds := rt.db.GetBanList(checkID)
	if errBanneds != nil {
		fmt.Println(errBanneds)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Ritorna gli IDs di tutte le foto pubblicate da tale utente
	photos, errPhotos := rt.db.GetPublishedPhotos(checkID)
	if errPhotos != nil {
		fmt.Println(errPhotos)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	// I valori ottenuti vengono salvati nella struct UserProfile appena creata
	profilo.UserID = user.UserID
	profilo.Username = user.Username

	for i := 0; i < len(followers); i++ {
		temp := followers[i]
		var follower User
		follower.UserID = temp.UserID
		follower.Username = temp.Username
		profilo.Followers = append(profilo.Followers, follower)
	}

	for i := 0; i < len(followings); i++ {
		temp := followings[i]
		var following User
		following.UserID = temp.UserID
		following.Username = temp.Username
		profilo.Followings = append(profilo.Followings, following)
	}

	for i := 0; i < len(banneds); i++ {
		temp := banneds[i]
		var ban User
		ban.UserID = temp.UserID
		ban.Username = temp.Username
		profilo.Banneds = append(profilo.Banneds, ban)
	}

	for i := 0; i < len(photos); i++ {
		temp := photos[i]
		var photo Photo
		photo.PhotoID = temp.PhotoID
		photo.PublisherID = temp.PublisherID
		photo.PublicationDate = temp.PublicationDate
		photo.File = "/tmp/userProfile/" + strconv.Itoa(photo.PublisherID) + "/publishedPhotos/" + strconv.Itoa(photo.PhotoID)

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

		profilo.PublishedPhotos = append(profilo.PublishedPhotos, photo)
	}
	// Viene creato un nuovo Encoder per trasformare i dati delle query in Json
	errEncode := json.NewEncoder(w).Encode(profilo)
	if errEncode != nil {
		fmt.Println(errEncode)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
