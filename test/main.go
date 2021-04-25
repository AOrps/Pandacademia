package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"text/template"
)

// CONSTANTS
const (
	PORT = 9990
)

// PAGES

func home(w http.ResponseWriter, r *http.Request) {
	index := template.Must(template.ParseFiles("templates/base.html"))
	index.Execute(w, nil)

	if r.Method == http.MethodPost {
		fmt.Fprint(w, "/ with POST")
	}
}

func results(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "results")
}

func infoHandler(w http.ResponseWriter, r *http.Request) {

}

// MAIN FUNCTION

func main() {

	fs := http.FileServer(http.Dir("static"))

	http.Handle("/assets/", http.StripPrefix("/assets", fs))

	http.HandleFunc("/", home)
	http.HandleFunc("/info/", infoHandler)

	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(PORT), nil))
}
