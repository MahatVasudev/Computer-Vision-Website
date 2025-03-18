package store

import (
	"context"
	"net/http"

	"github.com/MahatVasudev/Computer-Vision-Website/server/typestore"
)

type PostStore interface {
	GetPostDetail(ctx context.Context, post_id string) (*typestore.PostFullPicture, error)

	CreatePost(ctx context.Context, p *typestore.PostFullPicture) error
	UploadImages(ctx context.Context, r *http.Request) (string, error)

	CountOfPostOfEachUser(ctx context.Context, user_id string) (int, error)

	Get_All_Posts(ctx context.Context, limit int) (*[]typestore.Post, error)

	Get_All_Posts_From_User(
		ctx context.Context,
		user_name string,
		limit int,
	) (*[]typestore.Post, error)
}
