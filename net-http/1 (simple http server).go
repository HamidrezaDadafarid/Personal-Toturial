package main

import (
	"fmt"
	"net/http"
)

// hello is a handler function that writes "hello" as the response body.
// It takes two parameters: w, an http.ResponseWriter to write the response,
// and req, the incoming http request.
func hello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Hello ")
	// we can call Write() method on w too.
	w.Write([]byte("World\n"))
}

// headers is a handler function that writes each header of the incoming
// request to the response body. It iterates over each header in the request,
// then formats and writes it as "Header-Name: Header-Value" for each header
// line. This helps in debugging by showing all incoming request headers.
func headers(w http.ResponseWriter, req *http.Request) {
	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}

// main is the entry point of the program. It sets up two URL endpoints:
// "/hello" and "/headers", each associated with a specific handler function.
// It then starts an HTTP server on port 8080, which listens and serves
// incoming HTTP requests on these endpoints.
func main() {
	// Register the "/hello" route with the hello handler function.
	http.HandleFunc("/hello", hello)

	// Register the "/headers" route with the headers handler function.
	http.HandleFunc("/headers", headers)

	// Start the HTTP server on port 8080 and listen for incoming requests.
	// This function blocks and keeps the server running.
	http.ListenAndServe(":8080", nil)
}
