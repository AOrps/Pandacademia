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

type Answers struct {
	Name        string
	Feeling     string
	Temperature string
	Sick        string
	Sore        string
	Contact     string
	Location    string
}

type Question struct {
	QuestionText string
	Type         string
	NameElement  string
}

type QuestionPage struct {
	Ask       bool
	Questions []Question
	Code      string
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

	layout.ExecuteTemplate(w, "header", "🩺 Daily Screening")
	layout.ExecuteTemplate(w, "navNbody", L.GetNavBar())

	if r.Method == http.MethodPost {

		answers := Answers{
			Name:        r.FormValue("nm"),
			Feeling:     r.FormValue("feeling"),
			Temperature: r.FormValue("temp"),
			Sick:        r.FormValue("sick"),
			Sore:        r.FormValue("sore"),
			Contact:     r.FormValue("contact"),
			Location:    r.FormValue("location"),
		}

		fmt.Println(answers)

		fmt.Fprint(w, "Questions with POST")

		test := QuestionPage{
			Ask: false,
			Questions: []Question{
				{QuestionText: "Are you feeling sick?", Type: "radio", NameElement: "sick"},
				{QuestionText: "Do you have muscle soreness or respiratory trouble that you can’t attribute to another medical condition?", Type: "radio", NameElement: "sore"},
				{QuestionText: "Have you recently had close contact (within 6 feet of an infected person for at least 15 minutes) with someone with symptoms of COVID-19, tested for COVID-19, or diagnosed with COVID-19?", Type: "radio", NameElement: "contact"},
				{QuestionText: "Have you recently been in a nursing home, healthcare facility, or homeless shelter?", Type: "radio", NameElement: "location"},
			},
			Code: "2",
		}

		layout.ExecuteTemplate(w, "test-quest", test)
		layout.ExecuteTemplate(w, "footer", nil)
		return
	}

	testData := QuestionPage{
		Ask: true,
		Questions: []Question{
			{QuestionText: "Are you feeling sick?", Type: "radio", NameElement: "sick"},
			{QuestionText: "Do you have muscle soreness or respiratory trouble that you can’t attribute to another medical condition?", Type: "radio", NameElement: "sore"},
			{QuestionText: "Have you recently had close contact (within 6 feet of an infected person for at least 15 minutes) with someone with symptoms of COVID-19, tested for COVID-19, or diagnosed with COVID-19?", Type: "radio", NameElement: "contact"},
			{QuestionText: "Have you recently been in a nursing home, healthcare facility, or homeless shelter?", Type: "radio", NameElement: "location"},
		},
		Code: "0",
	}

	answers := Answers{
		Name:        r.FormValue("nm"),
		Feeling:     r.FormValue("feeling"),
		Temperature: r.FormValue("temp"),
		Sick:        r.FormValue("sick"),
		Sore:        r.FormValue("sore"),
		Contact:     r.FormValue("contact"),
		Location:    r.FormValue("location"),
	}

	fmt.Println(answers)

	layout.ExecuteTemplate(w, "test-quest", testData)
	layout.ExecuteTemplate(w, "footer", nil)

	// Debug
	// layout.ExecuteTemplate(os.Stdout, "header", "🩺 Daily Screening")
	// layout.ExecuteTemplate(os.Stdout, "navNbody", L.GetNavBar())
	// layout.ExecuteTemplate(os.Stdout, "test-quest", testData)
	// layout.ExecuteTemplate(os.Stdout, "footer", nil)

	fmt.Println()
	return

}

func resultsHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Results")
	// layout := template.Must(template.ParseFiles("templates/layout.html"))

	// results := template.Must(template.ParseFiles("templates/results.html"))

	// data := "1"

	// layout.ExecuteTemplate(w, "header", "🩺 Daily Screening")
	// layout.ExecuteTemplate(w, "navNbody", L.GetNavBar())
	// results.ExecuteTemplate(w, "results", data)
	// layout.ExecuteTemplate(w, "footer", nil)

	// layout.ExecuteTemplate(os.Stdout, "header", "🩺 Daily Screening")
	// layout.ExecuteTemplate(os.Stdout, "navNbody", L.GetNavBar())
	// results.ExecuteTemplate(os.Stdout, "requests", data)
	// layout.ExecuteTemplate(os.Stdout, "footer", nil)

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
