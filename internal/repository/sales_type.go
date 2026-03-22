package repository

import (
	"context"

	"github.com/BramAristyo/go-pos-mawish/internal/domain"
	"github.com/BramAristyo/go-pos-mawish/pkg/filter"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type SalesTypeRepository struct {
	DB *gorm.DB
}

func NewSalesTypeRepository(db *gorm.DB) *SalesTypeRepository {
	return &SalesTypeRepository{DB: db}
}

func (r *SalesTypeRepository) Paginate(ctx context.Context, req filter.PaginationWithInputFilter) (int64, []domain.SalesType, error) {
	salesTypes := make([]domain.SalesType, 0, req.PaginationInput.PageSize)
	var totalRows int64

	if err := r.DB.WithContext(ctx).Model(&domain.SalesType{}).Where("is_active = ?", true).Count(&totalRows).Error; err != nil {
		return 0, []domain.SalesType{}, err
	}

	if totalRows == 0 {
		return 0, []domain.SalesType{}, nil
	}

	if err := r.DB.WithContext(ctx).
		Where("is_active = ?", true).
		Offset(req.Offset()).
		Limit(req.PaginationInput.PageSize).
		Order("created_at desc").
		Find(&salesTypes).
		Error; err != nil {
		return 0, []domain.SalesType{}, err
	}

	return totalRows, salesTypes, nil
}

func (r *SalesTypeRepository) FindById(ctx context.Context, id uuid.UUID) (domain.SalesType, error) {
	var existing domain.SalesType
	if err := r.DB.WithContext(ctx).
		Where("id = ? AND is_active = ?", id, true).
		Preload("Charges").
		First(&existing).
		Error; err != nil {
		return domain.SalesType{}, err
	}

	return existing, nil
}

func (r *SalesTypeRepository) Store(ctx context.Context, s *domain.SalesType) (domain.SalesType, error) {
	if err := r.DB.WithContext(ctx).Create(s).Error; err != nil {
		return domain.SalesType{}, err
	}

	return *s, nil
}

func (r *SalesTypeRepository) SmartUpdate(ctx context.Context, id uuid.UUID, s *domain.SalesType) (domain.SalesType, error) {
	err := r.DB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		var existing domain.SalesType
		if err := tx.Where("id = ?", id).Preload("Charges").First(&existing).Error; err != nil {
			return err
		}

		updateData := map[string]any{
			"name": s.Name,
		}

		if err := tx.Model(&existing).Updates(updateData).Error; err != nil {
			return err
		}

		// Sync AdditionalCharge
		existingMap := make(map[uuid.UUID]bool)
		for _, c := range existing.Charges {
			existingMap[c.ID] = true
		}

		reqMap := make(map[uuid.UUID]bool)
		for _, c := range s.Charges {
			if c.ID != uuid.Nil {
				reqMap[c.ID] = true
			}
		}

		var toDelete []uuid.UUID
		for _, c := range existing.Charges {
			if !reqMap[c.ID] {
				toDelete = append(toDelete, c.ID)
			}
		}

		var toCreate []domain.AdditionalCharge
		var toUpdate []domain.AdditionalCharge
		for _, c := range s.Charges {
			if c.ID == uuid.Nil {
				c.SalesTypeID = id
				toCreate = append(toCreate, c)
			} else if existingMap[c.ID] {
				toUpdate = append(toUpdate, c)
			}
		}

		if len(toDelete) > 0 {
			if err := tx.Where("id IN ?", toDelete).Unscoped().Delete(&domain.AdditionalCharge{}).Error; err != nil {
				return err
			}
		}
		if len(toCreate) > 0 {
			if err := tx.Create(&toCreate).Error; err != nil {
				return err
			}
		}
		if len(toUpdate) > 0 {
			for _, c := range toUpdate {
				if err := tx.Model(&domain.AdditionalCharge{}).Where("id = ?", c.ID).Updates(map[string]any{
					"name":   c.Name,
					"amount": c.Amount,
					"type":   c.Type,
				}).Error; err != nil {
					return err
				}
			}
		}
		return nil
	})

	if err != nil {
		return domain.SalesType{}, err
	}

	return r.FindById(ctx, id)
}

func (r *SalesTypeRepository) UpdateStatus(ctx context.Context, id uuid.UUID, status bool) (domain.SalesType, error) {
	var existing domain.SalesType
	if err := r.DB.WithContext(ctx).Where("id = ?", id).First(&existing).Error; err != nil {
		return domain.SalesType{}, err
	}

	if err := r.DB.WithContext(ctx).Model(&existing).Update("is_active", status).Error; err != nil {
		return domain.SalesType{}, err
	}

	return existing, nil
}
