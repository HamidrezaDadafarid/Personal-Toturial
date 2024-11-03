package main

import (
	"log"
	"net/http"
)

// main is the entry point for the application
func main() {
	// Define the file path for the file to be served
	filePath := "./assets/1.png"

	// Set up an HTTP handler function for the "/my-file" route
	// When a request is made to "/my-file", this function will respond by serving the specified file
	http.HandleFunc("/my-file", func(w http.ResponseWriter, r *http.Request) {
		// Serve the file located at filePath to the client
		http.ServeFile(w, r, filePath)
	})

	// Start the HTTP server on port 8080 and listen for incoming requests
	// If an error occurs while starting the server, log the error and exit
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err) // Log the error if server fails to start
	}
}
