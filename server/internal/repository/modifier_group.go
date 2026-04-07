package repository

import (
	"context"

	"github.com/BramAristyo/saas-pos-core/server/internal/domain"
	"github.com/BramAristyo/saas-pos-core/server/internal/infrastructure/persistence/database"
	"github.com/BramAristyo/saas-pos-core/server/pkg/filter"
	"github.com/BramAristyo/saas-pos-core/server/pkg/usecase_errors"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ModifierGroupRepository struct {
	DB *gorm.DB
}

func NewModifierGroupRepository(db *gorm.DB) *ModifierGroupRepository {
	return &ModifierGroupRepository{
		DB: db,
	}
}

func (r *ModifierGroupRepository) Paginate(ctx context.Context, req filter.PaginationWithInputFilter) (int64, []domain.ModifierGroup, error) {
	var mg []domain.ModifierGroup
	var totalRows int64

	allowedFields := map[string]string{
		"name":       "name",
		"created_at": "created_at",
	}

	q := database.BuildQuery(r.DB.WithContext(ctx).Model(&domain.ModifierGroup{}), req.DynamicFilter, []string{"name"}, allowedFields)

	if err := q.Count(&totalRows).Error; err != nil {
		return 0, nil, err
	}

	if err := q.Offset(req.Offset()).Limit(req.PaginationInput.PageSize).Find(&mg).Error; err != nil {
		return 0, nil, err
	}

	return totalRows, mg, nil
}

func (r *ModifierGroupRepository) FindById(ctx context.Context, id uuid.UUID) (domain.ModifierGroup, error) {
	var mg domain.ModifierGroup

	err := r.DB.WithContext(ctx).
		Preload("ModifierOptions").
		Preload("ProductModifiers").
		Where("id = ?", id).
		First(&mg).
		Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return domain.ModifierGroup{}, usecase_errors.NotFound
		}
		return domain.ModifierGroup{}, err
	}

	return mg, nil
}

func (r *ModifierGroupRepository) Store(ctx context.Context, mg *domain.ModifierGroup) (domain.ModifierGroup, error) {
	if err := r.DB.WithContext(ctx).Create(mg).Error; err != nil {
		if usecase_errors.IsUniqueViolation(err) {
			return domain.ModifierGroup{}, usecase_errors.DuplicateEntry
		}
		return domain.ModifierGroup{}, err
	}

	return *mg, nil
}

func (r *ModifierGroupRepository) Update(ctx context.Context, id uuid.UUID, mg *domain.ModifierGroup) (domain.ModifierGroup, error) {
	var existing domain.ModifierGroup
	if err := r.DB.WithContext(ctx).Where("id = ?", id).First(&existing).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return domain.ModifierGroup{}, usecase_errors.NotFound
		}
		return domain.ModifierGroup{}, err
	}

	updateData := map[string]any{
		"name":        mg.Name,
		"is_required": mg.IsRequired,
	}

	if err := r.DB.WithContext(ctx).Model(&existing).Updates(updateData).Error; err != nil {
		return domain.ModifierGroup{}, err
	}

	return existing, nil
}

func (r *ModifierGroupRepository) Delete(ctx context.Context, id uuid.UUID) error {
	result := r.DB.WithContext(ctx).Delete(&domain.ModifierGroup{}, "id = ?", id)
	if result.RowsAffected == 0 {
		return usecase_errors.NotFound
	}
	return result.Error
}

func (r *ModifierGroupRepository) Restore(ctx context.Context, id uuid.UUID) error {
	result := r.DB.WithContext(ctx).
		Model(&domain.ModifierGroup{}).
		Unscoped().
		Where("id = ?", id).
		Update("deleted_at", nil)
	if result.RowsAffected == 0 {
		return usecase_errors.NotFound
	}
	return result.Error
}
