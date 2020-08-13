package server

import (
	"net/http"
	"regexp"
)

// BlogmHandler used to serve http
type BlogmHandler struct{}

func (h BlogmHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) { // multiplexer
	switch r.Method {
	case "GET":
		// Static routes only, dynamic ones in default
		switch r.URL.Path {
		case "/":
			serveHomePage(w, r)
		default:
			dynamicRoutingGET(w, r)
		}
	case "POST":
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("method pas autorise lol"))
	}
}

// Handle non static routing using regexp
// for example assets, cdn, view_post and view_page endpoints
func dynamicRoutingGET(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path

	match, err := regexp.MatchString("^/assets/", path) // match for /assets/* URLs, that serve the /assets/public/ folder
	check(err)
	if match {
		serveAssets(w, r)
		return
	}

	w.WriteHeader(http.StatusNotFound) // If none of the routes matched, go for 404 error
	w.Write([]byte("404 pas trouve lol"))
}