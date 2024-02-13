package routes

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

type TeamHandler struct{}

func (t TeamHandler) Routes() chi.Router {
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

func (t TeamHandler) List(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("List Teams"))
}

func (t TeamHandler) Get(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Get Teams"))
}

func (t TeamHandler) Create(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Create Teams"))
}

func (t TeamHandler) Update(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Update Teams"))
}

func (t TeamHandler) Delete(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Delete Match"))
}
