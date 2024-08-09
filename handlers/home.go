package handlers

import (
	"html/template"
	"net/http"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		renderErrorPage(w, "405 Wrong Request Method", http.StatusMethodNotAllowed)
		return
	}
	if r.URL.Path != "/" {
		renderErrorPage(w, "404 Page Not Found", http.StatusNotFound)
		return
	}
	template, err := template.ParseFiles("templates/index.html")
	if err != nil {
		renderErrorPage(w, "500 Internal Server Error", http.StatusInternalServerError)
		return
	}
	template.Execute(w, nil)
}
