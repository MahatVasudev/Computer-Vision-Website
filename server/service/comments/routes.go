package comments

import (
	"github.com/MahatVasudev/Computer-Vision-Website/server/store"
	"github.com/go-chi/chi/v5"
)

type Handler struct {
	store     store.CommentStore
	userStore store.UserStore
	postStore store.PostStore
}

func NewHandler(
	store store.CommentStore,
	userStore store.UserStore,
	postStore store.PostStore,
) *Handler {
	return &Handler{
		store,
		userStore,
		postStore,
	}
}

func (h *Handler) RegisteredRoutes(route_name string, router *chi.Mux) {
	router.Route(route_name, func(r chi.Router) {

	})
}
