package posts

import (
	"github.com/MahatVasudev/Computer-Vision-Website/server/middleware"
	"github.com/MahatVasudev/Computer-Vision-Website/server/store"
	"github.com/go-chi/chi/v5"
)

type Handler struct {
	store        store.PostStore
	userStore    store.UserStore
	commentStore store.CommentStore
}

func NewHandler(
	postStore store.PostStore,
	userStore store.UserStore,
	commentStore store.CommentStore,
) *Handler {
	return &Handler{
		store:        postStore,
		userStore:    userStore,
		commentStore: commentStore,
	}
}

func (h *Handler) RegisterRoutes(route_string string, router *chi.Mux) {
	router.Route(route_string, func(r chi.Router) {
		r.Get("/all", h.handleGetAllPosts) // All Posts

		r.With(middleware.AuthMiddleWareUser).Post("/make", h.handleCreatePosts) // create a post

		r.Get("/i/{post_id}", h.GetPostDetails) // Get Specific Post

	})
}
