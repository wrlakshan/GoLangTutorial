package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {
	// Define a simple HTTP handler function
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, World!")
	})

	// Define a RESTful API endpoint for retrieving a list of bills
	http.HandleFunc("/api/bills", getBills)

	// Start the server on port 8080
	http.ListenAndServe(":8080", nil)
}

func getBills(w http.ResponseWriter, r *http.Request) {
	// Set the Content-Type header to application/json
	w.Header().Set("Content-Type", "application/json")

	// Define the bill struct
	type Bill struct {
		ID       int     `json:"id"`
		Amount   float64 `json:"amount"`
		Payee    string  `json:"payee"`
		Category string  `json:"category"`
		Date     string  `json:"date"`
	}

	// Initialize the bills slice
	bills := []Bill{
		{ID: 1, Amount: 100.0, Payee: "John Doe", Category: "Groceries", Date: "2022-01-01"},
		{ID: 2, Amount: 200.0, Payee: "Jane Smith", Category: "Rent", Date: "2022-02-01"},
	}

	// Encode the bills slice as JSON and send it in the response
	json.NewEncoder(w).Encode(bills)
}
