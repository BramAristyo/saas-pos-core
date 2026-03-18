package repository

import (
	"context"

	"github.com/BramAristyo/go-pos-mawish/internal/domain"
	"github.com/BramAristyo/go-pos-mawish/pkg/filter"
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

	if err := r.DB.WithContext(ctx).Model(&domain.ModifierGroup{}).Where("is_active = ?", true).Count(&totalRows).Error; err != nil {
		return 0, nil, err
	}

	if err := r.DB.WithContext(ctx).Where("is_active = ?", true).Offset(req.Offset()).Limit(req.PaginationInput.PageSize).Find(&mg).Error; err != nil {
		return 0, nil, err
	}

	return totalRows, mg, nil
}

func (r *ModifierGroupRepository) FindById(ctx context.Context, id uuid.UUID) (*domain.ModifierGroup, error) {
	var mg domain.ModifierGroup

	err := r.DB.WithContext(ctx).
		Preload("ModifierOptions", "is_active = ?", true).
		Preload("ProductModifiers", "is_active = ?", true).
		Where("id = ?", id).
		First(&mg).
		Error

	if err != nil {
		return nil, err
	}

	return &mg, nil
}

func (r *ModifierGroupRepository) Store(ctx context.Context, mg *domain.ModifierGroup) (*domain.ModifierGroup, error) {
	if err := r.DB.WithContext(ctx).Create(mg).Error; err != nil {
		return nil, err
	}

	return mg, nil
}

func (r *ModifierGroupRepository) Update(ctx context.Context, id uuid.UUID, mg *domain.ModifierGroup) (*domain.ModifierGroup, error) {
	var existing domain.ModifierGroup
	if err := r.DB.WithContext(ctx).Where("id = ?", id).First(&existing).Error; err != nil {
		return nil, err
	}

	updateData := map[string]any{
		"name":        mg.Name,
		"is_required": mg.IsRequired,
		"is_active":   mg.IsActive,
	}

	if err := r.DB.WithContext(ctx).Model(&existing).Updates(updateData).Error; err != nil {
		return nil, err
	}

	return &existing, nil
}

func (r *ModifierGroupRepository) UpdateStatus(ctx context.Context, id uuid.UUID, status bool) (*domain.ModifierGroup, error) {
	var existing domain.ModifierGroup
	if err := r.DB.WithContext(ctx).Where("id = ?", id).First(&existing).Error; err != nil {
		return nil, err
	}

	if err := r.DB.WithContext(ctx).Model(&existing).Update("is_active", status).Error; err != nil {
		return nil, err
	}

	return &existing, nil
}
