package usecase

import (
	"fmt"
	"time"

	"github.com/BramAristyo/go-pos-mawish/internal/infrastructure/config"
	"github.com/BramAristyo/go-pos-mawish/internal/api/dto"
	"github.com/BramAristyo/go-pos-mawish/internal/repository"
	"github.com/BramAristyo/go-pos-mawish/pkg/usecase_errors"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type AuthUseCase struct {
	Repo *repository.UserRepository
	Cfg  *config.Config
}

func NewAuthUseCase(repo *repository.UserRepository, cfg *config.Config) *AuthUseCase {
	return &AuthUseCase{
		Repo: repo,
		Cfg:  cfg,
	}
}

func (s *AuthUseCase) Login(req dto.LoginRequest) (dto.LoginResponse, error) {
	user, err := s.Repo.FindByEmail(req.Email)
	if err != nil {
		return dto.LoginResponse{}, err
	}

	if !user.IsActive {
		return dto.LoginResponse{}, usecase_errors.UserNotActive
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		fmt.Println(user.Password, req.Password)
		return dto.LoginResponse{}, usecase_errors.InvalidPassword
	}

	atc := jwt.MapClaims{}

	atc["userID"] = user.ID
	atc["email"] = user.Email
	atc["role"] = user.Role
	atc["exp"] = time.Now().Add(s.Cfg.JWT.AccessTokenExpireDuration).Unix()

	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atc)
	token, err := at.SignedString([]byte(s.Cfg.JWT.Secret))
	if err != nil {
		return dto.LoginResponse{}, err
	}

	return dto.LoginResponse{
		Token: token,
		User:  dto.ToUserResponse(*user),
	}, nil
}
