package apiinfo

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
	"encoding/json"
)

// Controller ...
type Controller struct {}

type apiInfo struct{
	Name string `json:"name"`
	Version string `json:"version"`
}

func (c *Controller) Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	info := &apiInfo{
		Name: "LeilãoFácil",
		Version: "1.0",
	}

	data, _ := json.Marshal(info)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}