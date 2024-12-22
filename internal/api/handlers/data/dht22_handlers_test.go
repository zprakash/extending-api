package data

import (
	"bytes"
	"encoding/json"
	"goapi/internal/api/repository/models"
	"goapi/internal/api/service/dht22"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCreateDHT22Handler_Success(t *testing.T) {
	mockService := &dht22.MockDHT22ServiceSuccessful{}
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		CreateDHT22Handler(w, r, nil, mockService)
	})

	// Prepare a test DHT22Data object
	data := &models.DHT22Data{
		DeviceName:  "Test Sensor",
		Temperature: 25.5,
		Humidity:    60.0,
		DateTime:    "2024-12-22T12:00:00Z",
	}

	// Encode the data into JSON
	reqBody, _ := json.Marshal(data)
	req := httptest.NewRequest("POST", "/dht22", bytes.NewReader(reqBody))
	w := httptest.NewRecorder()

	// Call the handler
	handler.ServeHTTP(w, req)

	// Check the response code
	if w.Code != http.StatusCreated {
		t.Errorf("Expected status code %d, got %d", http.StatusCreated, w.Code)
	}

	// Check that the response body is the same as the input data
	var respData models.DHT22Data
	if err := json.NewDecoder(w.Body).Decode(&respData); err != nil {
		t.Fatalf("Failed to decode response body: %v", err)
	}

	if respData.DeviceName != data.DeviceName || respData.Temperature != data.Temperature || respData.Humidity != data.Humidity {
		t.Errorf("Expected response body %+v, got %+v", data, respData)
	}
}

func TestCreateDHT22Handler_Error(t *testing.T) {
	mockService := &dht22.MockDHT22ServiceError{}
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		CreateDHT22Handler(w, r, nil, mockService)
	})

	// Prepare test data
	data := &models.DHT22Data{
		DeviceName:  "Error Sensor",
		Temperature: 25.0,
		Humidity:    55.0,
		DateTime:    "2024-12-22T12:00:00Z",
	}

	reqBody, _ := json.Marshal(data)
	req := httptest.NewRequest("POST", "/dht22", bytes.NewReader(reqBody))
	w := httptest.NewRecorder()

	// Call the handler
	handler.ServeHTTP(w, req)

	// Check the response code for error
	if w.Code != http.StatusInternalServerError {
		t.Errorf("Expected status code %d, got %d", http.StatusInternalServerError, w.Code)
	}
}

func TestGetDHT22Handler_Success(t *testing.T) {
	mockService := &dht22.MockDHT22ServiceSuccessful{}
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		GetDHT22Handler(w, r, nil, mockService)
	})

	req := httptest.NewRequest("GET", "/dht22", nil)
	w := httptest.NewRecorder()

	// Call the handler
	handler.ServeHTTP(w, req)

	// Check the response code
	if w.Code != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, w.Code)
	}

	// Check that the response body contains the expected data
	var respData []*models.DHT22Data
	if err := json.NewDecoder(w.Body).Decode(&respData); err != nil {
		t.Fatalf("Failed to decode response body: %v", err)
	}

	if len(respData) != 2 {
		t.Errorf("Expected 2 records, got %d", len(respData))
	}
}

func TestGetDHT22ByIDHandler_NotFound(t *testing.T) {
	mockService := &dht22.MockDHT22ServiceNotFound{}
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		GetDHT22ByIDHandler(w, r, nil, mockService)
	})

	req := httptest.NewRequest("GET", "/dht22/999", nil) // Non-existent ID
	w := httptest.NewRecorder()

	// Call the handler
	handler.ServeHTTP(w, req)

	// Check that we get a 404 Not Found
	if w.Code != http.StatusNotFound {
		t.Errorf("Expected status code %d, got %d", http.StatusNotFound, w.Code)
	}
}

func TestGetDHT22ByIDHandler_Success(t *testing.T) {
	mockService := &dht22.MockDHT22ServiceSuccessful{}
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		GetDHT22ByIDHandler(w, r, nil, mockService)
	})

	req := httptest.NewRequest("GET", "/dht22/1", nil)
	w := httptest.NewRecorder()

	// Call the handler
	handler.ServeHTTP(w, req)

	// Check the response code
	if w.Code != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, w.Code)
	}

	// Check that the response body contains the expected data
	var respData models.DHT22Data
	if err := json.NewDecoder(w.Body).Decode(&respData); err != nil {
		t.Fatalf("Failed to decode response body: %v", err)
	}

	if respData.ID != 1 {
		t.Errorf("Expected ID 1, got %d", respData.ID)
	}
}

func TestUpdateDHT22Handler_Success(t *testing.T) {
	mockService := &dht22.MockDHT22ServiceSuccessful{}
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		UpdateDHT22Handler(w, r, nil, mockService)
	})

	// Prepare test data
	data := &models.DHT22Data{
		DeviceName:  "Updated Sensor",
		Temperature: 30.0,
		Humidity:    65.0,
		DateTime:    "2024-12-22T12:00:00Z",
	}

	reqBody, _ := json.Marshal(data)
	req := httptest.NewRequest("PUT", "/dht22/1", bytes.NewReader(reqBody))
	w := httptest.NewRecorder()

	// Call the handler
	handler.ServeHTTP(w, req)

	// Check the response code
	if w.Code != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, w.Code)
	}

	// Check that the response body contains a success message
	if w.Body.String() != "DHT22 data updated successfully" {
		t.Errorf("Expected success message, got %s", w.Body.String())
	}
}

func TestDeleteDHT22Handler_Success(t *testing.T) {
	mockService := &dht22.MockDHT22ServiceSuccessful{}
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		DeleteDHT22Handler(w, r, nil, mockService)
	})

	req := httptest.NewRequest("DELETE", "/dht22/1", nil)
	w := httptest.NewRecorder()

	// Call the handler
	handler.ServeHTTP(w, req)

	// Check the response code
	if w.Code != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, w.Code)
	}

	// Check that the response body contains a success message
	if w.Body.String() != "DHT22 data deleted successfully" {
		t.Errorf("Expected success message, got %s", w.Body.String())
	}
}
