package services

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

type PlayerService struct{}

func (p PlayerService) Routes() chi.Router {
	router := chi.NewRouter()

	router.Get("/", p.List)
	router.Post("/", p.Create)

	router.Route("/{id}", func(r chi.Router) {
		r.Get("/", p.Get)
		r.Put("/", p.Update)
		r.Delete("/", p.Delete)
	})
	return router
}

func (p PlayerService) List(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("List Player"))
}

func (p PlayerService) Get(w http.ResponseWriter, r *http.Request) {
	// id := chi.URLParam(r, "id")
	w.Write([]byte("Get Player "))
}

func (p PlayerService) Create(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Create Player"))
}

func (p PlayerService) Update(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Update Player"))
}

func (p PlayerService) Delete(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Delete Player"))
}
