package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Kektus8000/WasaPhoto/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
)

// getHelloWorld is an example of HTTP endpoint that returns "Hello world!" as a plain text
func (rt *_router) setMyUsername(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	w.Header().Set("content-type", "text/plain")

	//Check utente
	userID, errConv := strconv.Atoi(ps.ByName("userID"))
	if errConv != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	var newUsername string

	err := json.NewDecoder(r.Body).Decode(&newUsername)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	errUpdate := rt.db.setMyUserName(userID, newUsername)
	if errUpdate != nil {
		w.WriteHeader(http.StatusBadGateway)
		return
	}

	w.WriteHeader(http.StatusNoContent)
	return
}
