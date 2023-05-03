package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"time"
)

type Response struct {
	Message string `json:"message"`
}

func main() {
	// Seed the random number generator with the current time
	rand.Seed(time.Now().UnixNano())

	// Define a logger middleware that logs the incoming request
	logger := func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			log.Printf("%s %s %s", r.RemoteAddr, r.Method, r.URL)
			next(w, r)
		}
	}

	// Define the HTTP handler for the root endpoint
	http.HandleFunc("/", logger(func(w http.ResponseWriter, r *http.Request) {
		// Create a random message
		message := Response{Message: getRandomMessage()}

		// Convert the message to JSON
		jsonResponse, err := json.Marshal(message)
		if err != nil {
			log.Fatal(err)
		}

		// Set the Content-Type header and write the JSON response
		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonResponse)
	}))

	// Start the HTTP server on port 3000
	log.Fatal(http.ListenAndServe(":3000", nil))
}

func getRandomMessage() string {
	messages := []string{
		"Hello, World!",
		"Welcome to my Golang app!",
		"Have a nice day!",
		"Thanks for stopping by!",
		"This is only a test.",
		"Do not adjust your set.",
		"One of another set of responses",
	}

	return messages[rand.Intn(len(messages))]
}
