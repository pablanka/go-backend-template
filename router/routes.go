package router

import "github.com/pablanka/go-backend-template/handlers"

var routes = Routes{
	&Route{
		"Index",
		"GET",
		"/",
		handlers.Index,
	},
	&Route{
		"App",
		"GET",
		"/webapp",
		handlers.Webapp,
	},
	&Route{
		"GETUser",
		"GET",
		"/api/user.json",
		handlers.GETUser,
	},
	&Route{
		"POSTUser",
		"POST",
		"/api/user",
		handlers.POSTUser,
	},
}
