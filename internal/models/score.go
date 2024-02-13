package models

type Score struct {
	Id                int `json:"id,omitempty"`
	HomeTeamScore     int `json:"home_team_score,omitempty"`
	OpposingTeamScore int `json:"opposing_team_score,omitempty"`
}
