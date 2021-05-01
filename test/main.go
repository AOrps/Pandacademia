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
	fmt.Println()
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

		answers := L.Answers{
			Name:        r.FormValue("nm"),
			Feeling:     r.FormValue("feeling"),
			Temperature: r.FormValue("temp"),
			Sick:        r.FormValue("sick"),
			Sore:        r.FormValue("sore"),
			Contact:     r.FormValue("contact"),
			Location:    r.FormValue("location"),
		}

		answers.PrintAll()
		fmt.Println(answers)

		try := answers.CheckAll()
		fmt.Println(try)
		if try {
			fmt.Fprint(w, "Questions with POST")

			test := L.QuestionPage{
				Ask:       false,
				Questions: *L.MakeQuestionArr(),
				Code:      "2",
			}

			layout.ExecuteTemplate(w, "test-quest", test)
			layout.ExecuteTemplate(w, "footer", nil)
			return
		}

	}

	testData := L.QuestionPage{
		Ask:       true,
		Questions: *L.MakeQuestionArr(),
		Code:      "0",
	}

	layout.ExecuteTemplate(w, "test-quest", testData)
	layout.ExecuteTemplate(w, "footer", nil)

	fmt.Println()
	return

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
	http.HandleFunc("/test/", testFunc)

	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(PORT), nil))
}
