package main

import "net/http"

func main() {
	// Create a file server to serve files from the "assets" directory.
	// http.Dir("assets/") sets the root directory for the file server.
	fs := http.FileServer(http.Dir("assets/"))

	// Handle requests to the "/static/" URL path by serving files from "assets/".
	// http.StripPrefix("/static/", fs) removes "/static/" from the URL path
	// before the request is forwarded to the file server. For example, a request to
	// "/static/example.jpg" will look for "assets/example.jpg".
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// Start the HTTP server on port 8080. If users access "http://localhost:8080/static/",
	// they can access files from the "assets/" folder.
	http.ListenAndServe(":8080", nil)
}
