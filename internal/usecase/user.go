package usecase

import (
	"fmt"

	"github.com/BramAristyo/go-pos-mawish/internal/api/dto"
	"github.com/BramAristyo/go-pos-mawish/internal/repository"
	"github.com/BramAristyo/go-pos-mawish/pkg/usecase_errors"
	"golang.org/x/crypto/bcrypt"
)

type UserUseCase struct {
	Repo *repository.UserRepository
}

func NewUserUseCase(r *repository.UserRepository) *UserUseCase {
	return &UserUseCase{
		Repo: r,
	}
}

func (s *UserUseCase) GetAll() ([]dto.UserResponse, error) {
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

func (s *UserUseCase) FindById(id string) (dto.UserResponse, error) {
	user, err := s.Repo.FindById(id)
	if err != nil {
		return dto.UserResponse{}, err
	}

	res := dto.ToUserResponse(*user)
	return res, nil
}

func (s *UserUseCase) Store(req dto.CreateUserRequest) (dto.UserResponse, error) {
	exist, err := s.Repo.IsEmailExist(req.Email)
	if err != nil {
		return dto.UserResponse{}, err
	}

	if exist {
		return dto.UserResponse{}, usecase_errors.EmailExist
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

func (s *UserUseCase) Update(id string, req dto.UpdateUserRequest) (dto.UserResponse, error) {
	user := dto.ToUpdateUserModel(req)

	fmt.Println(id, req.Email)
	exist, err := s.Repo.IsEmailTaken(id, req.Email)

	if err != nil {
		return dto.UserResponse{}, err
	}

	if exist {
		return dto.UserResponse{}, usecase_errors.EmailExist
	}

	updatedUser, err := s.Repo.Update(id, &user)
	if err != nil {
		return dto.UserResponse{}, err
	}

	return dto.ToUserResponse(*updatedUser), nil
}

func (s *UserUseCase) UpdateStatus(id string, status bool) (dto.UserResponse, error) {
	user, err := s.Repo.UpdateStatus(id, status)
	if err != nil {
		return dto.UserResponse{}, err
	}

	return dto.ToUserResponse(*user), nil
}

func (s *UserUseCase) Destroy(id string) error {
	return s.Repo.Destroy(id)
}
