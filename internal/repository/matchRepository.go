package repository

import (
	"context"
	"matchkeeper/internal/models"
)

type MatchRepository interface {
	List(ctx context.Context) (map[int]models.Match, error)
	GetById(ctx context.Context, id int) (models.Match, error)
	Update(ctx context.Context, id int, match models.Match) interface{}
	Delete(ctx context.Context, id int) interface{}
	Create(ctx context.Context, match models.Match) interface{}
}

