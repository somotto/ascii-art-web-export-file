package handlers

import (
	"fmt"
	"net/http"
	"time"
)

func ExportHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		renderErrorPage(w, "405 Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	asciiArt := r.FormValue("ascii_art")

	filename := fmt.Sprintf("ascii_art_%s.txt", time.Now().Format("20060102_150405"))

	w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=%s", filename))
	w.Header().Set("Content-Type", "text/plain")
	w.Header().Set("Content-Length", fmt.Sprintf("%d", len(asciiArt)))

	_, err := w.Write([]byte(asciiArt))
	if err != nil {
		renderErrorPage(w, "500 Internal Server Error: Unable to write file", http.StatusInternalServerError)
		return
	}
}
