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

	var quest L.QuestionPage


	layout.ExecuteTemplate(w, "header", "🩺 Daily Screening")
	layout.ExecuteTemplate(w, "navNbody", L.GetNavBar())
	

	if r.Method == http.MethodPost {
		answers := L.GetAnswers(r)

		fmt.Println(answers)

		try := answers.CheckAll()
		fmt.Println(try)
		if try {
			test := L.QuestionPage{
				Ask:       false,
				Questions: L.MakeQuestionArr(),
				Code:      answers.Calc(),
			}

			layout.ExecuteTemplate(w, "questions", test)
			layout.ExecuteTemplate(w, "footer", nil)
			return
		} else {
			fmt.Fprintf(w, "Fill out all questions!")
		}

	}


	quest = L.QuestionPage{
		Ask:       true,
		Questions: L.MakeQuestionArr(),
		Code:      0,
	}

	layout.ExecuteTemplate(w, "questions", quest)
	layout.ExecuteTemplate(w, "viz", nil)
	layout.ExecuteTemplate(w, "footer", nil)
	fmt.Println()
}

func questionsHandler(w http.ResponseWriter, r *http.Request) {

	layout := template.Must(template.ParseGlob("templates/*.html"))

	layout.ExecuteTemplate(w, "header", "🩺 Daily Screening")
	layout.ExecuteTemplate(w, "navNbody", L.GetNavBar())

	if r.Method == http.MethodPost {

		answers := L.GetAnswers(r)

		fmt.Println(answers)

		try := answers.CheckAll()
		fmt.Println(try)
		if try {
			test := L.QuestionPage{
				Ask:       false,
				Questions: L.MakeQuestionArr(),
				Code:      answers.Calc(),
			}

			layout.ExecuteTemplate(w, "questions", test)
			layout.ExecuteTemplate(w, "footer", nil)
			return
		} else {
			fmt.Fprintf(w, "Fill out all questions!")
		}

	}

	testData := L.QuestionPage{
		Ask:       true,
		Questions: L.MakeQuestionArr(),
		Code:      0,
	}

	layout.ExecuteTemplate(w, "questions", testData)
	layout.ExecuteTemplate(w, "footer", nil)

	fmt.Println()
	return

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

// MAIN FUNCTION

func main() {

	fs := http.FileServer(http.Dir("static"))

	http.Handle("/assets/", http.StripPrefix("/assets", fs))

	http.HandleFunc("/", home)
	http.HandleFunc("/analysis/", analysisHandler)
	http.HandleFunc("/questions/", questionsHandler)
	http.HandleFunc("/viz/", vizHandler)
	http.HandleFunc("/info/", infoHandler)

	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(PORT), nil))
}
