package repository

import (
	"context"
	"matchkeeper/internal/database"
	"matchkeeper/internal/models"
)

//type mysqliteMatch struct {
//	ID           string    `db:"id"`
//	Hour         time.Time `db:"hour"`
//	Availability string    `db:"availability"`
//}

// sqlContextGetter is an interface provided both by transaction and standard db connection
type sqlContextGetter interface {
	GetContext(ctx context.Context, dest interface{}, query string, args ...interface{}) error
}

type MySqliteMatchRepository struct {
	db *database.Service
}

func NewMySQLMatchRepository(db *database.Service) *MySqliteMatchRepository {
	if db == nil {
		panic("missing db")
	}
	return &MySqliteMatchRepository{db: db}
}

func (m MySqliteMatchRepository) List() (map[int]models.Match, error) {
	//TODO implement me
	panic("implement me")
}

func (m MySqliteMatchRepository) Get(id int) (models.Match, error) {
	//TODO implement me
	panic("implement me")
}

func (m MySqliteMatchRepository) Update(id int, match models.Match) error {
	//TODO implement me
	panic("implement me")
}

func (m MySqliteMatchRepository) Delete(id int) error {
	//TODO implement me
	panic("implement me")
}

func (m MySqliteMatchRepository) Create(match models.Match) error {
	//TODO implement me
	panic("implement me")
}
