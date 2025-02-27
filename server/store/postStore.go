package store

import (
	"context"

	"github.com/MahatVasudev/Computer-Vision-Website/server/typestore"
)

type PostStore interface {
	GetPostDetail(ctx context.Context, post_id string) (*typestore.PostFullPicture, error)

	CreatePost(ctx context.Context, p *typestore.PostFullPicture) error
}
