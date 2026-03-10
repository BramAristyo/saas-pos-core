package dto

// https://gin-gonic.com/en/docs/examples/binding-and-validation/

type CreateUserRequest struct {
	Name     string `json:"name" binding:"required, min=2,max=100"`
	Email    string `json:"email" binding:"required,min=6"`
	Password string `json:"password" binding:"required,min=8"`
}

type UserResponse struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	IsActive  string `json:"is_active"`
	CreatedAt string `json:"created_at"`
}
