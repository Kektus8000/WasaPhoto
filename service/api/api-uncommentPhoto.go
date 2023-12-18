package api

import (
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

// getHelloWorld is an example of HTTP endpoint that returns "Hello world!" as a plain text
func (rt *_router) uncommentPhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Set("content-type", "text/plain")

	//Check utente
	commentID, errConv := strconv.Atoi(ps.ByName("commentID"))
	if errConv != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	errQuery := rt.db.uncommentPhoto(commentID)
	if errQuery != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusNoContent)
	return
}
