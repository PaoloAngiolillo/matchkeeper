package models

import "time"

type Match struct {
	Id           int       `json:"id,omitempty"`
	HomeTeam     Team      `json:"home_team"`
	OpposingTeam Team      `json:"opposing_team"`
	Score        Score     `json:"score"`
	CreatedDate  time.Time `json:"created_date"`
}
