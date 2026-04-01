package repository

import (
	"context"

	"github.com/BramAristyo/go-pos-mawish/internal/domain"
	"github.com/BramAristyo/go-pos-mawish/internal/infrastructure/persistence/database"
	"github.com/BramAristyo/go-pos-mawish/pkg/filter"
	"github.com/BramAristyo/go-pos-mawish/pkg/usecase_errors"
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

	allowedFields := map[string]string{
		"name":       "name",
		"created_at": "created_at",
	}

	q := database.BuildQuery(r.DB.WithContext(ctx).Model(&domain.Discount{}), req.DynamicFilter, []string{"name"}, allowedFields)

	if err := q.Count(&totalRows).Error; err != nil {
		return 0, []domain.Discount{}, err
	}

	if totalRows == 0 {
		return 0, []domain.Discount{}, nil
	}

	if err := q.Offset(req.Offset()).Limit(req.PaginationInput.PageSize).Find(&d).Error; err != nil {
		return 0, []domain.Discount{}, err
	}

	return totalRows, d, nil
}

func (r *DiscountRepository) FindById(ctx context.Context, id uuid.UUID) (domain.Discount, error) {
	var d domain.Discount

	if err := r.DB.WithContext(ctx).First(&d, "id = ?", id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return domain.Discount{}, usecase_errors.NotFound
		}
		return domain.Discount{}, err
	}

	return d, nil
}

func (r *DiscountRepository) Store(ctx context.Context, d *domain.Discount) (domain.Discount, error) {
	if err := r.DB.WithContext(ctx).Create(d).Error; err != nil {
		if usecase_errors.IsUniqueViolation(err) {
			return domain.Discount{}, usecase_errors.DuplicateEntry
		}
		return domain.Discount{}, err
	}

	return *d, nil
}

func (r *DiscountRepository) Update(ctx context.Context, id uuid.UUID, d *domain.Discount) (domain.Discount, error) {
	var existing domain.Discount
	if err := r.DB.WithContext(ctx).Where("id = ?", id).First(&existing).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return domain.Discount{}, usecase_errors.NotFound
		}
		return domain.Discount{}, err
	}

	updateData := map[string]any{
		"name":       d.Name,
		"type":       d.Type,
		"value":      d.Value,
		"start_date": d.StartDate,
		"end_date":   d.EndDate,
	}

	if err := r.DB.WithContext(ctx).Model(&existing).Updates(updateData).Error; err != nil {
		return domain.Discount{}, err
	}

	return existing, nil
}

func (r *DiscountRepository) Delete(ctx context.Context, id uuid.UUID) error {
	result := r.DB.WithContext(ctx).Delete(&domain.Discount{}, "id = ?", id)
	if result.RowsAffected == 0 {
		return usecase_errors.NotFound
	}
	return result.Error
}

func (r *DiscountRepository) Restore(ctx context.Context, id uuid.UUID) error {
	result := r.DB.WithContext(ctx).
		Model(&domain.Discount{}).
		Unscoped().
		Where("id = ?", id).
		Update("deleted_at", nil)
	if result.RowsAffected == 0 {
		return usecase_errors.NotFound
	}
	return result.Error
}

func (r *DiscountRepository) Usage(ctx context.Context, req filter.PaginationWithInputFilter) (int64, []domain.DiscountReport, error) {
	var totalRows int64

	allowedFields := map[string]string{
		"created_at": "o.created_at",
	}

	q := database.BuildQuery(r.DB.WithContext(ctx).Table("discounts d"), req.DynamicFilter, nil, allowedFields)

	q = q.
		Select("d.name, COUNT(o.id) as count, COALESCE(SUM(o.subtotal), 0) as gross_discount, COALESCE(SUM(o.discount_amount), 0) as discount").
		Joins("LEFT JOIN orders o ON o.discount_id = d.id AND o.status = ?", domain.OrderCompleted).
		Group("d.name")

	if err := q.Count(&totalRows).Error; err != nil {
		return 0, nil, err
	}

	if totalRows == 0 {
		return 0, nil, nil
	}

	var results []domain.DiscountReport
	if err := q.Offset(req.Offset()).Limit(req.PaginationInput.PageSize).Scan(&results).Error; err != nil {
		return 0, nil, err
	}

	return totalRows, results, nil
}
