package services

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-chi/chi/v5"
	"matchkeeper/internal/database"
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

type MatchService struct {
	repository repository.MatchRepository
}

var db = database.New()

func (mh *MatchService) Routes() chi.Router {
	router := chi.NewRouter()

	router.Get("/", mh.List)
	router.Post("/", mh.Create)
	router.Put("/", mh.Update)

	router.Route("/{id}", func(r chi.Router) {
		r.Get("/", mh.Get)
		r.Put("/", mh.Update)
		r.Delete("/", mh.Delete)
	})
	return router
}

func NewMatchService(mr repository.MatchRepository) *MatchService {
	return &MatchService{
		repository: mr,
	}
}

func (mh *MatchService) List(w http.ResponseWriter, r *http.Request) {
	matches, err := mh.repository.List()
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

func (mh *MatchService) Create(w http.ResponseWriter, r *http.Request) {
	// Match object that will be populated from json payload
	var match models.Match

	if err := json.NewDecoder(r.Body).Decode(&match); err != nil {
		InternalServerErrorHandler(w, r)
		return
	}

	if err := mh.repository.Create(match); err != nil {
		InternalServerErrorHandler(w, r)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Create Match"))
	fmt.Println(match)

}

func (mh *MatchService) Get(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	convertedId, err := strconv.Atoi(id)
	if err != nil {
		NotFoundHandler(w, r)
		return
	}

	match, err := mh.repository.Get(convertedId)
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

func (mh *MatchService) Update(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Update Match"))
}

func (mh *MatchService) Delete(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Delete Match"))
}
