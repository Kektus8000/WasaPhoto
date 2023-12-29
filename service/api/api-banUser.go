package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Kektus8000/WasaPhoto/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
)

// getHelloWorld is an example of HTTP endpoint that returns "Hello world!" as a plain text
func (rt *_router) banUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	w.Header().Set("content-type", "text/plain")

	//Check utente
	userID, errConv := strconv.Atoi(ps.ByName("userID"))
	if errConv != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var banned string
	errDecode := json.NewDecoder(r.Body).Decode(&banned)
	if errDecode != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	bannedID, errConv2 := strconv.Atoi(banned)
	if errConv2 != nil || userID == bannedID {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	errUpdate := rt.db.banUser(userID, bannedID)
	if errUpdate != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusNoContent)
	return
}
