package api

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/Kektus8000/WasaPhoto/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) BanUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	w.Header().Set("content-type", "application/json")

	// Recupera l'ID dell'utente della sessione, ritornadno 400 se impossibile
	bannerID := Authenticate(r.Header.Get("Authorization"))
	if bannerID == -1 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Recupera l'ID dell'utente da bannare dai parametri, ritornando 500 se vi sono errori
	bannedID, errConv := strconv.Atoi(ps.ByName("bannedID"))
	if errConv != nil {
		fmt.Println(errConv)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Aggiunge l'utente alla lista di utenti bannati, ritornando 500 se vi sono errori nella query
	errUpdate := rt.db.BanUser(bannerID, bannedID)
	if errUpdate != nil {
		fmt.Println(errUpdate)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Vengono rimossi i follow tra i due utenti, ritornando 500 se vi sono errori nella query
	errUnfollow := rt.db.UnFollowUser(bannerID, bannedID)
	if errUnfollow != nil {
		fmt.Println(errUnfollow)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	errUnfollow2 := rt.db.UnFollowUser(bannedID, bannerID)
	if errUnfollow2 != nil {
		fmt.Println(errUnfollow2)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Rimuove tutti i commenti pubblicati dall'utente bannato alle foto dell'utente bannante,
	// ritornando 500 se vi sono errori nella query
	errRemove1 := rt.db.RemoveAllComments(bannerID, bannedID)
	if errRemove1 != nil {
		fmt.Println(errRemove1)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Rimuove tutti i mi piace pubblicati dall'utente bannato alle foto dell'utente bannante,
	// ritornando 500 se vi sono errori nella query
	errRemove2 := rt.db.RemoveAllLikes(bannerID, bannedID)
	if errRemove2 != nil {
		fmt.Println(errRemove2)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
