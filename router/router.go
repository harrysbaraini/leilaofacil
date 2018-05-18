package router

import (
	"github.com/julienschmidt/httprouter"
	"log"
)

// Route defines a route
type Route struct {
	Name string
	Method string
	Pattern string
	Handler httprouter.Handle
}

// RouteGroup defines a list of routes related to the same endpoint
type RouteGroup struct {
	Prefix string
	Routes []Route
}

// Router extends httprouter.Router
type Router struct {
	HttpRouter *httprouter.Router
}

// InstallRoutes adds a list of routes to the router
func (r *Router) InstallRoutes(rg *RouteGroup) {
	for _, route := range rg.Routes {
		r.HttpRouter.Handle(route.Method, rg.Prefix + route.Pattern, route.Handler)

		log.Printf("Route registed: " + rg.Prefix + route.Pattern)
	}
}