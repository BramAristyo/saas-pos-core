package repositories

import (
	"context"

	"github.com/BramAristyo/go-pos-mawish/internal/models"
	"github.com/BramAristyo/go-pos-mawish/pkg/filter"
	"gorm.io/gorm"
)

type CategoryRepository struct {
	DB *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) *CategoryRepository {
	return &CategoryRepository{DB: db}
}

func (r *CategoryRepository) Paginate(ctx context.Context, req filter.PaginationWithInputFilter) (int64, []models.Category, error) {
	var categories []models.Category
	var totalRows int64

	if err := r.DB.WithContext(ctx).Model(&models.Category{}).Count(&totalRows).Error; err != nil {
		return 0, []models.Category{}, err
	}

	if err := r.DB.WithContext(ctx).Offset(req.Offset()).Limit(req.PaginationInput.PageSize).Find(&categories).Error; err != nil {
		return 0, []models.Category{}, err
	}

	return totalRows, categories, nil
}

func (r *CategoryRepository) FindById(ctx context.Context, id int) (*models.Category, error) {
	var category models.Category

	if err := r.DB.WithContext(ctx).First(&category, "id = ?", id).Error; err != nil {
		return nil, err
	}

	return &category, nil
}

func (r *CategoryRepository) Store(ctx context.Context, category *models.Category) (*models.Category, error) {
	if err := r.DB.WithContext(ctx).Create(category).Error; err != nil {
		return nil, err
	}

	return category, nil
}

func (r *CategoryRepository) Update(ctx context.Context, id int, data *models.Category) (*models.Category, error) {
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
