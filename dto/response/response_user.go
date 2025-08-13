package response

type UserResponse struct {
	Name   string `json:"name"`
	Email  string `json:"email"`
	UserID uint   `json:"user_id"`
}
