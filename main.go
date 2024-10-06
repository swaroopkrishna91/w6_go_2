package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type Car struct {
	ID     int    `json:"id"`
	Make   string `json:"make"`
	Model  string `json:"model"`
	Year   int    `json:"year"`
	Status string `json:"status"` // "available" or "sold"
}

var cars []Car

// Getting the Car Details
func getAllCars(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}
	fmt.Fprintf(w, "Car Databases \n")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(cars)
}

// Creating the Car Details
func createCar(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var newCar Car
	err := json.NewDecoder(r.Body).Decode(&newCar)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	newCar.ID = len(cars) + 1 // Assign a unique ID
	cars = append(cars, newCar)

	json.NewEncoder(w).Encode(newCar)
}

// Updating the Car Details
func updateCar(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, err := strconv.Atoi(r.URL.Path[len("/cars/"):])
	if err != nil {
		http.Error(w, "Invalid car ID", http.StatusBadRequest)
		return
	}

	var updatedCar Car
	err = json.NewDecoder(r.Body).Decode(&updatedCar)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	for i := range cars {
		if cars[i].ID == id {
			cars[i] = updatedCar
			json.NewEncoder(w).Encode(cars[i])
			return
		}
	}

	http.Error(w, "Car not found", http.StatusNotFound)
}

// Deleteing the Car Details
func deleteCar(w http.ResponseWriter, r *http.Request) {
	
	if r.Method != http.MethodGet {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}
	w.Header().Set("Content-Type", "application/json")

	id, err := strconv.Atoi(r.URL.Path[len("/cars/"):])
	if err != nil {
		http.Error(w, "Invalid car ID", http.StatusBadRequest)
		return
	}
	fmt.Fprintf(w, "Car Details Deleted \n")
	for i := range cars {
		if cars[i].ID == id {
			cars = append(cars[:i], cars[i+1:]...)
			return
		}
	}

	http.Error(w, "Car not found", http.StatusNotFound)
}

//Uncomment this if you want to see the orginial secret and line uncomment line no. 135
// var easterEggURL = "https://www.youtube.com/watch?v=dQw4w9WgXcQ&ab_channel=RickAstley"

// func easterEggsecret(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	fmt.Fprintf(w, `{"message": "Congratulations! Here's your secret!", "url": "%s"}`, easterEggURL)
// }

//New Section
var easterEgg = "🎉 You found the Easter Egg! 🐣"

func secret(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, `{"message": "Congratulations! Here's your secret!", "url": "%s"}`, easterEgg)
}


func main() {
	http.HandleFunc("/cars/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			getAllCars(w, r)
		case http.MethodPut:
			updateCar(w, r)
		case http.MethodDelete:
			deleteCar(w, r)
		case http.MethodPost:
			createCar(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	http.HandleFunc("/cars/secret", secret)
	// http.HandleFunc("/cars/orginialsecret", easterEggsecret)
	fmt.Println("Server listening on port 8497")
	http.ListenAndServe(":8497", nil)
}
