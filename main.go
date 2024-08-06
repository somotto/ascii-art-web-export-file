package main

import (
	"ascii/handlers"
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	if len(os.Args) != 1 {
		return
	}

	// Register the handler to serve static files
	http.HandleFunc("/static/", handlers.ServeStatic)

	// Register the handler for the home page.
	http.HandleFunc("/", handlers.HomeHandler)

	// Register the handler for the ASCII art generation.
	http.HandleFunc("/ascii-art", handlers.AsciiArtHandler)

	port := ":8000"

	fmt.Printf("Server is running on http://localhost%s", port)

	// Start the HTTP server and log any errors.
	log.Fatal(http.ListenAndServe(port, nil))
}
