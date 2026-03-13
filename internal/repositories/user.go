package repositories

import (
	"fmt"

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
	if err := r.DB.Create(user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func (r *UserRepository) FindById(id string) (*model.User, error) {
	var user model.User

	if err := r.DB.Where("id = ?", id).First(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *UserRepository) Update(id string, data *model.User) (*model.User, error) {
	user, err := r.FindById(id)
	if err != nil {
		return nil, err
	}

	// https://gorm.io/docs/update.html
	if err := r.DB.Model(user).Updates(map[string]any{"name": data.Name, "email": data.Email, "is_active": data.IsActive}).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func (r *UserRepository) Destroy(id string) error {
	user, err := r.FindById(id)
	if err != nil {
		return err
	}

	if err := r.DB.Delete(&user).Error; err != nil {
		return err
	}

	return nil
}

func (r *UserRepository) IsEmailExist(email string) (bool, error) {
	var count int64

	if err := r.DB.Model(&model.User{}).Where("email = ?", email).Count(&count).Error; err != nil {
		return false, err
	}

	return count > 0, nil
}

func (r *UserRepository) IsEmailTaken(id string, email string) (bool, error) {
	var count int64

	fmt.Println(id, email)
	if err := r.DB.Model(&model.User{}).Where("email = ? AND id != ?", email, id).Count(&count).Error; err != nil {
		return false, err
	}

	return count > 0, nil
}
