package repository

import (
	"context"
	"matchkeeper/internal/models"
)

type PlayerRepository interface {
	List(ctx context.Context) (map[int]models.Player, error)
	GetById(ctx context.Context, id int) (models.Player, error)
	Update(ctx context.Context, id int, player models.Player) interface{}
	Delete(ctx context.Context, id int) interface{}
	Create(ctx context.Context, match models.Player) interface{}
}
