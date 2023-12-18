package api

import (
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

// getHelloWorld is an example of HTTP endpoint that returns "Hello world!" as a plain text
func (rt *_router) getMyStream(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Set("content-type", "text/plain")

	//Check Utente
	userID, errConv := strconv.Atoi(ps.ByName("userID"))
	errConv != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	photos, err := rt.db.getMyStream(userID)
	if err != nil{
		
	}



}
