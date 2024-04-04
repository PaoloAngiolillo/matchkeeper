package server

import (
	"encoding/json"
	"github.com/go-chi/render"
	"log"
	"matchkeeper/internal/repository"
	"net/http"

	"matchkeeper/internal/services"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func (s *Server) RegisterRoutes() http.Handler {

	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.URLFormat)
	r.Use(render.SetContentType(render.ContentTypeJSON))

	r.Get("/", s.HelloWorldHandler)
	r.Get("/health", s.healthHandler)
	r.Mount("/", s.mountSubroutes())
	return r
}

func (s *Server) mountSubroutes() chi.Router {
	r := chi.NewRouter()
	sqliteMatchRepository := repository.NewMySqliteMatchRepository(&s.db)
	sqlitePlayerRepository := repository.NewMySqlitePlayerRepository(&s.db)
	matchService := services.NewMatchService(sqliteMatchRepository)
	playerService := services.NewPlayerService(sqlitePlayerRepository)

	teamService := services.TeamService{}
	r.Mount("/match", matchService.Routes())
	r.Mount("/team", teamService.Routes())
	r.Mount("/player", playerService.Routes())

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
