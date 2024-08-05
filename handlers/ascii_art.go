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
        renderErrorPage(w, "Method Not Allowed", http.StatusMethodNotAllowed)
        return
    }

    text := r.FormValue("text")
    banner := r.FormValue("banner")

    if text == "" {
        renderErrorPage(w, "Bad Request: Missing text", http.StatusBadRequest)
        return
    }

    if !functions.InputString(text) {
        renderErrorPage(w, "Bad Request: Input contains non-ASCII characters", http.StatusBadRequest)
        return
    }

    fileName := banner + ".txt"
    if functions.FileName(fileName) == "" {
        renderErrorPage(w, "Not Found: Invalid banner", http.StatusNotFound)
        return
    }

    lines, err := functions.Readfile(fileName)
    if err != nil {
        if errors.Is(err, functions.ErrFileMissing) {
            renderErrorPage(w, "Not Found: Banner file missing", http.StatusNotFound)
        } else {
            renderErrorPage(w, "Internal Server Error: Unable to process banner file", http.StatusInternalServerError)
        }
        return
    }

    result := functions.AsciiArt(text, lines)

    tmpl, err := template.ParseFiles("templates/index.html")
    if err != nil {
        renderErrorPage(w, "Internal Server Error", http.StatusInternalServerError)
        return
    }

    data := PageData{Result: result}
    tmpl.Execute(w, data)
}

func renderErrorPage(w http.ResponseWriter, message string, statusCode int) {
    w.WriteHeader(statusCode)
    tmpl, err := template.ParseFiles("templates/error.html")
    if err != nil {
        http.Error(w, "Internal Server Error", http.StatusInternalServerError)
        return
    }
    tmpl.Execute(w, struct{ ErrorMessage string }{ErrorMessage: message})
}