package models

import "context"

type Data struct {
	ID         int    `json:"id"`
	DeviceID   string `json:"device_id"`
	DeviceName string `json:"device_name"`
	//removed value and added Price and Serial Number
	Price        float64 `json:"price"`
	SerialNumber float64 `json:"SerialNumber"`
	Type         string  `json:"type"`
	DateTime     string  `json:"date_time"`
	Description  string  `json:"description"`
}

type DataRepository interface {
	Create(Data *Data, ctx context.Context) error
	ReadOne(id int, ctx context.Context) (*Data, error)
	ReadMany(page int, rowsPerPage int, ctx context.Context) ([]*Data, error)
	Update(data *Data, ctx context.Context) (int64, error)
	Delete(data *Data, ctx context.Context) (int64, error)
}

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
	Delete(id int, ctx context.Context) (int64, error)
}
