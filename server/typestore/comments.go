package typestore

type Comment struct {
	Id        string `json:"comment_id" validate:"required"`
	Commented string `json:"commented"  validate:"required"`
	UserId    string `json:"userid"     validate:"required"`
	Username  string `json:"username"   validate:"required"`
}
