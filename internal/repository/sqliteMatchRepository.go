package repository

import (
	"context"
	"database/sql"
)

//type mysqliteMatch struct {
//	ID           string    `db:"id"`
//	Hour         time.Time `db:"hour"`
//	Availability string    `db:"availability"`
//}

type MySqliteRepository struct {
	db *sql.DB
}

func NewMySQLHourRepository(db *sql.DB) *MySqliteRepository {
	if db == nil {
		panic("missing db")
	}
	return &MySqliteRepository{db: db}
}

// sqlContextGetter is an interface provided both by transaction and standard db connection
type sqlContextGetter interface {
	GetContext(ctx context.Context, dest interface{}, query string, args ...interface{}) error
}
