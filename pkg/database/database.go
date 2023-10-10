package database

import (
	"database/sql"
)

type Service interface {
	Connect() (*sql.DB, error)
}
