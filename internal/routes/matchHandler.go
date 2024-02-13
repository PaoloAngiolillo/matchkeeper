package routes

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	"matchkeeper/internal/models"
	"matchkeeper/internal/repository"
	"net/http"
	"strconv"
)

func InternalServerErrorHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte("500 Internal Server Error"))
}

func NotFoundHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte("404 Not Found"))
}

type MatchHandler struct {
	repository repository.MatchRepository
}

func (h *MatchHandler) Routes() chi.Router {
	inMemRepository := repository.NewInMemRepository()
	sqlLiteMatchRepository := repository.MySqliteRepository()
	matchHandler := NewMatchHandler(inMemRepository)

	router := chi.NewRouter()
	router.Use(middleware.RequestID)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Use(middleware.URLFormat)
	router.Use(render.SetContentType(render.ContentTypeJSON))
	router.Get("/", matchHandler.List)
	router.Post("/", matchHandler.Create)
	router.Put("/", matchHandler.Update)

	router.Route("/{id}", func(r chi.Router) {
		r.Get("/", matchHandler.Get)
		r.Put("/", matchHandler.Update)
		r.Delete("/", matchHandler.Delete)
	})
	return router
}

func NewMatchHandler(s repository.MatchRepository) *MatchHandler {
	return &MatchHandler{
		repository: s,
	}
}

func (h *MatchHandler) List(w http.ResponseWriter, r *http.Request) {
	matches, err := h.repository.List()
	fmt.Println(matches)
	jsonBytes, err := json.Marshal(matches)
	if err != nil {
		InternalServerErrorHandler(w, r)
		return
	}

	w.WriteHeader(http.StatusOK)
	//w.Write([]byte("List Matches"))
	w.Write(jsonBytes)
}

func (h *MatchHandler) Create(w http.ResponseWriter, r *http.Request) {
	// Match object that will be populated from json payload
	var match models.Match

	if err := json.NewDecoder(r.Body).Decode(&match); err != nil {
		InternalServerErrorHandler(w, r)
		return
	}

	if err := h.repository.Create(match); err != nil {
		InternalServerErrorHandler(w, r)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Create Match"))
	fmt.Println(match)

}

func (h *MatchHandler) Get(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	convertedId, err := strconv.Atoi(id)
	if err != nil {
		NotFoundHandler(w, r)
		return
	}

	match, err := h.repository.Get(convertedId)
	if err != nil {
		if errors.Is(err, repository.NotFoundErr) {
			NotFoundHandler(w, r)
			return
		}

		InternalServerErrorHandler(w, r)
		return
	}

	jsonBytes, err := json.Marshal(match)
	if err != nil {
		InternalServerErrorHandler(w, r)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(jsonBytes)
	w.Write([]byte("Get Match"))
}

func (h *MatchHandler) Update(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Update Match"))
}

func (h *MatchHandler) Delete(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Delete Match"))
}
