package payloads

type FollowOrUnFollow struct {
	Following_id string `json:"following_id" validate:"required"`
	Type_follow  bool   `json:"type"         validate:"required"`
}
