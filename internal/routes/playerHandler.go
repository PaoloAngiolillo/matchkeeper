package routes

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

type PlayerHandler struct{}

func (p PlayerHandler) Routes() chi.Router {
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

func (p PlayerHandler) List(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("List Player"))
}

func (p PlayerHandler) Get(w http.ResponseWriter, r *http.Request) {
	// id := chi.URLParam(r, "id")
	w.Write([]byte("Get Player "))
}

func (p PlayerHandler) Create(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Create Player"))
}

func (p PlayerHandler) Update(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Update Player"))
}

func (p PlayerHandler) Delete(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Delete Player"))
}
