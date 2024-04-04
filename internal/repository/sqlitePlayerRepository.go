package repository

import (
	"context"
	"database/sql"
	"log"
	"matchkeeper/internal/database"
	"matchkeeper/internal/models"
)

type MySqlitePlayerRepository struct {
	db database.Service
}

func NewMySqlitePlayerRepository(db *database.Service) *MySqlitePlayerRepository {
	if db == nil {
		panic("missing db")
	}
	return &MySqlitePlayerRepository{db: *db}
}

func (m MySqlitePlayerRepository) List(ctx context.Context) (map[int]models.Player, error) {
	log.Println("Listing Players")
	query := `SELECT * FROM players`
	rows, err := m.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {

		}
	}(rows)

	players := make(map[int]models.Player)
	for rows.Next() {
		var player models.Player
		if err := rows.Scan(&player.Id, &player.FirstName, &player.LastName); err != nil {
			return nil, err
		}
		players[player.Id] = player
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}
	return players, nil
}

func (m MySqlitePlayerRepository) GetById(ctx context.Context, id int) (models.Player, error) {
	query := `SELECT * FROM players WHERE id = ?`
	var player models.Player

	row := m.db.QueryRowContext(ctx, query, id)

	if err := row.Scan(&player.Id, &player.FirstName, &player.LastName); err != nil {
		return models.Player{}, err
	}
	return models.Player{}, nil
}

func (m MySqlitePlayerRepository) Update(ctx context.Context, id int, player models.Player) interface{} {

	// Prepare the SQL UPDATE statement
	query := `UPDATE players SET first_name = ?, last_name = ? WHERE id = ?`

	// Execute the SQL UPDATE statement
	_, err := m.db.ExecContext(ctx, query, player.FirstName, player.LastName, id)
	if err != nil {
		return err
	}

	return nil
}

func (m MySqlitePlayerRepository) Delete(ctx context.Context, id int) interface{} {
	// Prepare the SQL DELETE statement
	query := `
        DELETE FROM players
        WHERE id = ?
    `

	// Execute the SQL DELETE statement
	_, err := m.db.ExecContext(ctx, query, id)
	if err != nil {
		log.Println(err)
		return err
	}
	// If no error occurred, return nil
	return nil
}

func (m MySqlitePlayerRepository) Create(ctx context.Context, player models.Player) interface{} {
	//// Prepare the SQL INSERT statements
	playerQuery := `INSERT INTO players (first_name, last_name) VALUES (?, ?)`
	//// Execute the SQL INSERT statement for home player one
	_, err := m.db.ExecContext(ctx, playerQuery, player.FirstName, player.LastName)
	if err != nil {
		return err
	}
	// If no error occurred, return nil
	return nil
}
