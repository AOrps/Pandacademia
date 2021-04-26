package lib

import "net/http"

type Link struct {
	Linkname string
	Href     string
}

type Page struct {
	Title          string
	Navs           []Link
	TemplateRender []string
}

// Automates the creation of Handlers
func MakeHandler(fn func(http.ResponseWriter, *http.Request, Page)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}
