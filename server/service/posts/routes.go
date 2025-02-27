package posts

import (
	"github.com/MahatVasudev/Computer-Vision-Website/server/store"
	"github.com/go-chi/chi/v5"
)

type Handler struct {
	store     store.PostStore
	userStore store.UserStore
}

func NewHandler(postStore store.PostStore, userStore store.UserStore) *Handler {
	return &Handler{
		store:     postStore,
		userStore: userStore,
	}
}

func (h *Handler) RegisterRoutes(route_string string, router *chi.Mux) {
	router.Route(route_string, func(r chi.Router) {
		r.Get("/all", nil) // All Posts

		r.Get("/i/{id}", nil) // Get Specific Post

		r.Post("/make", nil) // create a post

		r.Get("/details", nil)
	})
}
