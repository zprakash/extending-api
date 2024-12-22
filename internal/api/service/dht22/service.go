package dht22

import (
	"context"
	"goapi/internal/api/repository/models"
)

// DHT22Service handles the business logic for DHT22Data operations
type DHT22Service interface {
	Create(data *models.DHT22Data, ctx context.Context) error
	ReadOne(id int, ctx context.Context) (*models.DHT22Data, error)
	ReadMany(page, rowsPerPage int, ctx context.Context) ([]*models.DHT22Data, error)
	Update(data *models.DHT22Data, ctx context.Context) error
	Delete(data *models.DHT22Data, ctx context.Context) error
}

type DHT22Error string

func (e DHT22Error) Error() string {
	return string(e)
}

// dht22Service implements the DHT22Service interface
type dht22Service struct {
	repository models.DHT22Repository
}

func NewDHT22Service(repository models.DHT22Repository) DHT22Service {
	return &dht22Service{
		repository: repository,
	}
}

func (s *dht22Service) Create(data *models.DHT22Data, ctx context.Context) error {
	// Call repository to create data
	if err := s.repository.Create(data, ctx); err != nil {
		return err
	}
	return nil
}

func (s *dht22Service) ReadOne(id int, ctx context.Context) (*models.DHT22Data, error) {
	// Call repository to fetch data by ID
	data, err := s.repository.ReadOne(id, ctx)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (s *dht22Service) ReadMany(page, rowsPerPage int, ctx context.Context) ([]*models.DHT22Data, error) {
	// Call repository to fetch multiple records
	data, err := s.repository.ReadMany(page, rowsPerPage, ctx)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (s *dht22Service) Update(data *models.DHT22Data, ctx context.Context) error {
	// Call repository to update data
	_, err := s.repository.Update(data, ctx)
	if err != nil {
		return err
	}
	return nil
}

func (s *dht22Service) Delete(data *models.DHT22Data, ctx context.Context) error {
	// Call repository to delete data using the DHT22Data object
	_, err := s.repository.Delete(data, ctx)
	if err != nil {
		return err
	}
	return nil
}
