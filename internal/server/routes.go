package server

import (
	"encoding/json"
	"log"
	"net/http"

	"matchkeeper/internal/routes"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func (s *Server) RegisterRoutes() http.Handler {
	matchHandler := routes.MatchHandler{}
	teamHandler := routes.TeamHandler{}
	playerHandler := routes.PlayerHandler{}
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Get("/", s.HelloWorldHandler)
	r.Get("/health", s.healthHandler)
	r.Mount("/match", matchHandler.Routes())
	r.Mount("/team", teamHandler.Routes())
	r.Mount("/player", playerHandler.Routes())
	return r
}

func (s *Server) HelloWorldHandler(w http.ResponseWriter, r *http.Request) {
	resp := make(map[string]string)
	resp["message"] = "Hello World"

	jsonResp, err := json.Marshal(resp)
	if err != nil {
		log.Fatalf("error handling JSON marshal. Err: %v", err)
	}

	_, _ = w.Write(jsonResp)
}

func (s *Server) healthHandler(w http.ResponseWriter, r *http.Request) {
	jsonResp, _ := json.Marshal(s.db.Health())
	_, _ = w.Write(jsonResp)
}
