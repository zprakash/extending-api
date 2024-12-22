package data

import (
	"encoding/json"
	"fmt"
	"goapi/internal/api/repository/models"
	"goapi/internal/api/service/dht22"
	"log"
	"net/http"
	"strings"
)

// PostHandler - Creates a new DHT22 record
func CreateDHT22Handler(w http.ResponseWriter, r *http.Request, logger *log.Logger, dht22Service dht22.DHT22Service) {
	var data models.DHT22Data
	// Decode the incoming request body to the DHT22Data struct
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		http.Error(w, fmt.Sprintf("Invalid request body: %v", err), http.StatusBadRequest)
		return
	}

	// Call the service to create the record
	err := dht22Service.Create(&data, r.Context())
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to create DHT22 data: %v", err), http.StatusInternalServerError)
		return
	}

	// Respond with the created data in JSON format
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	if err := json.NewEncoder(w).Encode(data); err != nil {
		http.Error(w, fmt.Sprintf("Failed to encode response: %v", err), http.StatusInternalServerError)
	}
}

// GetHandler - Fetches all DHT22 records with pagination
func GetDHT22Handler(w http.ResponseWriter, r *http.Request, logger *log.Logger, dht22Service dht22.DHT22Service) {
	// For now, let's just assume page and rowsPerPage are passed as query parameters
	page := 1
	rowsPerPage := 10

	data, err := dht22Service.ReadMany(page, rowsPerPage, r.Context())
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to fetch DHT22 data: %v", err), http.StatusInternalServerError)
		return
	}

	// Respond with the fetched data
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		http.Error(w, fmt.Sprintf("Failed to encode response: %v", err), http.StatusInternalServerError)
	}
}

// GetByIDHandler - Fetches a DHT22 record by ID
func GetDHT22ByIDHandler(w http.ResponseWriter, r *http.Request, logger *log.Logger, dht22Service dht22.DHT22Service) {
	// Extract the ID from the URL
	idStr := strings.TrimPrefix(r.URL.Path, "/dht22/")
	var id int
	if _, err := fmt.Sscanf(idStr, "%d", &id); err != nil {
		http.Error(w, "Invalid ID format", http.StatusBadRequest)
		return
	}

	// Fetch the DHT22 record by ID
	data, err := dht22Service.ReadOne(id, r.Context())
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to fetch DHT22 data: %v", err), http.StatusInternalServerError)
		return
	}

	// Respond with the fetched data
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		http.Error(w, fmt.Sprintf("Failed to encode response: %v", err), http.StatusInternalServerError)
	}
}

// PutHandler - Updates a DHT22 record by ID
func UpdateDHT22Handler(w http.ResponseWriter, r *http.Request, logger *log.Logger, dht22Service dht22.DHT22Service) {
	// Extract the ID from the URL
	idStr := strings.TrimPrefix(r.URL.Path, "/dht22/")
	var id int
	if _, err := fmt.Sscanf(idStr, "%d", &id); err != nil {
		http.Error(w, "Invalid ID format", http.StatusBadRequest)
		return
	}

	var data models.DHT22Data
	// Decode the incoming request body to the DHT22Data struct
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		http.Error(w, fmt.Sprintf("Invalid request body: %v", err), http.StatusBadRequest)
		return
	}

	// Set the ID on the data (for update)
	data.ID = id

	// Call the service to update the record
	if err := dht22Service.Update(&data, r.Context()); err != nil {
		http.Error(w, fmt.Sprintf("Failed to update DHT22 data: %v", err), http.StatusInternalServerError)
		return
	}

	// Respond with a success message
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("DHT22 data updated successfully"))
}

// DeleteHandler - Deletes a DHT22 record by ID
func DeleteDHT22Handler(w http.ResponseWriter, r *http.Request, logger *log.Logger, dht22Service dht22.DHT22Service) {
	// Extract the ID from the URL
	idStr := strings.TrimPrefix(r.URL.Path, "/dht22/")
	var id int
	if _, err := fmt.Sscanf(idStr, "%d", &id); err != nil {
		http.Error(w, "Invalid ID format", http.StatusBadRequest)
		return
	}

	// Create a DHT22Data object with the ID to pass to the service
	data := &models.DHT22Data{
		ID: id,
	}

	// Call the service to delete the record using the DHT22Data object
	if err := dht22Service.Delete(data, r.Context()); err != nil {
		http.Error(w, fmt.Sprintf("Failed to delete DHT22 data: %v", err), http.StatusInternalServerError)
		return
	}

	// Respond with a success message
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("DHT22 data deleted successfully"))
}
