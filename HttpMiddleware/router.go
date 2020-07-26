package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func Serve() {

	router := mux.NewRouter()

	for index := range routes {
		route := routes[index]
		var handler http.Handler
		handler = route.HandlerFunc
		handler = Middleware1(handler, "value1")
		handler = Middleware2(handler, "value2")

		router.
			Path(route.Path).
			Methods(route.Method).
			Handler(handler)
	}

	http.ListenAndServe("0.0.0.0:9000", router)

}

type Routes []Route

type Route struct {
	Path        string
	Method      string
	HandlerFunc http.HandlerFunc
}

var routes = Routes{
	Route{
		Path:        "/agents",
		Method:      http.MethodGet,
		HandlerFunc: dummy,
	},
	Route{
		Path:        "/agents",
		Method:      http.MethodPost,
		HandlerFunc: dummy,
	},
}

func dummy(w http.ResponseWriter, r *http.Request) {

}
