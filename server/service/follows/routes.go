package follows

import (
	"github.com/MahatVasudev/Computer-Vision-Website/server/middleware"
	"github.com/MahatVasudev/Computer-Vision-Website/server/store"
	"github.com/go-chi/chi/v5"
)

type Handler struct {
	store     store.FollowStore
	userStore store.UserStore
}

func NewHandler(followstore store.FollowStore, userstore store.UserStore) *Handler {
	return &Handler{
		store:     followstore,
		userStore: userstore,
	}
}

func (h *Handler) RegisterRoutes(route_string string, router *chi.Mux) {
	router.Route(route_string, func(r chi.Router) {

		// Get Function (get followers, followed etc...)

		r.With(middleware.AuthMiddleWareUser).Get("/{username}/followers", nil)

		r.With(middleware.AuthMiddleWareUser).Get("/{username}/following", nil)

		// Post Function (follow others)

		r.With(middleware.AuthMiddleWareUser).Post("/follow", h.HandleFollow)

		r.With(middleware.AuthMiddleWareUser).Delete("/{id}/follow", nil)
	})
}
