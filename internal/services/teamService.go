package services

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

type TeamService struct{}

func (t TeamService) Routes() chi.Router {
	router := chi.NewRouter()

	router.Get("/", t.List)
	router.Post("/", t.Create)
	router.Put("/", t.Update)

	router.Route("/{id}", func(r chi.Router) {
		r.Get("/", t.Get)
		r.Put("/", t.Update)
		r.Delete("/", t.Delete)
	})
	return router
}

func (t TeamService) List(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("List Teams"))
}

func (t TeamService) Get(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Get Teams"))
}

func (t TeamService) Create(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Create Teams"))
}

func (t TeamService) Update(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Update Teams"))
}

func (t TeamService) Delete(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Delete Match"))
}
