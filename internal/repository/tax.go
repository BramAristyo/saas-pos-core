package repository

import (
	"context"

	"github.com/BramAristyo/go-pos-mawish/internal/domain"
	"github.com/BramAristyo/go-pos-mawish/pkg/filter"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type TaxRepository struct {
	DB *gorm.DB
}

func NewTaxRepository(db *gorm.DB) *TaxRepository {
	return &TaxRepository{DB: db}
}

func (r *TaxRepository) Paginate(ctx context.Context, req filter.PaginationWithInputFilter) (int64, []domain.Tax, error) {
	t := make([]domain.Tax, 0, req.PaginationInput.PageSize)
	var totalRows int64

	if err := r.DB.WithContext(ctx).Model(&domain.Tax{}).Count(&totalRows).Error; err != nil {
		return 0, []domain.Tax{}, err
	}

	if totalRows == 0 {
		return 0, []domain.Tax{}, nil
	}

	if err := r.DB.WithContext(ctx).Offset(req.Offset()).Limit(req.PaginationInput.PageSize).Find(&t).Error; err != nil {
		return 0, []domain.Tax{}, err
	}

	return totalRows, t, nil
}

func (r *TaxRepository) FindById(ctx context.Context, id uuid.UUID) (domain.Tax, error) {
	var t domain.Tax

	if err := r.DB.WithContext(ctx).First(&t, "id = ?", id).Error; err != nil {
		return domain.Tax{}, err
	}

	return t, nil
}

func (r *TaxRepository) Store(ctx context.Context, t *domain.Tax) (domain.Tax, error) {
	if err := r.DB.WithContext(ctx).Create(t).Error; err != nil {
		return domain.Tax{}, err
	}

	return *t, nil
}

func (r *TaxRepository) Update(ctx context.Context, id uuid.UUID, t *domain.Tax) (domain.Tax, error) {
	var existing domain.Tax
	if err := r.DB.WithContext(ctx).Where("id = ?", id).First(&existing).Error; err != nil {
		return domain.Tax{}, err
	}

	updateData := map[string]any{
		"name":       t.Name,
		"percentage": t.Percentage,
		"is_active":  t.IsActive,
	}

	if err := r.DB.WithContext(ctx).Model(&existing).Updates(updateData).Error; err != nil {
		return domain.Tax{}, err
	}

	return existing, nil
}

func (r *TaxRepository) UpdateStatus(ctx context.Context, id uuid.UUID, status bool) (domain.Tax, error) {
	var t domain.Tax
	if err := r.DB.WithContext(ctx).Where("id = ?", id).First(&t).Error; err != nil {
		return domain.Tax{}, err
	}

	if err := r.DB.WithContext(ctx).Model(&t).Update("is_active", status).Error; err != nil {
		return domain.Tax{}, err
	}

	return t, nil
}

func (r *TaxRepository) DeactiveAll(ctx context.Context) error {
	err := r.DB.Model(domain.Tax{}).Where("is_active = ?", true).Update("is_active", false).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *TaxRepository) FindActiveTaxes(ctx context.Context, req filter.PaginationWithInputFilter) (int64, []domain.Tax, error) {
	t := make([]domain.Tax, 0, req.PaginationInput.PageSize)
	var totalRows int64

	if err := r.DB.WithContext(ctx).Model(&domain.Tax{}).Where("is_active = ?", true).Count(&totalRows).Error; err != nil {
		return 0, []domain.Tax{}, err
	}

	if totalRows == 0 {
		return 0, []domain.Tax{}, nil
	}

	if err := r.DB.WithContext(ctx).Where("is_active = ?", true).Offset(req.Offset()).Limit(req.PaginationInput.PageSize).Find(&t).Error; err != nil {
		return 0, []domain.Tax{}, err
	}

	return totalRows, t, nil
}
