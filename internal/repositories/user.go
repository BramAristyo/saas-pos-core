package repositories

import (
	model "github.com/BramAristyo/go-pos-mawish/internal/models"
	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		DB: db,
	}
}

func (r *UserRepository) GetAll() ([]model.User, error) {
	var users []model.User

	// https://pkg.go.dev/gorm.io/gorm@v1.31.1#DB.Find
	if err := r.DB.Find(&users).Error; err != nil {
		return nil, err
	}

	return users, nil
}

func (r *UserRepository) Store(user *model.User) (*model.User, error) {
	if err := r.DB.Create(&user).Error; err != nil {
		return &model.User{}, nil
	}

	return user, nil
}
