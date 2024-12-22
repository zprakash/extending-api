package SQLite

import (
	"context"
	"database/sql"
	"goapi/internal/api/repository/DAL"
	"goapi/internal/api/repository/models"
)

type DHT22Repository struct {
	sqlDB *sql.DB
	createStmt,
	readStmt,
	readManyStmt,
	updateStmt,
	deleteStmt *sql.Stmt
	ctx context.Context
}

// NewDHT22Repository initializes the repository for DHT22Data.
func NewDHT22Repository(sqlDB DAL.SQLDatabase, ctx context.Context) (models.DHT22Repository, error) {

	repo := &DHT22Repository{
		sqlDB: sqlDB.Connection(),
		ctx:   ctx,
	}

	if _, err := repo.sqlDB.Exec(`DROP TABLE IF EXISTS dht22_data`); err != nil {
		repo.sqlDB.Close()
		return nil, err
	}

	// Create the `dht22_data` table
	if _, err := repo.sqlDB.Exec(`CREATE TABLE IF NOT EXISTS dht22_data (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		device_name VARCHAR(50) NOT NULL,
		temperature FLOAT NOT NULL,
		humidity FLOAT NOT NULL,
		date_time TIMESTAMP NOT NULL
	);`); err != nil {
		repo.sqlDB.Close()
		return nil, err
	}

	// Prepare SQL statements
	createStmt, err := repo.sqlDB.Prepare(`INSERT INTO dht22_data (device_name, temperature, humidity, date_time) VALUES (?, ?, ?, ?)`)
	if err != nil {
		repo.sqlDB.Close()
		return nil, err
	}
	repo.createStmt = createStmt

	readStmt, err := repo.sqlDB.Prepare("SELECT id, device_name, temperature, humidity, date_time FROM dht22_data WHERE id = ?")
	if err != nil {
		repo.sqlDB.Close()
		return nil, err
	}
	repo.readStmt = readStmt

	readManyStmt, err := repo.sqlDB.Prepare("SELECT id, device_name, temperature, humidity, date_time FROM dht22_data LIMIT ? OFFSET ?")
	if err != nil {
		repo.sqlDB.Close()
		return nil, err
	}
	repo.readManyStmt = readManyStmt

	updateStmt, err := repo.sqlDB.Prepare("UPDATE dht22_data SET device_name = ?, temperature = ?, humidity = ?, date_time = ? WHERE id = ?")
	if err != nil {
		repo.sqlDB.Close()
		return nil, err
	}
	repo.updateStmt = updateStmt

	deleteStmt, err := repo.sqlDB.Prepare("DELETE FROM dht22_data WHERE id = ?")
	if err != nil {
		repo.sqlDB.Close()
		return nil, err
	}
	repo.deleteStmt = deleteStmt

	// Handle cleanup when the context is canceled
	go CloseDHT22(ctx, repo)

	return repo, nil
}

// Cleanup resources when the context is canceled
func CloseDHT22(ctx context.Context, r *DHT22Repository) {
	<-ctx.Done()
	r.createStmt.Close()
	r.readStmt.Close()
	r.updateStmt.Close()
	r.deleteStmt.Close()
	r.readManyStmt.Close()
	r.sqlDB.Close()
}

// Implement CRUD operations

func (r *DHT22Repository) Create(data *models.DHT22Data, ctx context.Context) error {
	res, err := r.createStmt.ExecContext(ctx, data.DeviceName, data.Temperature, data.Humidity, data.DateTime)
	if err != nil {
		return err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return err
	}
	data.ID = int(id)
	return nil
}

func (r *DHT22Repository) ReadOne(id int, ctx context.Context) (*models.DHT22Data, error) {
	row := r.readStmt.QueryRowContext(ctx, id)
	var data models.DHT22Data
	err := row.Scan(&data.ID, &data.DeviceName, &data.Temperature, &data.Humidity, &data.DateTime)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &data, nil
}

func (r *DHT22Repository) ReadMany(page int, rowsPerPage int, ctx context.Context) ([]*models.DHT22Data, error) {
	offset := rowsPerPage * (page - 1)
	rows, err := r.readManyStmt.QueryContext(ctx, rowsPerPage, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var data []*models.DHT22Data
	for rows.Next() {
		var d models.DHT22Data
		err := rows.Scan(&d.ID, &d.DeviceName, &d.Temperature, &d.Humidity, &d.DateTime)
		if err != nil {
			return nil, err
		}
		data = append(data, &d)
	}
	return data, nil
}

func (r *DHT22Repository) Update(data *models.DHT22Data, ctx context.Context) (int64, error) {
	res, err := r.updateStmt.ExecContext(ctx, data.DeviceName, data.Temperature, data.Humidity, data.DateTime, data.ID)
	if err != nil {
		return 0, err
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return 0, err
	}
	return rowsAffected, nil
}

func (r *DHT22Repository) Delete(data *models.DHT22Data, ctx context.Context) (int64, error) {
	res, err := r.deleteStmt.ExecContext(ctx, data.ID)
	if err != nil {
		return 0, err
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return 0, err
	}
	return rowsAffected, nil
}
