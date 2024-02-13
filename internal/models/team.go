package models

type Team struct {
	Id        int    `json:"id,omitempty"`
	PlayerOne Player `json:"player_one"`
	PlayerTwo Player `json:"player_two"`
}
