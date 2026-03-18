package repository

import (
	"context"

	"github.com/BramAristyo/go-pos-mawish/internal/model"
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

func (r *CategoryRepository) Paginate(ctx context.Context, req filter.PaginationWithInputFilter) (int64, []model.Category, error) {
	var categories []model.Category
	var totalRows int64

	if err := r.DB.WithContext(ctx).Model(&model.Category{}).Count(&totalRows).Error; err != nil {
		return 0, []model.Category{}, err
	}

	if err := r.DB.WithContext(ctx).Offset(req.Offset()).Limit(req.PaginationInput.PageSize).Find(&categories).Error; err != nil {
		return 0, []model.Category{}, err
	}

	return totalRows, categories, nil
}

func (r *CategoryRepository) FindById(ctx context.Context, id uuid.UUID) (*model.Category, error) {
	var category model.Category

	if err := r.DB.WithContext(ctx).First(&category, "id = ?", id).Error; err != nil {
		return nil, err
	}

	return &category, nil
}

func (r *CategoryRepository) Store(ctx context.Context, category *model.Category) (*model.Category, error) {
	if err := r.DB.WithContext(ctx).Create(category).Error; err != nil {
		return nil, err
	}

	return category, nil
}

func (r *CategoryRepository) Update(ctx context.Context, id uuid.UUID, data *model.Category) (*model.Category, error) {
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
