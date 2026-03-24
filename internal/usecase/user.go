package usecase

import (
	"context"
	"fmt"

	"github.com/BramAristyo/go-pos-mawish/internal/api/dto"
	"github.com/BramAristyo/go-pos-mawish/internal/domain"
	"github.com/BramAristyo/go-pos-mawish/internal/repository"
	"github.com/BramAristyo/go-pos-mawish/pkg/helper"
	"github.com/BramAristyo/go-pos-mawish/pkg/usecase_errors"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type UserUseCase struct {
	Repo       *repository.UserRepository
	LogUseCase *AuditLogUseCase
}

func NewUserUseCase(r *repository.UserRepository, log *AuditLogUseCase) *UserUseCase {
	return &UserUseCase{
		Repo:       r,
		LogUseCase: log,
	}
}

func (u *UserUseCase) GetAll(ctx context.Context) ([]dto.UserResponse, error) {
	users, err := u.Repo.GetAll()
	if err != nil {
		return nil, err
	}

	responses := make([]dto.UserResponse, 0, len(users))
	for i := range users {
		responses = append(responses, dto.ToUserResponse(&users[i]))
	}

	return responses, nil
}

func (u *UserUseCase) FindById(ctx context.Context, id uuid.UUID) (dto.UserResponse, error) {
	user, err := u.Repo.FindById(id)
	if err != nil {
		return dto.UserResponse{}, err
	}

	return dto.ToUserResponse(&user), nil
}

func (u *UserUseCase) Store(ctx context.Context, req dto.CreateUserRequest) (dto.UserResponse, error) {
	userId, _ := helper.ExtractUserID(ctx)

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

	user := dto.ToCreateUserModel(&req)

	stored, err := u.Repo.Store(&user)
	if err != nil {
		return dto.UserResponse{}, err
	}

	go u.LogUseCase.Log(context.Background(), domain.AuditLog{
		UserID:      userId,
		Action:      domain.ActionCreate,
		Entity:      domain.EntityUser,
		EntityID:    &stored.ID,
		Description: "User created a new user account: " + stored.Email,
	})

	return dto.ToUserResponse(&stored), nil
}

func (u *UserUseCase) Update(ctx context.Context, id uuid.UUID, req dto.UpdateUserRequest) (dto.UserResponse, error) {
	userId, _ := helper.ExtractUserID(ctx)

	user := dto.ToUpdateUserModel(&req)

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

	go u.LogUseCase.Log(context.Background(), domain.AuditLog{
		UserID:      userId,
		Action:      domain.ActionUpdate,
		Entity:      domain.EntityUser,
		EntityID:    &updatedUser.ID,
		Description: "User updated user account: " + updatedUser.Email,
	})

	return dto.ToUserResponse(&updatedUser), nil
}

func (u *UserUseCase) UpdateStatus(ctx context.Context, id uuid.UUID, status bool) (dto.UserResponse, error) {
	userId, _ := helper.ExtractUserID(ctx)

	user, err := u.Repo.UpdateStatus(id, status)
	if err != nil {
		return dto.UserResponse{}, err
	}

	action := domain.ActionActivate
	desc := "User activated user account: " + user.Email
	if !status {
		action = domain.ActionDeactivate
		desc = "User deactivated user account: " + user.Email
	}

	go u.LogUseCase.Log(context.Background(), domain.AuditLog{
		UserID:      userId,
		Action:      action,
		Entity:      domain.EntityUser,
		EntityID:    &user.ID,
		Description: desc,
	})

	return dto.ToUserResponse(&user), nil
}

func (u *UserUseCase) Destroy(ctx context.Context, id uuid.UUID) error {
	userId, _ := helper.ExtractUserID(ctx)

	err := u.Repo.Destroy(id)
	if err != nil {
		return err
	}

	go u.LogUseCase.Log(context.Background(), domain.AuditLog{
		UserID:      userId,
		Action:      domain.ActionDelete,
		Entity:      domain.EntityUser,
		EntityID:    &id,
		Description: fmt.Sprintf("User deleted user account with ID: %s", id.String()),
	})

	return nil
}
