package main

import (
	"log"
	"fmt"
	"github.com/gorilla/mux"
	"html/template"
	"net/http"
)

type Page struct {
	Title string
}

var templates = template.Must(template.ParseFiles("index.html"))

func LoginHandler(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-type", "text/html")
	err := request.ParseForm()
	if err != nil {
		http.Error(response, fmt.Sprintf("error parsing url %v", err), 500)
	}
	templates.ExecuteTemplate(response, "index.html", Page{Title: "Login"})
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", LoginHandler)
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.Handle("/", r)

	err := http.ListenAndServe("localhost:8080", r)
	if err != nil {
		log.Fatal("Error starting server: ", err)
	}
}




