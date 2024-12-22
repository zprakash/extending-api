package dht22

import (
	"context"
	"goapi/internal/api/repository/models"
)

type MockDHT22ServiceSuccessful struct{}

func (m *MockDHT22ServiceSuccessful) Create(data *models.DHT22Data, ctx context.Context) error {
	return nil
}

func (m *MockDHT22ServiceSuccessful) ReadOne(id int, ctx context.Context) (*models.DHT22Data, error) {
	return &models.DHT22Data{
		ID:          1,
		DeviceName:  "DHT22 Sensor",
		Temperature: 10,
		Humidity:    23.0,
		DateTime:    "2024-12-22T12:00:00Z",
	}, nil
}

func (m *MockDHT22ServiceSuccessful) ReadMany(page, rowsPerPage int, ctx context.Context) ([]*models.DHT22Data, error) {
	return []*models.DHT22Data{
		{
			ID:          1,
			DeviceName:  "DHT22 Sensor 1",
			Temperature: 22.5,
			Humidity:    50.0,
			DateTime:    "2024-12-22T10:00:00Z",
		},
		{
			ID:          2,
			DeviceName:  "DHT22 Sensor 2",
			Temperature: 24.3,
			Humidity:    55.0,
			DateTime:    "2024-12-22T11:00:00Z",
		},
	}, nil
}

func (m *MockDHT22ServiceSuccessful) Update(data *models.DHT22Data, ctx context.Context) error {
	return nil
}

func (m *MockDHT22ServiceSuccessful) Delete(data *models.DHT22Data, ctx context.Context) error {
	return nil
}

// MockDHT22ServiceNotFound: Simulates not found responses
type MockDHT22ServiceNotFound struct{}

func (m *MockDHT22ServiceNotFound) Create(data *models.DHT22Data, ctx context.Context) error {
	return nil
}

func (m *MockDHT22ServiceNotFound) ReadOne(id int, ctx context.Context) (*models.DHT22Data, error) {
	return nil, nil
}

func (m *MockDHT22ServiceNotFound) ReadMany(page, rowsPerPage int, ctx context.Context) ([]*models.DHT22Data, error) {
	return []*models.DHT22Data{}, nil
}

func (m *MockDHT22ServiceNotFound) Update(data *models.DHT22Data, ctx context.Context) error {
	return nil
}

func (m *MockDHT22ServiceNotFound) Delete(data *models.DHT22Data, ctx context.Context) error {
	return nil
}

// MockDHT22ServiceError: Simulates error responses
type MockDHT22ServiceError struct{}

func (m *MockDHT22ServiceError) Create(data *models.DHT22Data, ctx context.Context) error {
	return DHT22Error("Error creating DHT22 data")
}

func (m *MockDHT22ServiceError) ReadOne(id int, ctx context.Context) (*models.DHT22Data, error) {
	return nil, DHT22Error("Error reading DHT22 data")
}

func (m *MockDHT22ServiceError) ReadMany(page, rowsPerPage int, ctx context.Context) ([]*models.DHT22Data, error) {
	return nil, DHT22Error("Error reading multiple DHT22 data entries")
}

func (m *MockDHT22ServiceError) Update(data *models.DHT22Data, ctx context.Context) error {
	return DHT22Error("Error updating DHT22 data")
}

func (m *MockDHT22ServiceError) Delete(data *models.DHT22Data, ctx context.Context) error {
	return DHT22Error("Error deleting DHT22 data")
}
