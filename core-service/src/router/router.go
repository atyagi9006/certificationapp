package router

import (
	"net/http"

	"github.com/atyagi9006/certificationapp/core-service/src/logger"
	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	routes := getRoutes()
	for _, route := range routes {
		var handler http.Handler

		handler = route.HandlerFunc
		handler = logger.LogNHandle(handler, route.Name)
		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}
	return router
}
