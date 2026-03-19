package repository

import (
	"context"

	"github.com/BramAristyo/go-pos-mawish/internal/domain"
	"github.com/BramAristyo/go-pos-mawish/pkg/filter"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ModifierOptionRepository struct {
	DB *gorm.DB
}

func NewModifierOptionRepository(db *gorm.DB) *ModifierOptionRepository {
	return &ModifierOptionRepository{DB: db}
}

func (r *ModifierOptionRepository) Paginate(ctx context.Context, req filter.PaginationWithInputFilter) (int64, []domain.ModifierOption, error) {
	var mo []domain.ModifierOption
	var totalRows int64

	if err := r.DB.WithContext(ctx).Model(&domain.ModifierOption{}).Count(&totalRows).Error; err != nil {
		return 0, []domain.ModifierOption{}, err
	}

	if err := r.DB.WithContext(ctx).Preload("ModifierGroup").Offset(req.Offset()).Limit(req.PaginationInput.PageSize).Find(&mo).Error; err != nil {
		return 0, []domain.ModifierOption{}, err
	}

	return totalRows, mo, nil
}

func (r *ModifierOptionRepository) FindById(ctx context.Context, id uuid.UUID) (*domain.ModifierOption, error) {
	var mo domain.ModifierOption

	if err := r.DB.WithContext(ctx).Preload("ModifierGroup").First(&mo, "id = ?", id).Error; err != nil {
		return nil, err
	}

	return &mo, nil
}

func (r *ModifierOptionRepository) Store(ctx context.Context, mo *domain.ModifierOption) (*domain.ModifierOption, error) {
	if err := r.DB.WithContext(ctx).Create(mo).Error; err != nil {
		return nil, err
	}

	return mo, nil
}

func (r *ModifierOptionRepository) Update(ctx context.Context, id uuid.UUID, mo *domain.ModifierOption) (*domain.ModifierOption, error) {
	var existing domain.ModifierOption
	if err := r.DB.WithContext(ctx).Where("id = ?", id).First(&existing).Error; err != nil {
		return nil, err
	}

	updateData := map[string]any{
		"modifier_group_id": mo.ModifierGroupID,
		"name":              mo.Name,
		"price_adjustment":  mo.PriceAdjustment,
		"cogs_adjustment":   mo.CogsAdjustment,
		"is_active":         mo.IsActive,
	}

	if err := r.DB.WithContext(ctx).Model(&existing).Updates(updateData).Error; err != nil {
		return nil, err
	}

	return &existing, nil
}

func (r *ModifierOptionRepository) UpdateStatus(ctx context.Context, id uuid.UUID, status bool) (*domain.ModifierOption, error) {
	var mo domain.ModifierOption
	if err := r.DB.WithContext(ctx).Where("id = ?", id).First(&mo).Error; err != nil {
		return nil, err
	}

	if err := r.DB.WithContext(ctx).Model(&mo).Update("is_active", status).Error; err != nil {
		return nil, err
	}

	return &mo, nil
}
