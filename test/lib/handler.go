package lib

import (
	"html/template"
	"net/http"
	"os"
)

// ===============================================
// Types
// ===============================================

type Link struct {
	Linkname string
	Href     string
}

type Page struct {
	Title    string
	Location string
}

// type Page struct {
// 	Title          string
// 	Navs           []Link
// 	TemplateRender []string
// 	Cra            Crank
// }

// ===============================================
// Functions
// ===============================================

func GetNavBar() []Link {
	yes := []Link{
		{Linkname: "Questions", Href: "/questions"},
		{Linkname: "Data Analytics", Href: "/analysis"},
		{Linkname: "Data Visulizations", Href: "/viz"},
		{Linkname: "Info", Href: "/info"},
	}
	return yes
}

// Automates the creation of Handlers
// func MakeHandler(fn func(http.ResponseWriter, *http.Request, Page)) http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request, p Page) {
// 		layout := template.Must(template.ParseFiles("templates/layout.html"))

// 		layout.ExecuteTemplate(w, "header", p.Title)
// 		layout.ExecuteTemplate(w, "navNbody", GetNavBar())

// 		layout.ExecuteTemplate(w, "footer", nil)

// 	}
// }

func SetupPage(w http.ResponseWriter, p Page, debug bool) {
	layout := template.Must(template.ParseFiles("templates/layout.html"))

	loc := p.Location

	purpose := template.Must(template.ParseFiles("templates/" + loc + ".html"))

	layout.ExecuteTemplate(w, "header", p.Title)
	layout.ExecuteTemplate(w, "navNbody", GetNavBar())
	purpose.ExecuteTemplate(w, loc, nil)
	layout.ExecuteTemplate(w, "footer", nil)

	if debug {
		layout.ExecuteTemplate(os.Stdout, "header", p.Title)
		layout.ExecuteTemplate(os.Stdout, "navNbody", GetNavBar())
		purpose.ExecuteTemplate(os.Stdout, loc, nil)
		layout.ExecuteTemplate(os.Stdout, "footer", nil)
	}
}
