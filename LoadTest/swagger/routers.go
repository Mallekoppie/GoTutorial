/*
 * LoadTest.V1
 *
 * No description
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

type Route struct {
	Name          string
	Method        string
	Pattern       string
	HandlerFunc   http.HandlerFunc
	RequiredRoles []string
}

type Routes []Route

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		var handler http.Handler
		handler = route.HandlerFunc
		handler = Middleware(handler, route.Name, route.RequiredRoles)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}

	return router
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		Index,
		[]string{"user", "admin"},
	},

	Route{
		"GetSmallUser",
		strings.ToUpper("Get"),
		"/small/{userID}",
		GetSmallUser,
		[]string{},
	},

	Route{
		"HeadersDelete",
		strings.ToUpper("Delete"),
		"/headers",
		HeadersDelete,
		[]string{},
	},

	Route{
		"HeadersGet",
		strings.ToUpper("Get"),
		"/headers",
		HeadersGet,
		[]string{},
	},

	Route{
		"LargeSizedGet",
		strings.ToUpper("Get"),
		"/largeSized",
		LargeSizedGet,
		[]string{},
	},

	Route{
		"LargeSizedLargeIdGet",
		strings.ToUpper("Get"),
		"/largeSized/{largeId}",
		LargeSizedLargeIdGet,
		[]string{},
	},

	Route{
		"LargeSizedPost",
		strings.ToUpper("Post"),
		"/largeSized",
		LargeSizedPost,
		[]string{},
	},

	Route{
		"MediumSizedGet",
		strings.ToUpper("Get"),
		"/mediumSized",
		MediumSizedGet,
		[]string{},
	},

	Route{
		"MediumSizedMediumIdDelete",
		strings.ToUpper("Delete"),
		"/mediumSized/{mediumId}",
		MediumSizedMediumIdDelete,
		[]string{},
	},

	Route{
		"MediumSizedPost",
		strings.ToUpper("Post"),
		"/mediumSized",
		MediumSizedPost,
		[]string{},
	},

	Route{
		"ResponseCodeCheckGet",
		strings.ToUpper("Get"),
		"/responseCodeCheck",
		ResponseCodeCheckGet,
		[]string{},
	},

	Route{
		"SlowGet",
		strings.ToUpper("Get"),
		"/slow",
		SlowGet,
		[]string{},
	},

	Route{
		"SlowPost",
		strings.ToUpper("Post"),
		"/slow",
		SlowPost,
		[]string{},
	},

	Route{
		"SmallGet",
		strings.ToUpper("Get"),
		"/small",
		SmallGet,
		[]string{},
	},

	Route{
		"SmallUserIDDelete",
		strings.ToUpper("Delete"),
		"/small/{userID}",
		SmallUserIDDelete,
		[]string{},
	},

	Route{
		"SmallUserIDPost",
		strings.ToUpper("Post"),
		"/small/{userID}",
		SmallUserIDPost,
		[]string{},
	},

	Route{
		"TimeoutGet",
		strings.ToUpper("Get"),
		"/timeout",
		TimeoutGet,
		[]string{},
	},
}
