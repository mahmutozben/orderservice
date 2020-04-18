package interfaces

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

func NewRouter(handlers *HandlerRepository) *mux.Router {

	router := mux.NewRouter().StrictSlash(true)
	for _, route := range getRoutes(handlers) {
		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(route.HandlerFunc)
	}

	return router
}

type Routes []Route

func getRoutes(handlers *HandlerRepository) []Route {
	var routes = Routes{
		Route{
			"SaveUser",
			"POST",
			"/users",
			handlers.UserHandler.SaveUser,
		},
		Route{
			"GetUsers",
			"GET",
			"/users",
			handlers.UserHandler.GetUsers,
		},
		Route{
			"GetUser",
			"GET",
			"/users/{userId}",
			handlers.UserHandler.GetUser,
		},
	}

	return routes
}
