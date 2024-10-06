package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"
)

type CarDatabase struct {
	ID      int    `json:"id"`
	CarName string `json:"car_name"`
	Company string `json:"company"`
	Year    int    `json:"year"`
}

var cars []CarDatabase
var nextID int = 1

// GET Car Database
func getHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}
	fmt.Fprintf(w, "Car Database")
}

// POST Car Database
func postHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}
	var cardetails CarDatabase
	// Read body of the POST request
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Invalid read request body", http.StatusBadRequest)
		return
	}
	// Parse the JSON data
	err = json.Unmarshal(body, &cardetails)
	if err != nil {
		http.Error(w, "Invalid JSON format", http.StatusBadRequest)
		return
	}
	// Response back with thesame message
	w.Header().Set("Content/Type", "application/json")
	json.NewEncoder(w).Encode(cardetails)
}

// Update handler
func updateHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	id, err := extractID(r.URL.Path)

	if err != nil {
		http.Error(w, "Invalid car ID", http.StatusBadRequest)
		return
	}

	for i, car := range cars {
		if car.ID == id {
			json.NewDecoder(r.Body).Decode(&car)
			cars[i] = car
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(car)
			return
		}
	}
	http.Error(w, "car not found", http.StatusNotFound)
}

// Extracting ID from URL
func extractID(path string) (int, error) {
	parts := strings.Split(path, "/")
	if len(parts) < 3 {
		return 0, fmt.Errorf("invalid path")
	}
	return strconv.Atoi(parts[2])
}

func main() {
	http.HandleFunc("/get", getHandler)
	http.HandleFunc("/post", postHandler)
	http.HandleFunc("/put", updateHandler)
	fmt.Println("Server running on port: 8497")
	log.Fatal(http.ListenAndServe(":8497", nil))

}
