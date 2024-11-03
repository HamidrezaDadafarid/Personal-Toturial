package main

import (
	"log"
	"net/http"
)

func main() {
	filePath := "./assets/1.png"

	http.HandleFunc("/my-file", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, filePath)
	})

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
