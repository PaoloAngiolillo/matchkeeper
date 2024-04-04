package services

import (
	"encoding/json"
	"fmt"
	"matchkeeper/internal/models"
	"matchkeeper/internal/repository"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type PlayerService struct {
	repository repository.PlayerRepository
}

func NewPlayerService(pr *repository.MySqlitePlayerRepository) *PlayerService {
	return &PlayerService{
		repository: pr,
	}
}

func (p PlayerService) Routes() chi.Router {
	router := chi.NewRouter()

	router.Get("/", p.List)
	router.Post("/", p.Create)

	router.Route("/{id}", func(r chi.Router) {
		r.Get("/", p.GetById)
		r.Put("/", p.Update)
		r.Delete("/", p.Delete)
	})
	return router
}

func (p PlayerService) List(w http.ResponseWriter, r *http.Request) {
	players, err := p.repository.List(r.Context())
	fmt.Println(players)
	jsonBytes, err := json.Marshal(players)
	if err != nil {
		InternalServerErrorHandler(w, r)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(jsonBytes)
	w.Write([]byte("List Player"))
}

func (p PlayerService) GetById(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	convertedId, err := strconv.Atoi(id)

	player, err := p.repository.GetById(r.Context(), convertedId)
	if err != nil {
		InternalServerErrorHandler(w, r)
		return
	}
	jsonBytes, err := json.Marshal(player)
	if err != nil {
		InternalServerErrorHandler(w, r)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(jsonBytes)
	w.Write([]byte("Get Player "))
}

func (p PlayerService) Create(w http.ResponseWriter, r *http.Request) {
	var player models.Player
	err := json.NewDecoder(r.Body).Decode(&player)
	if err != nil {
		InternalServerErrorHandler(w, r)
		return
	}
	if err := p.repository.Create(r.Context(), player); err != nil {
		InternalServerErrorHandler(w, r)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Create Player"))
}

func (p PlayerService) Update(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	convertedId, err := strconv.Atoi(id)
	if err != nil {
		InternalServerErrorHandler(w, r)
		return
	}
	var player models.Player
	err = json.NewDecoder(r.Body).Decode(&player)
	if err != nil {
		InternalServerErrorHandler(w, r)
		return
	}
	if err := p.repository.Update(r.Context(), convertedId, player); err != nil {
		InternalServerErrorHandler(w, r)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Println(player)
	w.Write([]byte("Update Player"))
}

func (p PlayerService) Delete(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	convertedId, err := strconv.Atoi(id)
	if err != nil {
		InternalServerErrorHandler(w, r)
		return
	}
	if err := p.repository.Delete(r.Context(), convertedId); err != nil {
		InternalServerErrorHandler(w, r)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Delete Player"))
}
