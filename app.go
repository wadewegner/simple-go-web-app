package main

import (
  "os"
  "net/http"
  "html/template"
  "github.com/gorilla/mux"
)

var templates = template.Must(template.ParseFiles(
    "header.html", 
    "footer.html", 
    "home.html", 
    "route.html"))

type Page struct {
  Title string
  Param string
  Id string
}

func display(w http.ResponseWriter, tmpl string, data interface{}) {
  templates.ExecuteTemplate(w, tmpl, data)
}

func main() {

  port := os.Getenv("PORT")
  if port == "" {
    port = "3000"
  }
    
  r := mux.NewRouter()

  r.HandleFunc("/route1", route1Handler)
  r.HandleFunc("/route2/{param}", route2Handler)
  r.HandleFunc("/route3/{param}/{id:[0-9]+}", route3Handler)
  r.HandleFunc("/", homeHandler)

  http.Handle("/", r)
  http.ListenAndServe(":3000", nil)
}

func route1Handler(w http.ResponseWriter, r *http.Request) {
  templates.ExecuteTemplate(w, "route", &Page{Title: "Route"})
}

func route2Handler(w http.ResponseWriter, r *http.Request) {
  param :=  mux.Vars(r)["param"]
  templates.ExecuteTemplate(w, "route", Page{Title: "Route", Param: param})
}

func route3Handler(w http.ResponseWriter, r *http.Request) {
  param :=  mux.Vars(r)["param"]
  id :=  mux.Vars(r)["id"]

  templates.ExecuteTemplate(w, "route", Page{Title: "Route", Param: param, Id: id})
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
  templates.ExecuteTemplate(w, "home", &Page{Title: "Home"})
}
