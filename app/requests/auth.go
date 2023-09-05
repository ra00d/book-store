package requests

type LoginRequestBody struct {
	Email    string `json:"email" validate:"email"`
	Password string `json:"password" validate:"required,min=6"`
}
