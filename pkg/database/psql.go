package database

import (
	"database/sql"
	_ "github.com/lib/pq"
)

type PostgreSQLService struct {
	connectionString string
}

func NewPostgreSQLService(connectionString string) *PostgreSQLService {
	return &PostgreSQLService{
		connectionString,
	}
}

func (service *PostgreSQLService) Connect() (*sql.DB, error) {

	db, err := sql.Open("postgres", service.connectionString)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
