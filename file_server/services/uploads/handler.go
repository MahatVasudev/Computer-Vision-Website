package uploads

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

type Handler struct {
}

func (h *Handler) RegisterRoutes(route string, router *chi.Mux) {
	router.Route(route, func(r chi.Router) {
		r.Post("/profile/{userid}/title", nil)
		r.Post("/profile/{userid}/banner", nil)

		r.Post("/post/{userid}/{postid}/{imageid}/", nil)

		r.Post("/post", nil)

	})
}

func (h *Handler) UploadImage(w http.ResponseWriter, r *http.Request) {

}
