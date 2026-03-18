package repository

import (
	"context"

	"github.com/BramAristyo/go-pos-mawish/internal/model"
	"github.com/BramAristyo/go-pos-mawish/pkg/filter"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ProductRepository struct {
	DB *gorm.DB
}

func NewProductRepository(db *gorm.DB) *ProductRepository {
	return &ProductRepository{
		DB: db,
	}
}

func (r *ProductRepository) Paginate(ctx context.Context, req filter.PaginationWithInputFilter) (int64, []model.Product, error) {
	var products []model.Product
	var totalRows int64

	if err := r.DB.WithContext(ctx).Model(&model.Product{}).Where("is_active = ?", true).Count(&totalRows).Error; err != nil {
		return 0, nil, err
	}

	if err := r.DB.WithContext(ctx).Preload("Category").Where("is_active = ?", true).Offset(req.Offset()).Limit(req.PaginationInput.PageSize).Find(&products).Error; err != nil {
		return 0, nil, err
	}

	return totalRows, products, nil
}

func (r *ProductRepository) FindById(ctx context.Context, id uuid.UUID) (*model.Product, error) {
	var product model.Product

	if err := r.DB.WithContext(ctx).Preload("Category").Where("id = ? AND is_active = ?", id, true).First(&product).Error; err != nil {
		return nil, err
	}

	return &product, nil
}

func (r *ProductRepository) Store(ctx context.Context, product *model.Product) (*model.Product, error) {
	if err := r.DB.WithContext(ctx).Create(product).Error; err != nil {
		return nil, err
	}

	return product, nil
}

func (r *ProductRepository) Update(ctx context.Context, id uuid.UUID, product *model.Product) (*model.Product, error) {
	var existing model.Product
	if err := r.DB.WithContext(ctx).Where("id = ? AND is_active = ?", id, true).First(&existing).Error; err != nil {
		return nil, err
	}

	if err := r.DB.WithContext(ctx).Model(&existing).Updates(product).Error; err != nil {
		return nil, err
	}

	return &existing, nil
}

func (r *ProductRepository) ChangeStatus(ctx context.Context, id uuid.UUID, status bool) (*model.Product, error) {
	var product model.Product
	if err := r.DB.WithContext(ctx).Where("id = ?", id).First(&product).Error; err != nil {
		return nil, err
	}

	if err := r.DB.WithContext(ctx).Model(&product).Update("is_active", status).Error; err != nil {
		return nil, err
	}

	return &product, nil
}
