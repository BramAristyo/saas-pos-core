package service

import (
	"fmt"

	"github.com/BramAristyo/go-pos-mawish/internal/dto"
	"github.com/BramAristyo/go-pos-mawish/internal/repository"
	"github.com/BramAristyo/go-pos-mawish/pkg/service_errors"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	Repo *repository.UserRepository
}

func NewUserService(r *repository.UserRepository) *UserService {
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
		responses = append(responses, dto.ToUserResponse(u))
	}

	return responses, nil
}

func (s *UserService) FindById(id string) (dto.UserResponse, error) {
	user, err := s.Repo.FindById(id)
	if err != nil {
		return dto.UserResponse{}, err
	}

	res := dto.ToUserResponse(*user)
	return res, nil
}

func (s *UserService) Store(req dto.CreateUserRequest) (dto.UserResponse, error) {
	exist, err := s.Repo.IsEmailExist(req.Email)
	if err != nil {
		return dto.UserResponse{}, err
	}

	if exist {
		return dto.UserResponse{}, service_errors.EmailExist
	}

	hashed, err := bcrypt.GenerateFromPassword([]byte(req.Password), 10)
	if err != nil {
		return dto.UserResponse{}, err
	}

	req.Password = string(hashed)

	user := dto.ToCreateUserModel(req)

	_, err = s.Repo.Store(&user)
	if err != nil {
		return dto.UserResponse{}, err
	}

	return dto.ToUserResponse(user), nil
}

func (s *UserService) Update(id string, req dto.UpdateUserRequest) (dto.UserResponse, error) {
	user := dto.ToUpdateUserModel(req)

	fmt.Println(id, req.Email)
	exist, err := s.Repo.IsEmailTaken(id, req.Email)

	if err != nil {
		return dto.UserResponse{}, err
	}

	if exist {
		return dto.UserResponse{}, service_errors.EmailExist
	}

	updatedUser, err := s.Repo.Update(id, &user)
	if err != nil {
		return dto.UserResponse{}, err
	}

	return dto.ToUserResponse(*updatedUser), nil
}

func (s *UserService) Destroy(id string) error {
	return s.Repo.Destroy(id)
}
