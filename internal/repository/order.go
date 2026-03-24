package repository

import (
	"context"

	"github.com/BramAristyo/go-pos-mawish/internal/domain"
	"github.com/BramAristyo/go-pos-mawish/pkg/filter"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type OrderRepository struct {
	DB *gorm.DB
}

func NewOrderRepository(db *gorm.DB) *OrderRepository {
	return &OrderRepository{DB: db}
}

func (r *OrderRepository) Paginate(ctx context.Context, req filter.PaginationWithInputFilter) (int64, []domain.Order, error) {
	var totalRows int64

	if err := r.DB.WithContext(ctx).Model(&domain.Order{}).Count(&totalRows).Error; err != nil {
		return 0, nil, err
	}

	if totalRows == 0 {
		return 0, nil, nil
	}

	orders := make([]domain.Order, 0, req.PaginationInput.PageSize)
	if err := r.DB.WithContext(ctx).
		Preload("Cashier").
		Offset(req.Offset()).
		Limit(req.PaginationInput.PageSize).
		Find(&orders).
		Error; err != nil {
		return 0, nil, err
	}

	return totalRows, orders, nil
}

func (r *OrderRepository) FindById(ctx context.Context, id uuid.UUID) (domain.Order, error) {
	var existing domain.Order
	if err := r.DB.WithContext(ctx).
		Where("id = ?", id).
		Preload("Shift").
		Preload("Cashier").
		Preload("SalesType").
		Preload("Tax").
		Preload("Discount").
		Preload("VoidedByUser").
		Preload("Items.Product").
		Preload("Items.Bundling").
		Preload("Items.Discount").
		Preload("Payments").
		First(&existing).
		Error; err != nil {
		return domain.Order{}, err
	}

	return existing, nil
}

func (r *OrderRepository) Store(ctx context.Context, order *domain.Order) (domain.Order, error) {
	err := r.DB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(order).Error; err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		return domain.Order{}, err
	}

	return r.FindById(ctx, order.ID)
}

func (r *OrderRepository) Void(ctx context.Context, order *domain.Order) (domain.Order, error) {
	err := r.DB.WithContext(ctx).Model(order).Updates(map[string]any{
		"status":      domain.OrderVoided,
		"void_reason": order.VoidReason,
		"voided_by":   order.VoidedBy,
		"voided_at":   order.VoidedAt,
	}).Error

	if err != nil {
		return domain.Order{}, err
	}

	return r.FindById(ctx, order.ID)
}

func (r *OrderRepository) FindByOrderNumber(ctx context.Context, orderNumber string) (domain.Order, error) {
	var existing domain.Order
	if err := r.DB.WithContext(ctx).
		Where("order_number = ?", orderNumber).
		Preload("Shift").
		Preload("Cashier").
		Preload("SalesType").
		Preload("Tax").
		Preload("Discount").
		Preload("VoidedByUser").
		Preload("Items.Product").
		Preload("Items.Bundling").
		Preload("Items.Discount").
		Preload("Payments").
		First(&existing).
		Error; err != nil {
		return domain.Order{}, err
	}

	return existing, nil
}

func (r *OrderRepository) GetLatestOrder(ctx context.Context) (domain.Order, error) {
	var latest domain.Order
	if err := r.DB.WithContext(ctx).
		Order("created_at desc").
		First(&latest).
		Error; err != nil {
		return domain.Order{}, err
	}

	return latest, nil
}
