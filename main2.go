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
var currentID = 1

// Getting the Car Details
func getCars(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(cars)
}

// Getting the Car Details
func getCar(w http.ResponseWriter, r *http.Request) {
	idParam := r.URL.Path[len("/cars/"):] // Extract the ID from URL path
	id, err := strconv.Atoi(idParam)
	if err != nil {
		http.Error(w, "Invalid car ID", http.StatusBadRequest)
		return
	}

	for _, car := range cars {
		if car.ID == id {
			json.NewEncoder(w).Encode(car)
			return
		}
	}
	http.Error(w, "Car not found", http.StatusNotFound)
}

// Creating the Car Details
func createCar(w http.ResponseWriter, r *http.Request) {
	var car Car
	err := json.NewDecoder(r.Body).Decode(&car)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	car.ID = currentID
	currentID++
	cars = append(cars, car)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(car)
}

// Updating the Car Details
func updateCar(w http.ResponseWriter, r *http.Request) {
	idParam := r.URL.Path[len("/cars/"):] // Extract the ID from URL path
	id, err := strconv.Atoi(idParam)
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

	for index, car := range cars {
		if car.ID == id {
			updatedCar.ID = car.ID
			cars[index] = updatedCar
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(updatedCar)
			return
		}
	}
	http.Error(w, "Car not found", http.StatusNotFound)
}

// Deleteing the Car Details
func deleteCar(w http.ResponseWriter, r *http.Request) {
	idParam := r.URL.Path[len("/cars/"):] // Extract the ID from URL path
	id, err := strconv.Atoi(idParam)
	if err != nil {
		http.Error(w, "Invalid car ID", http.StatusBadRequest)
		return
	}

	for index, car := range cars {
		if car.ID == id {
			cars = append(cars[:index], cars[index+1:]...)
			w.WriteHeader(http.StatusNoContent)
			return
		}
	}
	http.Error(w, "Car not found", http.StatusNotFound)
}

func main() {
	http.HandleFunc("/cars", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			getCars(w, r)
		case "POST":
			createCar(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	http.HandleFunc("/cars/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			getCar(w, r)
		case "PUT":
			updateCar(w, r)
		case "DELETE":
			deleteCar(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	fmt.Println("Server running on port 8497")
	http.ListenAndServe(":8497", nil)
}
