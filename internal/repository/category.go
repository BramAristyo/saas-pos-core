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
	var c []domain.Category
	var totalRows int64

	if err := r.DB.WithContext(ctx).Model(&domain.Category{}).Count(&totalRows).Error; err != nil {
		return 0, []domain.Category{}, err
	}

	if err := r.DB.WithContext(ctx).Offset(req.Offset()).Limit(req.PaginationInput.PageSize).Find(&c).Error; err != nil {
		return 0, []domain.Category{}, err
	}

	return totalRows, c, nil
}

func (r *CategoryRepository) FindById(ctx context.Context, id uuid.UUID) (domain.Category, error) {
	var c domain.Category

	if err := r.DB.WithContext(ctx).First(&c, "id = ?", id).Error; err != nil {
		return domain.Category{}, err
	}

	return c, nil
}

func (r *CategoryRepository) Store(ctx context.Context, c *domain.Category) (domain.Category, error) {
	if err := r.DB.WithContext(ctx).Create(c).Error; err != nil {
		return domain.Category{}, err
	}

	return *c, nil
}

func (r *CategoryRepository) Update(ctx context.Context, id uuid.UUID, c *domain.Category) (domain.Category, error) {
	var existing domain.Category
	if err := r.DB.WithContext(ctx).Where("id = ?", id).First(&existing).Error; err != nil {
		return domain.Category{}, err
	}

	updateData := map[string]any{
		"name":        c.Name,
		"description": c.Description,
		"is_active":   c.IsActive,
	}

	if err := r.DB.WithContext(ctx).Model(&existing).Updates(updateData).Error; err != nil {
		return domain.Category{}, err
	}

	return existing, nil
}

func (r *CategoryRepository) UpdateStatus(ctx context.Context, id uuid.UUID, status bool) (domain.Category, error) {
	var c domain.Category
	if err := r.DB.WithContext(ctx).Where("id = ?", id).First(&c).Error; err != nil {
		return domain.Category{}, err
	}

	if err := r.DB.WithContext(ctx).Model(&c).Update("is_active", status).Error; err != nil {
		return domain.Category{}, err
	}

	return c, nil
}
