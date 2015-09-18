package main

import (
        "log"
        "fmt"
        "github.com/gorilla/mux"
        "github.com/gorilla/handlers"
        "html/template"
        "net/http"
        "os"
)

type Page struct {
        Title string
        Error string
}

var templates = template.Must(template.ParseFiles("index.html", "admin.html"))

func LoginHandler(response http.ResponseWriter, request *http.Request) {
        if (request.Method == "GET") {
                response.Header().Set("Content-type", "text/html")
                err := request.ParseForm()
                if err != nil {
                        http.Error(response, fmt.Sprintf("error parsing url %v", err), 500)
                }
                templates.ExecuteTemplate(response, "index.html", Page{Title: "Login"})
        } else if (request.Method == "POST") {
                request.ParseForm()
                username := request.PostFormValue("inputEmail")
                password := request.PostFormValue("inputPassword")

                if (username == "admin@example.com" && password == "P4$$w0rd!") {
			log.Println("Successfull login for admin@example.com")
                        http.Redirect(response, request, "/admin", 302)
                } else {
			log.Println("Login failed for admin@example.com")
			page := Page{Title: "Login", Error: "Username or Password Invalid"}
                        templates.ExecuteTemplate(response, "index.html", page)
                }
        }
}

func AdminHandler(response http.ResponseWriter, request *http.Request) {
        response.Header().Set("Content-type", "text/html")
        err := request.ParseForm()
        if err != nil {
                http.Error(response, fmt.Sprintf("error parsing url %v", err), 500)
        }
        templates.ExecuteTemplate(response, "admin.html", Page{Title: "Admin"})
}

func main() {
        r := mux.NewRouter()
        r.Handle("/", handlers.LoggingHandler(os.Stdout, http.HandlerFunc(LoginHandler)))
        r.Handle("/admin", handlers.LoggingHandler(os.Stdout, http.HandlerFunc(AdminHandler)))
        r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
        http.Handle("/", r)

        err := http.ListenAndServe("localhost:8080", r)
        if err != nil {
                log.Fatal("Error starting server: ", err)
        }
}
