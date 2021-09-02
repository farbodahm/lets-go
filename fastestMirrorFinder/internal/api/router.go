package api

import (
	"github.com/gorilla/mux"
)

func NewMuxRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	for _, route := range Routes {
		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(route.HandlerFunc)

	}

	return router
}
