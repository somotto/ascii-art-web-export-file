package handlers

import (
	"html/template"
	"net/http"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodGet {
        renderErrorPage(w, "Wrong Request Method", http.StatusMethodNotAllowed)
        return
    }
    if r.URL.Path != "/" {
        renderErrorPage(w, "Page Not Found", http.StatusNotFound)
        return
    }
    tmpl, err := template.ParseFiles("templates/index.html")
    if err != nil {
        renderErrorPage(w, "Internal Server Error", http.StatusInternalServerError)
        return
    }
    tmpl.Execute(w, nil)
}