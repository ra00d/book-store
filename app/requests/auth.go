package requests

type LoginRequestBody struct {
	Email    string `json:"email" validate:"email"`
	Password string `json:"password" validate:"required,min=6"`
}
type SignUpRequest struct {
	Email                string `json:"email" form:"email" validate:"required,email"`
	Name                 string `json:"name" form:"name" validate:"required"`
	Password             string `json:"password" form:"password" validate:"min=6"`
	PasswordConfirmation string `json:"passwordConfirmation" form:"passwordConfirmation" validate:"required_with=Password,eqcsfield=Password"`
	Phone                string `json:"phone,omitempty" form:"phone,omitempty" `
}
