package router

import "github.com/pablanka/go-backend-template/handlers"

var routes = Routes{
	&Route{
		Name:        "Index",
		Method:      "GET",
		Pattern:     "/",
		HandlerFunc: handlers.Index,
	},
	&Route{
		Name:        "App",
		Method:      "GET",
		Pattern:     "/webapp",
		HandlerFunc: handlers.Statics,
		IsStatic:    true,
	},
	&Route{
		Name:        "GETUser",
		Method:      "GET",
		Pattern:     "/api/user.json",
		HandlerFunc: handlers.GETUser,
	},
	&Route{
		Name:        "POSTUser",
		Method:      "POST",
		Pattern:     "/api/user",
		HandlerFunc: handlers.POSTUser,
	},
}
