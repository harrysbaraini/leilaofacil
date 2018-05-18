package apiinfo

import "github.com/harrysbaraini/leilaofacil/router"

var controller = &Controller{}

// Routes ...
var Routes = &router.RouteGroup{
	Prefix: "",
	Routes: []router.Route{
		{
			"ApiInfo",
			"GET",
			"/",
			controller.Index,
		},
	},
}