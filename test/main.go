package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"

	L "./lib"
)

// CONSTANTS
const (
	PORT = 9991
)

type Header struct {
	Title string
}

func home(w http.ResponseWriter, r *http.Request) {

	layout := template.Must(template.ParseFiles("templates/layout.html"))

	layout.ExecuteTemplate(w, "header", "start")
	layout.ExecuteTemplate(w, "navNbody", L.GetNavBar())
	layout.ExecuteTemplate(w, "footer", nil)
	fmt.Println()
}

func questionsHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "results")

	if r.Method == http.MethodPost {
		fmt.Fprint(w, "Questions with POST")
	}
}

func analysisHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "results")
}

func vizHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "results")
}

func infoHandler(w http.ResponseWriter, r *http.Request) {

	page := L.Page{Title: "Info", Location: "info"}

	L.SetupPage(w, page, true)
}

func resultsHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "results")
}

// MAIN FUNCTION

func main() {

	fs := http.FileServer(http.Dir("static"))

	http.Handle("/assets/", http.StripPrefix("/assets", fs))

	http.HandleFunc("/", home)
	http.HandleFunc("/analysis/", analysisHandler)
	http.HandleFunc("/questions/", questionsHandler)
	http.HandleFunc("/viz/", vizHandler)
	http.HandleFunc("/info/", infoHandler)
	http.HandleFunc("/results/", resultsHandler)

	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(PORT), nil))
}

// PAGES

// func makeHandler(fn func(http.ResponseWriter, *http.Request, Page)) http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		m := valid

// 	}
// }

// data := Page{
// 	Title: "Home",
// 	Navs: []Link{
// 		{Linkname: "Questions", Href: "/questions", Done: true},
// 		{Linkname: "Data Analytics", Href: "/dataAna", Done: true},
// 		{Linkname: "Data Visulizations", Href: "/dataViz", Done: true},
// 		{Linkname: "Info", Href: "/info", Done: true},
// 	},
// 	TemplateRender: []RenderTemp{
// 		{TemplateName: "questions"},
// 		{TemplateName: "dataAna"},
// 		{TemplateName: "dataViz"},
// 		{TemplateName: "info", Render: false},
// 		{TemplateName: "results"},
// 	},
// }

// for _, val := range data.TemplateRender {
// 	tmplName := val.TemplateName
// 	if val.Render {
// 		index.ExecuteTemplate(w, tmplName, nil)
// 		fmt.Println(tmplName)
// 	}
// }

// NEW BASE
// index := template.Must(template.ParseGlob("templates/*.html"))

// header := Header{Title: "bruh"}

// index.ExecuteTemplate(w, "header", header)
// index.ExecuteTemplate(w, "navNbody", L.GetNavBar())
// index.ExecuteTemplate(w, "footer", nil)
