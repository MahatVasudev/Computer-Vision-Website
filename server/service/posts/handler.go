package posts

import (
	"context"
	"net/http"
	"time"
)

func (h *Handler) handleCreatePosts(w http.ResponseWriter, r *http.Request) {

	_, cancel := context.WithTimeout(context.TODO(), time.Second*10)

	defer cancel()

}
