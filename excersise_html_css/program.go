package main

import (
	"fmt"
	"net/http"
)

func main() {
	// Serve the HTML and CSS files
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "welcome.html")
	})
	http.HandleFunc("/style.css", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "style.css")
	})

	// Start the server
	fmt.Println("Listening on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
