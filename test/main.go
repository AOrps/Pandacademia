package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"strconv"

	L "./lib"
)

// CONSTANTS
const (
	PORT = 9991
)

type questPage struct {
	Question string
	Type     string
}

func home(w http.ResponseWriter, r *http.Request) {
	layout := template.Must(template.ParseGlob("templates/*.html"))

	layout.ExecuteTemplate(w, "header", "🩺 Daily Screening")
	layout.ExecuteTemplate(w, "navNbody", L.GetNavBar())
	layout.ExecuteTemplate(w, "questions", nil)
	layout.ExecuteTemplate(w, "viz", nil)
	layout.ExecuteTemplate(w, "footer", nil)
	fmt.Println()

	if r.Method == http.MethodPost {
		fmt.Fprintf(w, "Root with POST")
	}
}

func questionsHandler(w http.ResponseWriter, r *http.Request) {

	page := L.Page{Title: "Questions", Location: "questions"}

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

// TEST
func testFunc(w http.ResponseWriter, r *http.Request) {
	layout := template.Must(template.ParseGlob("templates/*.html"))

	testData := []questPage{
		{Question: "Are you feeling sick?", Type: "radio"},
		{Question: "Do you have muscle soreness or respiratory trouble that you can’t attribute to another medical condition?", Type: "radio"},
		{Question: "Have you recently had close contact (within 6 feet of an infected person for at least 15 minutes) with someone with symptoms of COVID-19, tested for COVID-19, or diagnosed with COVID-19?", Type: "radio"},
		{Question: "Have you recently been in a nursing home, healthcare facility, or homeless shelter?", Type: "radio"},
	}

	layout.ExecuteTemplate(w, "header", "🩺 Daily Screening")
	layout.ExecuteTemplate(w, "navNbody", L.GetNavBar())
	layout.ExecuteTemplate(w, "test-quest", testData)
	layout.ExecuteTemplate(w, "footer", nil)

	// Debug
	layout.ExecuteTemplate(os.Stdout, "header", "🩺 Daily Screening")
	layout.ExecuteTemplate(os.Stdout, "navNbody", L.GetNavBar())
	layout.ExecuteTemplate(os.Stdout, "test-quest", testData)
	layout.ExecuteTemplate(os.Stdout, "footer", nil)

	fmt.Println()

	if r.Method == http.MethodPost {
		fmt.Fprint(w, "Questions with POST")
	}
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
	http.HandleFunc("/test/", testFunc)

	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(PORT), nil))
}
