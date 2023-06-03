package database

import (
	"database/sql"
)

type DatabaseService interface {
	Connect() (*sql.DB, error)
}
