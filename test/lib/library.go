package lib

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
)

// ===============================================
// Types
// ===============================================

type Page struct {
	Title    string
	Location string
}

type Results struct {
	Code string
}

type trialType struct {
	Templ    *template.Template
	Location string
}

// ===============================================
// Functions
// ===============================================

// Automates the creation of Handlers
func MakeHandler(fn func(http.ResponseWriter, *http.Request, Page)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

// Specific Function that Parses Page and Executes that Template
func SetupSinglePage(w http.ResponseWriter, p Page, debug bool) {
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

// Does all the backend and Routes the Questions,
func Backend(w http.ResponseWriter, r *http.Request) {
	yes := "bruh"
	fmt.Println(yes)
}
