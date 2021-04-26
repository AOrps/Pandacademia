package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
)

// CONSTANTS
const (
	PORT = 9990
)

// CUSTOM STRUCTS
type Link struct {
	Linkname string
	Href     string
	Done     bool
}

type Header struct {
	Title string
}

type Page struct {
	Links []Link
}

// PAGES

// func makeHandler(fn func(http.ResponseWriter, *http.Request, string)) http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		m := valid

// 	}
// }

func home(w http.ResponseWriter, r *http.Request) {

	var allFiles []string
	files, err := ioutil.ReadDir("./templates")
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		filename := file.Name()
		if strings.HasSuffix(filename, ".html") {
			allFiles = append(allFiles, "./templates/"+filename)
		}
	}

	templates, err := template.ParseFiles(allFiles...)

	if err != nil {
		log.Fatal(err)
	}

	data := Page{
		Links: []Link{
			{Linkname: "Questions", Href: "/questions", Done: true},
			{Linkname: "Data Analytics", Href: "/info", Done: true},
			{Linkname: "Data Visulizations", Href: "/info", Done: true},
			{Linkname: "Info", Href: "/info", Done: true},
		},
	}

	s1 := templates.Lookup("header.html")
	s1.ExecuteTemplate(w, "header", &Header{Title: "Home"})
	fmt.Println()
	s4 := template.Lookup("navbar.html")
	// TEMPLATE Name must be referred to in  ExecuteTemplate
	s2 := templates.Lookup("base.html")
	s2.ExecuteTemplate(w, "content", data)
	fmt.Println()
	s3 := templates.Lookup("footer.html")
	s3.ExecuteTemplate(w, "footer", nil)
	fmt.Println()
	// s3.Execute(os.Stdout, nil)
	// head := template.Must(template.ParseFiles("templates/header.html"))
	// head.Execute(w, &Header{Title: "Home"})
	// index := template.Must(template.ParseFiles("templates/base.html"))
	// index.Execute(w, data)
	// foot := template.Must(template.ParseFiles("templates/footer.html"))
	// foot.Execute(w, nil)

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
