package models

import "context"

type DHT22Data struct {
	ID          int     `json:"id"`
	DeviceName  string  `json:"device_name"`
	Temperature float64 `json:"temperature"`
	Humidity    float64 `json:"humidity"`
	DateTime    string  `json:"date_time"`
}

type DHT22Repository interface {
	Create(data *DHT22Data, ctx context.Context) error
	ReadOne(id int, ctx context.Context) (*DHT22Data, error)
	ReadMany(page int, rowsPerPage int, ctx context.Context) ([]*DHT22Data, error)
	Update(data *DHT22Data, ctx context.Context) (int64, error)
	Delete(data *DHT22Data, ctx context.Context) (int64, error)
}
