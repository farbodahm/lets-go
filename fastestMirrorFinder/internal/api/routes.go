package api

import "net/http"

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

var Routes = [...]Route{
	{
		"GetFastestMirror",
		"GET",
		"/fastest-mirror",
		GetFastestMirror,
	},
}
