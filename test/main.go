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

func home(w http.ResponseWriter, r *http.Request) {

	index := template.Must(template.ParseGlob("templates/*.html"))

	data := L.Page{
		Title: "Home",
		Navs: []L.Link{
			{Linkname: "Questions", Href: "/questions"},
			{Linkname: "Data Analytics", Href: "/analysis"},
			{Linkname: "Data Visulizations", Href: "/viz"},
			{Linkname: "Info", Href: "/info"},
		},
		TemplateRender: []string{"questions", "viz"},
	}

	index.ExecuteTemplate(w, "layout", data)
	index.ExecuteTemplate(w, "questions", nil)
	// index.ExecuteTemplate(os.Stdout, "layout", data) // To see from STDOUT
	fmt.Println()
}

func questionsHandler(w http.ResponseWriter, r *http.Request) {
	index := template.Must(template.ParseGlob("templates/*.html"))

	data := L.Page{
		Title: "Home",
		Navs: []L.Link{
			{Linkname: "Questions", Href: "/questions"},
			{Linkname: "Data Analytics", Href: "/analysis"},
			{Linkname: "Data Visulizations", Href: "/viz"},
			{Linkname: "Info", Href: "/info"},
		},
		TemplateRender: []string{"questions"},
	}

	index.ExecuteTemplate(w, "layout", data)
	fmt.Println()

	if r.Method == http.MethodPost {
		fmt.Fprint(w, "Questions with POST")
	}
}

func analysisHandler(w http.ResponseWriter, r *http.Request) {
	index := template.Must(template.ParseGlob("templates/*.html"))

	data := L.Page{
		Title: "Home",
		Navs: []L.Link{
			{Linkname: "Questions", Href: "/questions"},
			{Linkname: "Data Analytics", Href: "/analysis"},
			{Linkname: "Data Visulizations", Href: "/viz"},
			{Linkname: "Info", Href: "/info"},
		},
		TemplateRender: []string{"analysis"},
	}

	index.ExecuteTemplate(w, "layout", data)
	fmt.Println()
}

func vizHandler(w http.ResponseWriter, r *http.Request) {
	index := template.Must(template.ParseGlob("templates/*.html"))

	data := L.Page{
		Title: "Home",
		Navs: []L.Link{
			{Linkname: "Questions", Href: "/questions"},
			{Linkname: "Data Analytics", Href: "/analysis"},
			{Linkname: "Data Visulizations", Href: "/viz"},
			{Linkname: "Info", Href: "/info"},
		},
		TemplateRender: []string{"viz"},
	}

	index.ExecuteTemplate(w, "layout", data)
	fmt.Println()
}

func infoHandler(w http.ResponseWriter, r *http.Request) {
	index := template.Must(template.ParseGlob("templates/*.html"))

	data := L.Page{
		Title: "Home",
		Navs: []L.Link{
			{Linkname: "Questions", Href: "/questions"},
			{Linkname: "Data Analytics", Href: "/analysis"},
			{Linkname: "Data Visulizations", Href: "/viz"},
			{Linkname: "Info", Href: "/info"},
		},
		TemplateRender: []string{"info"},
	}

	index.ExecuteTemplate(w, "layout", data)
	fmt.Println()
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
