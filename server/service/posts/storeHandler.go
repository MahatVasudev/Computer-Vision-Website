package posts

import (
	"context"

	"github.com/MahatVasudev/Computer-Vision-Website/server/typestore"
)

// GetPostDetail implements store.PostStore.
func (s *Store) GetPostDetail(
	ctx context.Context,
	post_id string,
) (*typestore.PostFullPicture, error) {
	panic("unimplemented")
}

// CreatePost implements store.PostStore.
func (s *Store) CreatePost(ctx context.Context, p *typestore.PostFullPicture) error {
	panic("unimplemented")
}
