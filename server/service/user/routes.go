package user

import (
	"github.com/MahatVasudev/Computer-Vision-Website/server/types"
	"github.com/go-chi/chi/v5"
)

type Handler struct {
	store types.UserStore
}

func NewHandler(store types.UserStore) *Handler {
	return &Handler{
		store: store,
	}
}

func (h *Handler) RegisterRoutes(addr string, router *chi.Mux) {
	router.Route(addr, func(router chi.Router) {
		router.Get("/all", h.handleGetAllData)
	})
}
