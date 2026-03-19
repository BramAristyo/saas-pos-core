package repository

import (
	"context"

	"github.com/BramAristyo/go-pos-mawish/internal/domain"
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

func (r *ProductRepository) Paginate(ctx context.Context, req filter.PaginationWithInputFilter) (int64, []domain.Product, error) {
	var p []domain.Product
	var totalRows int64

	if err := r.DB.WithContext(ctx).Model(&domain.Product{}).Count(&totalRows).Error; err != nil {
		return 0, nil, err
	}

	if err := r.DB.WithContext(ctx).
		Preload("Category").
		Preload("ProductModifiers.ModifierGroup.ModifierOptions").
		Offset(req.Offset()).Limit(req.PaginationInput.PageSize).Find(&p).Error; err != nil {
		return 0, nil, err
	}

	return totalRows, p, nil
}

func (r *ProductRepository) FindById(ctx context.Context, id uuid.UUID) (*domain.Product, error) {
	var p domain.Product

	if err := r.DB.WithContext(ctx).
		Preload("Category").
		Preload("ProductModifiers.ModifierGroup.ModifierOptions").
		Where("id = ?", id).First(&p).Error; err != nil {
		return nil, err
	}

	return &p, nil
}

func (r *ProductRepository) Store(ctx context.Context, p *domain.Product) (*domain.Product, error) {
	if err := r.DB.WithContext(ctx).Create(p).Error; err != nil {
		return nil, err
	}

	return p, nil
}

func (r *ProductRepository) Update(ctx context.Context, id uuid.UUID, p *domain.Product) (*domain.Product, error) {
	err := r.DB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		var existing domain.Product
		if err := tx.Where("id = ?", id).First(&existing).Error; err != nil {
			return err
		}

		updateData := map[string]any{
			"category_id": p.CategoryID,
			"name":        p.Name,
			"description": p.Description,
			"price":       p.Price,
			"cogs":        p.Cogs,
			"image_url":   p.ImageURL,
			"is_active":   p.IsActive,
		}

		if err := tx.Model(&existing).Updates(updateData).Error; err != nil {
			return err
		}

		// Sync ProductModifiers
		var existingIds []uuid.UUID
		tx.Model(&domain.ProductModifier{}).Where("product_id = ?", id).Pluck("modifier_group_id", &existingIds)

		existingMap := make(map[uuid.UUID]bool)
		for _, id := range existingIds {
			existingMap[id] = true
		}

		reqMap := make(map[uuid.UUID]bool)
		for _, pm := range p.ProductModifiers {
			reqMap[pm.ModifierGroupID] = true
		}

		var toDelete []uuid.UUID
		for _, id := range existingIds {
			if !reqMap[id] {
				toDelete = append(toDelete, id)
			}
		}

		var toCreate []domain.ProductModifier
		for _, pm := range p.ProductModifiers {
			if !existingMap[pm.ModifierGroupID] {
				pm.ProductID = id
				toCreate = append(toCreate, pm)
			}
		}

		if len(toDelete) > 0 {
			tx.Where("product_id = ? AND modifier_group_id IN ?", id, toDelete).
				Delete(&domain.ProductModifier{})
		}

		if len(toCreate) > 0 {
			tx.Create(&toCreate)
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return r.FindById(ctx, id)
}

func (r *ProductRepository) UpdateStatus(ctx context.Context, id uuid.UUID, status bool) (*domain.Product, error) {
	var existing domain.Product
	if err := r.DB.WithContext(ctx).Where("id = ?", id).First(&existing).Error; err != nil {
		return nil, err
	}

	if err := r.DB.WithContext(ctx).Model(&existing).Update("is_active", status).Error; err != nil {
		return nil, err
	}

	return &existing, nil
}
