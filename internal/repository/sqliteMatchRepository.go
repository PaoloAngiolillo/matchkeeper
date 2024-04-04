package repository

import (
	"context"
	"log"
	"matchkeeper/internal/database"
	"matchkeeper/internal/models"
)

type MySqliteMatchRepository struct {
	db database.Service
}

func NewMySqliteMatchRepository(db *database.Service) *MySqliteMatchRepository {
	if db == nil {
		panic("missing db")
	}
	return &MySqliteMatchRepository{db: *db}
}

func (m MySqliteMatchRepository) List(ctx context.Context) (map[int]models.Match, error) {
	log.Println("Listing Matches")
	query := `
	SELECT m.id, 
		p1.first_name || ' ' || p1.last_name as home_player_one,
		p2.first_name || ' ' || p2.last_name as home_player_two,
		p3.first_name || ' ' || p3.last_name as opposing_player_one,
		p4.first_name || ' ' || p4.last_name as opposing_player_two,
		s.home_team_score, s.opposing_team_score, m.created_date
	FROM matches m
	INNER JOIN teams t1 ON m.home_team_id = t1.id
	INNER JOIN teams t2 ON m.opposing_team_id = t2.id
	INNER JOIN players p1 ON t1.player_one_id = p1.id
	INNER JOIN players p2 ON t1.player_two_id = p2.id
	INNER JOIN players p3 ON t2.player_one_id = p3.id
	INNER JOIN players p4 ON t2.player_two_id = p4.id
	INNER JOIN scores s ON m.score_id = s.id
	`
	rows, err := m.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	matches := make(map[int]models.Match)
	for rows.Next() {
		var match models.Match
		var homePlayerOne, homePlayerTwo, opposingPlayerOne, opposingPlayerTwo string
		var homeTeamScore, opposingTeamScore int
		if err := rows.Scan(&match.Id, &homePlayerOne, &homePlayerTwo, &opposingPlayerOne, &opposingPlayerTwo, &homeTeamScore, &opposingTeamScore, &match.CreatedDate); err != nil {
			return nil, err
		}

		match.HomeTeam = models.Team{PlayerOne: models.Player{Name: homePlayerOne}, PlayerTwo: models.Player{Name: homePlayerTwo}}
		match.OpposingTeam = models.Team{PlayerOne: models.Player{Name: opposingPlayerOne}, PlayerTwo: models.Player{Name: opposingPlayerTwo}}
		match.Score = models.Score{HomeTeamScore: homeTeamScore, OpposingTeamScore: opposingTeamScore}

		matches[match.Id] = match
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}
	return matches, nil
}

func (m MySqliteMatchRepository) GetById(ctx context.Context, id int) (models.Match, error) {
	query := `
	SELECT m.id,
		p1.first_name || ' ' || p1.last_name as home_player_one,
		p2.first_name || ' ' || p2.last_name as home_player_two,
		p3.first_name || ' ' || p3.last_name as opposing_player_one,
		p4.first_name || ' ' || p4.last_name as opposing_player_two,
		s.home_team_score, s.opposing_team_score, m.created_date
	FROM matches m
	INNER JOIN teams t1 ON m.home_team_id = t1.id
	INNER JOIN teams t2 ON m.opposing_team_id = t2.id
	INNER JOIN players p1 ON t1.player_one_id = p1.id
	INNER JOIN players p2 ON t1.player_two_id = p2.id
	INNER JOIN players p3 ON t2.player_one_id = p3.id
	INNER JOIN players p4 ON t2.player_two_id = p4.id
	INNER JOIN scores s ON m.score_id = s.id
	WHERE m.id = ?
	`
	row := m.db.QueryRowContext(ctx, query, id)

	var match models.Match
	var homePlayerOne, homePlayerTwo, opposingPlayerOne, opposingPlayerTwo string
	var homeTeamScore, opposingTeamScore int
	if err := row.Scan(&match.Id, &homePlayerOne, &homePlayerTwo, &opposingPlayerOne, &opposingPlayerTwo, &homeTeamScore, &opposingTeamScore, &match.CreatedDate); err != nil {
		return models.Match{}, err
	}

	match.HomeTeam = models.Team{PlayerOne: models.Player{Name: homePlayerOne}, PlayerTwo: models.Player{Name: homePlayerTwo}}
	match.OpposingTeam = models.Team{PlayerOne: models.Player{Name: opposingPlayerOne}, PlayerTwo: models.Player{Name: opposingPlayerTwo}}
	match.Score = models.Score{HomeTeamScore: homeTeamScore, OpposingTeamScore: opposingTeamScore}

	return match, nil
}

