package server

import (
	"net/http"
	"regexp"
)

// RblogHandler used to serve http
type RblogHandler struct{}

func (h RblogHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) { // multiplexer
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

	match, err = regexp.MatchString("^/pages/", path) // match for /assets/* URLs, that serve the /assets/public/ folder
	check(err)
	if match {
		servePage(w, r)
		return
	}

	match, err = regexp.MatchString("^/posts/*$", path) // match for /posts, /posts/, /posts// to serve the lists of posts
	check(err)
	if match {
		serveListPosts(w, r)
		return
	}

	match, err = regexp.MatchString("^/posts/", path) // match for /posts/post-name to serve a specific post
	check(err)
	if match {
		servePost(w, r)
		return
	}

	match, err = regexp.MatchString("^/resources/", path) // match for /resources/ to serve public resources
	check(err)
	if match {
		serveResource(w, r)
		return
	}

	serve404(w, r) // If none of the routes matched, go for 404 error
}
