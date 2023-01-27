package usersdto

type UserResponse struct {
	ID       int    `json:"id"`
	Image    string `json:"image" `
	Username  string `json:"username" form:"username" validate:"required"`
	Password string `json:"password" form:"password" validate:"required"`
}
