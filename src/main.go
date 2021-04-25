package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

// CONSTANTS
const (
	PORT = 9990
)

// PAGES

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "/")
	if r.Method == http.MethodPost {
		fmt.Fprint(w, "/ with POST")
	}
}

func info(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "info")
}

func results(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "results")
}

// MAIN FUNCTION

func main() {

	fs := http.FileServer(http.Dir("static/"))

	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/", home)
	http.HandleFunc("/info", info)
	http.HandleFunc("/results", results)

	if err := http.ListenAndServe(":"+strconv.Itoa(PORT), nil); err != nil {
		log.Fatal("Error Starting the HTTP Server: ", err)

	}
}
