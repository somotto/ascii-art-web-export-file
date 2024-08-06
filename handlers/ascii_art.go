package handlers

import (
	"ascii/functions"
	"errors"
	"html/template"
	"net/http"
)

type PageData struct {
    Result string
}

func AsciiArtHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        renderErrorPage(w, "405 Method Not Allowed", http.StatusMethodNotAllowed)
        return
    }

    text := r.FormValue("text")
    banner := r.FormValue("banner")

    if text == "" {
        renderErrorPage(w, "400 Bad Request: Missing text", http.StatusBadRequest)
        return
    }

    if !functions.InputString(text) {
        renderErrorPage(w, "400 Bad Request: Input contains non-ASCII characters", http.StatusBadRequest)
        return
    }

    fileName := banner + ".txt"
    if functions.FileName(fileName) == "" {
        renderErrorPage(w, "404 Not Found: Invalid banner", http.StatusNotFound)
        return
    }

    lines, err := functions.Readfile(fileName)
    if err != nil {
        if errors.Is(err, functions.ErrFileMissing) {
            renderErrorPage(w, "404 Not Found: Banner file missing", http.StatusNotFound)
        } else {
            renderErrorPage(w, "500 Internal Server Error: Unable to process banner file", http.StatusInternalServerError)
        }
        return
    }

    result := functions.AsciiArt(text, lines)

    tmpl, err := template.ParseFiles("templates/index.html")
    if err != nil {
        renderErrorPage(w, "500 Internal Server Error", http.StatusInternalServerError)
        return
    }

    data := PageData{Result: result}
    tmpl.Execute(w, data)
}

func renderErrorPage(w http.ResponseWriter, message string, statusCode int) {
    w.WriteHeader(statusCode)
    tmpl, err := template.ParseFiles("templates/error.html")
    if err != nil {
        http.Error(w, "500\nInternal Server Error", http.StatusInternalServerError)
        return
    }
    tmpl.Execute(w, struct{ ErrorMessage string }{ErrorMessage: message})
}