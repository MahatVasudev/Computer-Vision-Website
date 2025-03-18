package posts

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/MahatVasudev/Computer-Vision-Website/server/msg"
	"github.com/MahatVasudev/Computer-Vision-Website/server/payloads"
	"github.com/MahatVasudev/Computer-Vision-Website/server/typestore"
	"github.com/MahatVasudev/Computer-Vision-Website/server/writer"
	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

var validate = validator.New()

func (h *Handler) handleGetAllPosts(w http.ResponseWriter, r *http.Request) {

	ctx, cancel := context.WithTimeout(context.TODO(), time.Second*10)

	defer cancel()

	limit := 100

	posts_all, err := h.store.Get_All_Posts(ctx, limit)

	if err != nil {
		if err == msg.ErrorNotFound {
			writer.WriteNotFound(w, err)
			return
		}

		writer.WriteServerError(w, err)
		return
	}

	writer.WriteOk(w, *posts_all)

}

func (h *Handler) handleCreatePosts(w http.ResponseWriter, r *http.Request) {

	ctx, cancel := context.WithTimeout(context.TODO(), time.Second*10)

	defer cancel()

	key, _ := r.Cookie(msg.SSID)

	user_session, err := h.userStore.ReadUserKey(ctx, key.Value)

	if err != nil {
		log.Println(err)
		writer.WriteNotAuthorized(w, msg.ErrorUnAuthorized)
		return
	}

	payload := payloads.PostPayload{
		Title:       r.FormValue("title"),
		Description: r.FormValue("description"),
	}

	if err := validate.Struct(payload); err != nil {
		log.Println("Structure Error:", err)
		writer.WriteServerError(w, msg.ErrorServerSide)
		return
	}

	image_location, err := h.store.UploadImages(ctx, r)

	if err != nil {
		log.Println("Upload Error:", err)
		writer.WriteServerError(w, err)
		return
	}
	post_id := uuid.NewString()
	post := typestore.PostFullPicture{
		PostID:      post_id,
		UserID:      user_session.Id,
		UserName:    user_session.Username,
		Title:       r.FormValue("title"),
		Description: r.FormValue("description"),
		Images:      image_location,
	}

	err = h.store.CreatePost(ctx, &post)

	if err != nil {
		writer.WriteServerError(w, err)
		return
	}

	writer.WriteOk(w, map[string]string{
		"status":  "success",
		"post_id": post_id,
	})
}

func (h *Handler) GetPostDetails(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.TODO(), time.Second*10)

	defer cancel()

	key, err := r.Cookie(msg.SSID)

	user_id := "*"
	if err == nil {
		user_session, err := h.userStore.ReadUserKey(ctx, key.Value)

		if err != nil {
			log.Println(err)
			writer.WriteNotAuthorized(w, msg.ErrorUnAuthorized)
			return
		}

		user_id = user_session.Id
	}

	post_id := chi.URLParam(r, "post_id")

	PostDetails, err := h.store.GetPostDetail(ctx, post_id)

	if err != nil {
		log.Println(err)
		if err == msg.ErrorNotFound {
			writer.WriteNotFound(w, msg.ErrorNotFound)
			return
		}
		writer.WriteServerError(w, msg.ErrorServerSide)
		return
	}

	permission := []string{}

	if user_id == PostDetails.UserID {
		permission = []string{
			"post:edit",
			"post:delete",
			"post:analytics",
		}
	}

	writer.WriteOk(w, map[string]any{
		"post_data":   PostDetails,
		"permissions": permission,
	})
}
