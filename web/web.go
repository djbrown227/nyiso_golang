package main

import (
    "net/http"
    "html/template"
)

func main() {
    http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("web/static"))))
    
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        tmpl := template.Must(template.ParseFiles("web/templates/index.html"))
        tmpl.Execute(w, nil)
    })
    
    http.ListenAndServe(":8080", nil)
}
