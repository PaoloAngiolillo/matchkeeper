package repository

import (
	"matchkeeper/internal/models"
)

type MatchRepository interface {
	List() (map[int]models.Match, error)
	Get(id int) (models.Match, error)
	Update(id int, match models.Match) error
	Delete(id int) error
	Create(match models.Match) error
}
