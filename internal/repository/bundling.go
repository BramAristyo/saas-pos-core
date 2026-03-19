package repository

import (
	"context"

	"github.com/BramAristyo/go-pos-mawish/internal/domain"
	"github.com/BramAristyo/go-pos-mawish/pkg/filter"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type BundlingRepository struct {
	DB *gorm.DB
}

func NewBundlingRepository(db *gorm.DB) *BundlingRepository {
	return &BundlingRepository{DB: db}
}

func (r *BundlingRepository) Paginate(ctx context.Context, req filter.PaginationWithInputFilter) (int64, []domain.BundlingPackage, error) {
	var bps []domain.BundlingPackage
	var totalRows int64

	if err := r.DB.WithContext(ctx).
		Model(&domain.BundlingPackage{}).
		Where("is_active = ?", true).
		Count(&totalRows).Error; err != nil {
		return 0, nil, err
	}

	if err := r.DB.WithContext(ctx).
		Offset(req.Offset()).
		Limit(req.PaginationInput.PageSize).
		Find(&bps).Error; err != nil {
		return 0, nil, err
	}

	return totalRows, bps, nil
}

func (r *BundlingRepository) FindById(ctx context.Context, id uuid.UUID) (*domain.BundlingPackage, error) {
	var bp domain.BundlingPackage
	if err := r.DB.WithContext(ctx).
		Preload("BundlingItems.Product.Category").
		Where("id = ?", id).
		First(&bp).Error; err != nil {
		return nil, err
	}

	return &bp, nil
}

func (r *BundlingRepository) Store(ctx context.Context, bp *domain.BundlingPackage) (*domain.BundlingPackage, error) {
	if err := r.DB.WithContext(ctx).Create(bp).Error; err != nil {
		return nil, err
	}

	return bp, nil
}

func (r *BundlingRepository) Update(ctx context.Context, id uuid.UUID, bp *domain.BundlingPackage) (*domain.BundlingPackage, error) {
	err := r.DB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		var existing domain.BundlingPackage
		if err := tx.Where("id = ?", id).Preload("BundlingItems").First(&existing).Error; err != nil {
			return err
		}

		updateData := map[string]any{
			"name":        bp.Name,
			"description": bp.Description,
			"price":       bp.Price,
			"cogs":        bp.Cogs,
			"image_url":   bp.ImageURL,
			"is_active":   bp.IsActive,
		}

		if err := tx.Model(&existing).Updates(updateData).Error; err != nil {
			return err
		}

		// TODO: Need Smart Sync SOON
		tx.Where("bundling_package_id = ?", id).Delete(&domain.BundlingItem{})
		for i := range bp.BundlingItems {
			bp.BundlingItems[i].BundlingPackageID = id
		}
		tx.Create(&bp.BundlingItems)

		return nil
	})

	if err != nil {
		return nil, err
	}

	return r.FindById(ctx, id)
}

func (r *BundlingRepository) UpdateStatus(ctx context.Context, id uuid.UUID, status bool) (*domain.BundlingPackage, error) {
	var existing domain.BundlingPackage
	if err := r.DB.WithContext(ctx).Where("id = ?", id).First(&existing).Error; err != nil {
		return nil, err
	}

	if err := r.DB.WithContext(ctx).Model(&existing).Update("is_active", status).Error; err != nil {
		return nil, err
	}

	return &existing, nil
}
