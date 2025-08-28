package main

import (
	"fmt"
	"net/http"

	asciiart "asciiart/functions"
)

// main starts the HTTP server and sets up route handlers for the ASCII art web app.
func main() {
	http.HandleFunc("/styles/", asciiart.Style)
	http.HandleFunc("/", asciiart.Home)
	http.HandleFunc("/ascii-art", asciiart.Ascii)
	fmt.Println("Server started: http://localhost:8080")

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error in starting Server")
		return
	}
}
