package services

import (
	"github.com/BramAristyo/go-pos-mawish/internal/dto"
	"github.com/BramAristyo/go-pos-mawish/internal/repositories"
)

type UserService struct {
	Repo *repositories.UserRepository
}

func NewUserService(r *repositories.UserRepository) *UserService {
	return &UserService{
		Repo: r,
	}
}

func (s *UserService) GetAll() ([]dto.UserResponse, error) {
	users, err := s.Repo.GetAll()
	if err != nil {
		return nil, err
	}

	// https://pkg.go.dev/builtin#make
	responses := make([]dto.UserResponse, 0, len(users))
	for _, u := range users {
		responses = append(responses, dto.UserResponse{
			ID:    u.ID,
			Name:  u.Name,
			Email: u.Email,
		})
	}

	return responses, nil
}
