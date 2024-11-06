package main

import (
	"log"
	"net/http"
)

func main() {
	// Serve the current directory (including HTML file)
	fs := http.FileServer(http.Dir("."))
	http.Handle("/", fs)

	// Set the port for the server
	port := "8000"
	log.Printf("Starting server on http://localhost:%s\n", port)

	// Start the server
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}
}