func (m MySqliteMatchRepository) Update(ctx context.Context, id int, match models.Match) interface{} {

	// Prepare the SQL UPDATE statement
	matchUpdateQuery := `UPDATE matches SET home_team_id = ?, opposing_team_id = ?, score_id = ? WHERE id = ?`
	_, err := m.db.ExecContext(ctx, matchUpdateQuery, match.HomeTeam.Id, match.OpposingTeam.Id, match.Score.Id, id)
	if err != nil {
		return err
	}

	return nil
}

func (m MySqliteMatchRepository) Delete(ctx context.Context, id int) interface{} {
	// Prepare the SQL DELETE statement
	query := `
        DELETE FROM matches
        WHERE id = ?
    `

	// Execute the SQL DELETE statement
	_, err := m.db.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}

	// If no error occurred, return nil
	return nil
}

func (m MySqliteMatchRepository) Create(ctx context.Context, match models.Match) interface{} {
	// Prepare the SQL INSERT statements
	playerQuery := `INSERT INTO players (first_name, last_name) VALUES (?, ?)`
	teamQuery := `INSERT INTO teams (player_one_id, player_two_id) VALUES (?, ?)`
	scoreQuery := `INSERT INTO scores (home_team_score, opposing_team_score) VALUES (?, ?)`
	matchQuery := `INSERT INTO matches (home_team_id, opposing_team_id, score_id) VALUES (?, ?, ?)`
	idQuery := `SELECT last_insert_rowid()`

	// Execute the SQL INSERT statement for home player one
	_, err := m.db.ExecContext(ctx, playerQuery, match.HomeTeam.PlayerOne.FirstName, match.HomeTeam.PlayerOne.LastName)
	if err != nil {
		return err
	}
	var homePlayerOneId int
	err = m.db.QueryRowContext(ctx, idQuery).Scan(&homePlayerOneId)
	if err != nil {
		return err
	}

	// Execute the SQL INSERT statement for home player two
	_, err = m.db.ExecContext(ctx, playerQuery, match.HomeTeam.PlayerTwo.FirstName, match.HomeTeam.PlayerTwo.LastName)
	if err != nil {
		return err
	}
	var homePlayerTwoId int
	err = m.db.QueryRowContext(ctx, idQuery).Scan(&homePlayerTwoId)
	if err != nil {
		return err
	}

	// Execute the SQL INSERT statement for home team
	_, err = m.db.ExecContext(ctx, teamQuery, homePlayerOneId, homePlayerTwoId)
	if err != nil {
		return err
	}

	var homeTeamId int
	err = m.db.QueryRowContext(ctx, idQuery).Scan(&homeTeamId)
	if err != nil {
		return err
	}

	// Execute the SQL INSERT statement for opposing team
	_, err = m.db.ExecContext(ctx, teamQuery, match.OpposingTeam.PlayerOne.Id, match.OpposingTeam.PlayerTwo.Id)
	if err != nil {
		return err
	}
	var opposingTeamId int
	err = m.db.QueryRowContext(ctx, idQuery).Scan(&opposingTeamId)
	if err != nil {
		return err
	}

	// Execute the SQL INSERT statement for score
	_, err = m.db.ExecContext(ctx, scoreQuery, match.Score.HomeTeamScore, match.Score.OpposingTeamScore)
	if err != nil {
		return err
	}
	var scoreId int
	err = m.db.QueryRowContext(ctx, idQuery).Scan(&scoreId)
	if err != nil {
		return err
	}

	// Execute the SQL INSERT statement for match
	_, err = m.db.ExecContext(ctx, matchQuery, homeTeamId, opposingTeamId, scoreId)
	if err != nil {
		return err
	}

	// If no error occurred, return nil
	return nil
}
