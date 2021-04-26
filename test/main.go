package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

// CONSTANTS
const (
	PORT = 9990
)

// var Links = []Link{
// 	{Linkname: "Questions", Href: "/questions", Done: true},
// 	{Linkname: "Data Analytics", Href: "/dataAna", Done: true},
// 	{Linkname: "Data Visulizations", Href: "/dataViz", Done: true},
// 	{Linkname: "Info", Href: "/info", Done: true},
// }

// CUSTOM STRUCTS
type Link struct {
	Linkname string
	Href     string
	Done     bool
}

// type RenderTemp struct {
// 	TemplateName string
// 	DoRender     bool
// }

type Page struct {
	Title string
	Navs  []Link
}

// PAGES

// func makeHandler(fn func(http.ResponseWriter, *http.Request, string)) http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		m := valid

// 	}
// }

func home(w http.ResponseWriter, r *http.Request) {
	data := Page{
		Title: "Home",
		Navs: []Link{
			{Linkname: "Questions", Href: "/questions", Done: true},
			{Linkname: "Data Analytics", Href: "/dataAna", Done: true},
			{Linkname: "Data Visulizations", Href: "/dataViz", Done: true},
			{Linkname: "Info", Href: "/info", Done: true},
		},
	}

	index := template.Must(template.ParseFiles("templates/layout.html"))
	index.Execute(w, data)
	// index.Execute(os.Stdout, data)
	fmt.Println()
}

func questionsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		fmt.Fprint(w, "Questions with POST")
	}
}

func dataAnaHandler(w http.ResponseWriter, r *http.Request) {

}

func dataVizHandler(w http.ResponseWriter, r *http.Request) {

}

func infoHandler(w http.ResponseWriter, r *http.Request) {

}

func resultsHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "info")
}

// MAIN FUNCTION

func main() {

	fs := http.FileServer(http.Dir("static"))

	http.Handle("/assets/", http.StripPrefix("/assets", fs))

	http.HandleFunc("/", home)
	http.HandleFunc("/info/", infoHandler)

	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(PORT), nil))
}
