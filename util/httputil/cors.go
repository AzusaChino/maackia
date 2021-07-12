package httputil

import (
	"net/http"
	"regexp"
)

var corsHeaders = map[string]string{
	"Access-Control-Allow-Headers":  "Accept, Authorization, Content-Type, Origin",
	"Access-Control-Allow-Methods":  "GET, POST, OPTIONS",
	"Access-Control-Expose-Headers": "Date",
	"Vary":                          "Origin",
}

func SetCors(w http.ResponseWriter, o *regexp.Regexp, r *http.Request) {
	origin := r.Header.Get("Origin")
	// no cors
	if origin == "" {
		return
	}
	for k, v := range corsHeaders {
		w.Header().Set(k, v)
	}
	// allows all origins
	if o.String() == "^(?:.*)$" {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		return
	}

	if o.MatchString(origin) {
		w.Header().Set("Access-Control-Allow-Origin", origin)
	}
}
