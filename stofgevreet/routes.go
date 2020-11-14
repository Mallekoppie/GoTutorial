package main

import (
	"net/http"

	"tutorial/stofgevreet/service"

	"github.com/Mallekoppie/goslow/platform"
)

var Routes = platform.Routes{
	platform.Route{
		Path:        "/point",
		Method:      http.MethodPost,
		HandlerFunc: service.SavePoint,
		SlaMs:       0,
	},
	platform.Route{
		Path:        "/scan",
		Method:      http.MethodPost,
		HandlerFunc: service.SaveScan,
		SlaMs:       0,
	},
	platform.Route{
		Path:        "/stopwatch",
		Method:      http.MethodPost,
		HandlerFunc: service.SaveStopwatch,
		SlaMs:       0,
	},
	platform.Route{
		Path:        "/car",
		Method:      http.MethodGet,
		HandlerFunc: service.SaveStopwatch,
		SlaMs:       0,
	},
	platform.Route{
		Path:        "/handicap",
		Method:      http.MethodPost,
		HandlerFunc: service.SaveStopwatch,
		SlaMs:       0,
	},
}

/*
TODO:
Create get to return all car data
Push for handicap
finish handicap table, sprocs and functions


*/
