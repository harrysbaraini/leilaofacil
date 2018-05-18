package main

import (
	"net/http"
	"log"
	"github.com/julienschmidt/httprouter"
	"github.com/harrysbaraini/leilaofacil/router"
	"github.com/harrysbaraini/leilaofacil/users"
	"github.com/harrysbaraini/leilaofacil/apiinfo"
)

func main() {
	log.Printf("Server starting ...")

	r := &router.Router{
		HttpRouter: httprouter.New(),
	}

	log.Printf("HTTP ROUTER created.")

	installRoutes(r)

	log.Printf("Routes registered.")

	if err := http.ListenAndServe(":3000", r.HttpRouter); err != nil {
		log.Fatal(err)
	}

	log.Printf("Listen and serving on :3000")
}

func installRoutes(r *router.Router) {
	r.InstallRoutes(apiinfo.Routes)
	r.InstallRoutes(users.Routes)
}