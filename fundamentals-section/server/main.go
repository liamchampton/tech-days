package main

import (
	"fmt"
	"net/http"
)

// welcomeHandler is a handler function.
func welcomeHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Hello, friends! Welcome to the Microsoft Tech Days with Liam & Adelina!\n")
}

func main() {
	// Register the welcomeHandler function to serve the root endpoint.
	http.HandleFunc("/", welcomeHandler)

	// Start the default router on port 4321 and block the main goroutine from terminating.
	fmt.Println("Listening on 4321...")
	http.ListenAndServe(":4321", nil)
}