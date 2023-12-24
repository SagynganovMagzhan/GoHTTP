package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type JsonRequest struct {
	Message string `json:"message"`
}

type JsonResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

func main() {
	// Await incoming POST requests on port 8080
	http.HandleFunc("/process", handleRequest)
	fmt.Println("Server is listening on port 8080...")
	http.ListenAndServe(":8080", nil)
}

func handleRequest(w http.ResponseWriter, r *http.Request) {
	// Receive JSON data from the request body
	decoder := json.NewDecoder(r.Body)
	var request JsonRequest
	err := decoder.Decode(&request)
	if err != nil {
		// Invalid JSON format
		http.Error(w, "Invalid JSON format", http.StatusBadRequest)
		return
	}

	// Check if the "message" field is present
	if request.Message == "" {
		// Return an HTTP error code 400 if the "message" field is absent or altered
		http.Error(w, "Invalid JSON message", http.StatusBadRequest)
		return
	}

	// Print the message to the server console
	fmt.Println("Received Message:", request.Message)

	// Send a JSON response
	response := JsonResponse{
		Status:  "success",
		Message: "Data successfully received",
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
