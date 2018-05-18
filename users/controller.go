package users

import (
	"net/http"
	"encoding/json"
	"github.com/julienschmidt/httprouter"
)

type Controller struct {
	Repository Repository
}

func (c *Controller) Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	users := c.Repository.GetUsers()

	data, _ := json.Marshal(users)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}