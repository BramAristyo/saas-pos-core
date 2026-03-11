package dto

import "github.com/BramAristyo/go-pos-mawish/internal/models"

// https://gin-gonic.com/en/docs/examples/binding-and-validation/

type CreateUserRequest struct {
	Name     string `json:"name" binding:"required, min=2,max=100"`
	Email    string `json:"email" binding:"required,min=6"`
	Password string `json:"password" binding:"required,min=8"`
}

type UpdateUserRequest struct {
	Name     string `json:"name" binding:"omitempty"`
	Email    string `json:"email" binding:"omitempty"`
	IsActive *bool  `json:"is_active" binding:"omitempty"`
}

type UserResponse struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	IsActive  string `json:"is_active"`
	CreatedAt string `json:"created_at"`
}

func ToUserResponse(u models.User) UserResponse {
	return UserResponse{
		ID:        u.ID,
		Name:      u.Name,
		Email:     u.Email,
		CreatedAt: u.CreatedAt.Format("2006-01-02 15:04:05"),
	}
}

func ToCreateUserModel(req CreateUserRequest) models.User {
	return models.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
		IsActive: true,
	}
}

func ToUpdateUserModel(req UpdateUserRequest) models.User {
	return models.User{
		Name:     req.Name,
		Email:    req.Email,
		IsActive: *req.IsActive,
	}
}
