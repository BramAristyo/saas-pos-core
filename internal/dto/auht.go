package dto

type LoginRequest struct {
	Email    string `json:"email" binding:"email,required"`
	Password string `json:"password" binding:"required,min=8"`
}

type RegisterRequest struct {
	Name                 string `json:"name" binding:"required,min=2,max=100"`
	Email                string `json:"email" binding:"required,min=6"`
	Password             string `json:"password" binding:"required,min=8"`
	PasswordConfirmation string `json:"passwordConfirmation" binding:"required,min=8"`
}
