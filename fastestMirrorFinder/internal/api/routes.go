package api

import (
	"net/http"

	"github.com/farbodahm/lets-go/fastestMirrorFinder/internal/api/handlers"
)

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
		handlers.GetFastestMirror,
	},
}
