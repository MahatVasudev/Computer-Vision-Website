package user

import (
	"github.com/MahatVasudev/Computer-Vision-Website/server/middleware"
	"github.com/MahatVasudev/Computer-Vision-Website/server/store"
	"github.com/go-chi/chi/v5"
)

type Handler struct {
	store        store.UserStore
	followStore  store.FollowStore
	postStore    store.PostStore
	commentStore store.CommentStore
}

func NewHandler(
	store store.UserStore,
	followStore store.FollowStore,
	postStore store.PostStore,
	commentStore store.CommentStore,
) *Handler {
	return &Handler{
		store:        store,
		followStore:  followStore,
		postStore:    postStore,
		commentStore: commentStore,
	}
}

func (h *Handler) RegisterRoutes(addr string, router *chi.Mux) {
	router.Route(addr, func(router chi.Router) {
		// admin control
		router.With(middleware.AuthMiddleWareUser).Get("/all", h.handleGetAllData)

		// authorization and authentication
		router.Post("/create", h.handleCreateUser)
		router.Post("/login", h.handleLoginUser)
		router.With(middleware.AuthMiddleWareUser).Get("/auth/verify", h.handleVerifyLogin)
		router.Post("/auth/verification/send", h.handleVerifyEmail)
		router.Post("/auth/verification/verify", h.handleVerifyOTP)
		router.Post("/auth/setup", h.handleSetup)
		router.Get("/details/un/{username}/posts", h.handleUsersPosts)
		// single user details operation
		router.With(middleware.AuthMiddleWareUser).
			Get("/details/un/{username}", h.handleUserNameDetails)
		router.Get("/check/un/", h.SeeUsernameExists)
		router.Post("/details/un/{username}", nil)
		router.Patch("/details/un/{username}", nil)
		router.Delete("/details/un/{username}", nil)

	})
}
