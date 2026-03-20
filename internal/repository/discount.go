package repository

import (
	"context"

	"github.com/BramAristyo/go-pos-mawish/internal/domain"
	"github.com/BramAristyo/go-pos-mawish/pkg/filter"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type DiscountRepository struct {
	DB *gorm.DB
}

func NewDiscountRepository(db *gorm.DB) *DiscountRepository {
	return &DiscountRepository{DB: db}
}

func (r *DiscountRepository) Paginate(ctx context.Context, req filter.PaginationWithInputFilter) (int64, []domain.Discount, error) {
	d := make([]domain.Discount, 0, req.PaginationInput.PageSize)
	var totalRows int64

	if err := r.DB.WithContext(ctx).Model(&domain.Discount{}).Count(&totalRows).Error; err != nil {
		return 0, []domain.Discount{}, err
	}

	if totalRows == 0 {
		return 0, []domain.Discount{}, nil
	}

	if err := r.DB.WithContext(ctx).Offset(req.Offset()).Limit(req.PaginationInput.PageSize).Find(&d).Error; err != nil {
		return 0, []domain.Discount{}, err
	}

	return totalRows, d, nil
}

func (r *DiscountRepository) FindById(ctx context.Context, id uuid.UUID) (domain.Discount, error) {
	var d domain.Discount

	if err := r.DB.WithContext(ctx).First(&d, "id = ?", id).Error; err != nil {
		return domain.Discount{}, err
	}

	return d, nil
}

func (r *DiscountRepository) Store(ctx context.Context, d *domain.Discount) (domain.Discount, error) {
	if err := r.DB.WithContext(ctx).Create(d).Error; err != nil {
		return domain.Discount{}, err
	}

	return *d, nil
}

func (r *DiscountRepository) Update(ctx context.Context, id uuid.UUID, d *domain.Discount) (domain.Discount, error) {
	var existing domain.Discount
	if err := r.DB.WithContext(ctx).Where("id = ?", id).First(&existing).Error; err != nil {
		return domain.Discount{}, err
	}

	updateData := map[string]any{
		"name":       d.Name,
		"type":       d.Type,
		"value":      d.Value,
		"start_date": d.StartDate,
		"end_date":   d.EndDate,
		"is_active":  d.IsActive,
	}

	if err := r.DB.WithContext(ctx).Model(&existing).Updates(updateData).Error; err != nil {
		return domain.Discount{}, err
	}

	return existing, nil
}

func (r *DiscountRepository) UpdateStatus(ctx context.Context, id uuid.UUID, status bool) (domain.Discount, error) {
	var d domain.Discount
	if err := r.DB.WithContext(ctx).Where("id = ?", id).First(&d).Error; err != nil {
		return domain.Discount{}, err
	}

	if err := r.DB.WithContext(ctx).Model(&d).Update("is_active", status).Error; err != nil {
		return domain.Discount{}, err
	}

	return d, nil
}
