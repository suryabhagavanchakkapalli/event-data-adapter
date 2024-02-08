package main

import (
	"event-data-adapter/server"
	"fmt"
	"net/http"
	"sync"
)

var mutex sync.Mutex

func main() {
	// Create a channel for receiving requests
	requestChannel := make(chan map[string]interface{})

	// Start the worker
	go server.Worker(requestChannel)

	// HTTP endpoint to receive requests
	http.HandleFunc("/receive", func(w http.ResponseWriter, r *http.Request) {
		server.HandleRequest(w, r, requestChannel)
	})

	// Start the HTTP server
	fmt.Println("Server listening on :8080")
	http.ListenAndServe(":8080", nil)
}
