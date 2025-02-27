package typestore

import "time"

type Post struct {
	ID          string    `json:"post_id"    validate:"required"`
	Title       string    `json:"post_title" validate:"required"`
	Description string    `json:"post_desc"  validate:"required"`
	UserID      string    `json:"user_id"    validate:"required"`
	CreatedAt   time.Time `json:"createdAt"  validate:"required"`
	UpdatedAt   time.Time `json:"updatedAt"  validate:"required"`
}

type Image struct {
	ID       string `json:"image_id" validate:"required"`
	Location string `json:"url"      validate:"required"`
	PostID   string `json:"post_id"  validate:"required"`
}

type PostFullPicture struct {
	PostID      string    `json:"post_id"    validate:"required"`
	Title       string    `json:"post_title" validate:"required"`
	Description string    `json:"post_desc"  validate:"required"`
	UserID      string    `json:"user_id"    validate:"required"`
	UserName    string    `json:"username"   validate:"required"`
	Images      []string  `json:"assets"     validate:"required"`
	CreatedAt   time.Time `json:"createdAt"  validate:"required"`
	UpdatedAt   time.Time `json:"updatedAt"  validate:"required"`
}
