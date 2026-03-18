package repository

import (
	"context"

	"github.com/BramAristyo/go-pos-mawish/internal/domain"
	"github.com/BramAristyo/go-pos-mawish/pkg/filter"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CategoryRepository struct {
	DB *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) *CategoryRepository {
	return &CategoryRepository{DB: db}
}

func (r *CategoryRepository) Paginate(ctx context.Context, req filter.PaginationWithInputFilter) (int64, []domain.Category, error) {
	var categories []domain.Category
	var totalRows int64

	if err := r.DB.WithContext(ctx).Model(&domain.Category{}).Count(&totalRows).Error; err != nil {
		return 0, []domain.Category{}, err
	}

	if err := r.DB.WithContext(ctx).Offset(req.Offset()).Limit(req.PaginationInput.PageSize).Find(&categories).Error; err != nil {
		return 0, []domain.Category{}, err
	}

	return totalRows, categories, nil
}

func (r *CategoryRepository) FindById(ctx context.Context, id uuid.UUID) (*domain.Category, error) {
	var category domain.Category

	if err := r.DB.WithContext(ctx).First(&category, "id = ?", id).Error; err != nil {
		return nil, err
	}

	return &category, nil
}

func (r *CategoryRepository) Store(ctx context.Context, category *domain.Category) (*domain.Category, error) {
	if err := r.DB.WithContext(ctx).Create(category).Error; err != nil {
		return nil, err
	}

	return category, nil
}

func (r *CategoryRepository) Update(ctx context.Context, id uuid.UUID, data *domain.Category) (*domain.Category, error) {
	category, err := r.FindById(ctx, id)
	if err != nil {
		return nil, err
	}

	err = r.DB.WithContext(ctx).Model(category).Updates(map[string]any{"name": data.Name, "description": data.Description, "is_active": data.IsActive}).Error
	if err != nil {
		return nil, err
	}

	return category, nil
}
