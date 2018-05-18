package users

import "github.com/harrysbaraini/leilaofacil/router"

var controller = &Controller{Repository: Repository{}}

// Routes defines the list of routes handling Users
var Routes = &router.RouteGroup{
	Prefix: "/users",
	Routes: []router.Route {
		{
			"Index",
			"GET",
			"/",
			controller.Index,
		},
	},
}