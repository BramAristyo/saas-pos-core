package repository

import (
	"github.com/BramAristyo/go-pos-mawish/internal/domain"
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

func (r *UserRepository) GetAll() ([]domain.User, error) {
	var u []domain.User

	if err := r.DB.Find(&u).Error; err != nil {
		return nil, err
	}

	return u, nil
}

func (r *UserRepository) Store(u *domain.User) (*domain.User, error) {
	if err := r.DB.Create(u).Error; err != nil {
		return nil, err
	}

	return u, nil
}

func (r *UserRepository) FindById(id string) (*domain.User, error) {
	var u domain.User

	if err := r.DB.Where("id = ?", id).First(&u).Error; err != nil {
		return nil, err
	}

	return &u, nil
}

func (r *UserRepository) FindByEmail(email string) (*domain.User, error) {
	var u domain.User

	if err := r.DB.Where("email = ?", email).First(&u).Error; err != nil {
		return nil, err
	}
	return &u, nil
}

func (r *UserRepository) Update(id string, u *domain.User) (*domain.User, error) {
	var existing domain.User
	if err := r.DB.Where("id = ?", id).First(&existing).Error; err != nil {
		return nil, err
	}

	updateData := map[string]any{
		"name":      u.Name,
		"email":     u.Email,
		"role":      u.Role,
		"is_active": u.IsActive,
	}

	if err := r.DB.Model(&existing).Updates(updateData).Error; err != nil {
		return nil, err
	}

	return &existing, nil
}

func (r *UserRepository) UpdateStatus(id string, status bool) (*domain.User, error) {
	var u domain.User
	if err := r.DB.Where("id = ?", id).First(&u).Error; err != nil {
		return nil, err
	}

	if err := r.DB.Model(&u).Update("is_active", status).Error; err != nil {
		return nil, err
	}

	return &u, nil
}

func (r *UserRepository) Destroy(id string) error {
	var u domain.User
	if err := r.DB.Where("id = ?", id).First(&u).Error; err != nil {
		return err
	}

	if err := r.DB.Delete(&u).Error; err != nil {
		return err
	}

	return nil
}

func (r *UserRepository) IsEmailExist(email string) (bool, error) {
	var count int64

	if err := r.DB.Model(&domain.User{}).Where("email = ?", email).Count(&count).Error; err != nil {
		return false, err
	}

	return count > 0, nil
}

func (r *UserRepository) IsEmailTaken(id string, email string) (bool, error) {
	var count int64

	if err := r.DB.Model(&domain.User{}).Where("email = ? AND id != ?", email, id).Count(&count).Error; err != nil {
		return false, err
	}

	return count > 0, nil
}
