package api

import (
	"encoding/json"
	"net/http"

	"github.com/Kektus8000/WasaPhoto/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) SearchUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	w.Header().Set("content-type", "application/json")

	risultato, errQuery := rt.db.GetUserByUsername(ps.ByName("searchName"))
	if errQuery != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if risultato.UserID == 0 {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	var ID int = risultato.UserID

	err := json.NewDecoder(r.Body).Decode(&ID)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
}
