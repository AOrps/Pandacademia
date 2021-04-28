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

	page := L.Page{Title: "Screening", Location: "questions"}

	L.SetupSinglePage(w, page, false)

	if r.Method == http.MethodPost {
		fmt.Fprint(w, "Questions with POST")
	}
}

func analysisHandler(w http.ResponseWriter, r *http.Request) {
	page := L.Page{Title: "Data Analysis", Location: "analysis"}

	L.SetupSinglePage(w, page, false)
}

func vizHandler(w http.ResponseWriter, r *http.Request) {
	page := L.Page{Title: "Data Visualization", Location: "viz"}

	L.SetupSinglePage(w, page, false)
}

func infoHandler(w http.ResponseWriter, r *http.Request) {

	page := L.Page{Title: "Info", Location: "info"}

	L.SetupSinglePage(w, page, false)
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
