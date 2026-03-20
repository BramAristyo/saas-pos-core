package usecase

import (
	"fmt"

	"github.com/BramAristyo/go-pos-mawish/internal/api/dto"
	"github.com/BramAristyo/go-pos-mawish/internal/repository"
	"github.com/BramAristyo/go-pos-mawish/pkg/usecase_errors"
	"github.com/google/uuid"
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

func (u *UserUseCase) GetAll() ([]dto.UserResponse, error) {
	users, err := u.Repo.GetAll()
	if err != nil {
		return nil, err
	}

	// https://pkg.go.dev/builtin#make
	responses := make([]dto.UserResponse, 0, len(users))
	for _, usr := range users {
		responses = append(responses, dto.ToUserResponse(usr))
	}

	return responses, nil
}

func (u *UserUseCase) FindById(id uuid.UUID) (dto.UserResponse, error) {
	user, err := u.Repo.FindById(id)
	if err != nil {
		return dto.UserResponse{}, err
	}

	res := dto.ToUserResponse(user)
	return res, nil
}

func (u *UserUseCase) Store(req dto.CreateUserRequest) (dto.UserResponse, error) {
	exist, err := u.Repo.IsEmailExist(req.Email)
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

	_, err = u.Repo.Store(&user)
	if err != nil {
		return dto.UserResponse{}, err
	}

	return dto.ToUserResponse(user), nil
}

func (u *UserUseCase) Update(id uuid.UUID, req dto.UpdateUserRequest) (dto.UserResponse, error) {
	user := dto.ToUpdateUserModel(req)

	fmt.Println(id, req.Email)
	exist, err := u.Repo.IsEmailTaken(id, req.Email)

	if err != nil {
		return dto.UserResponse{}, err
	}

	if exist {
		return dto.UserResponse{}, usecase_errors.EmailExist
	}

	updatedUser, err := u.Repo.Update(id, &user)
	if err != nil {
		return dto.UserResponse{}, err
	}

	return dto.ToUserResponse(updatedUser), nil
}

func (u *UserUseCase) UpdateStatus(id uuid.UUID, status bool) (dto.UserResponse, error) {
	user, err := u.Repo.UpdateStatus(id, status)
	if err != nil {
		return dto.UserResponse{}, err
	}

	return dto.ToUserResponse(user), nil
}

func (u *UserUseCase) Destroy(id uuid.UUID) error {
	return u.Repo.Destroy(id)
}
