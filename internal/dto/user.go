package dto

import (
	"github.com/BramAristyo/go-pos-mawish/internal/model"
	"github.com/google/uuid"
)

type CreateUserRequest struct {
	Name     string `json:"name" binding:"required,min=2,max=100"`
	Email    string `json:"email" binding:"required,min=6"`
	Password string `json:"password" binding:"required,min=8"`
}

type UpdateUserRequest struct {
	Name     string `json:"name" binding:"required,min=2,max=100"`
	Email    string `json:"email" binding:"required,min=6"`
	IsActive bool   `json:"isActive"`
}

type UserResponse struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	IsActive  bool      `json:"isActive"`
	CreatedAt string    `json:"createdAt"`
}

func ToUserResponse(u model.User) UserResponse {
	return UserResponse{
		ID:        u.ID,
		Name:      u.Name,
		Email:     u.Email,
		IsActive:  u.IsActive,
		CreatedAt: u.CreatedAt.Format("2006-01-02 15:04:05"),
	}
}

func ToCreateUserModel(req CreateUserRequest) model.User {
	return model.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
		IsActive: true,
	}
}

func ToUpdateUserModel(req UpdateUserRequest) model.User {
	return model.User{
		Name:     req.Name,
		Email:    req.Email,
		IsActive: req.IsActive,
	}
}
