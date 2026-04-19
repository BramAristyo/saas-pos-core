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

type ExpenseRepository struct {
	DB *gorm.DB
}

func NewExpenseRepository(db *gorm.DB) *ExpenseRepository {
	return &ExpenseRepository{
		DB: db,
	}
}

func (r *ExpenseRepository) Paginate(ctx context.Context, req filter.PaginationWithInputFilter) (int64, []domain.Expense, error) {
	var totalRows int64
	expenses := make([]domain.Expense, 0, req.PaginationInput.PageSize)

	allowedFields := map[string]string{
		"date":        "date",
		"description": "description",
		"coa_id":      "coa_id",
		"amount":       "amount",
		"created_at":  "created_at",
	}

	q := database.BuildQuery(r.DB.WithContext(ctx).Model(&domain.Expense{}), req.DynamicFilter, []string{"description"}, allowedFields)

	if err := q.Count(&totalRows).Error; err != nil {
		return 0, []domain.Expense{}, err
	}

	if totalRows == 0 {
		return 0, []domain.Expense{}, nil
	}

	if err := q.Preload("COA").Offset(req.Offset()).Limit(req.PaginationInput.PageSize).Find(&expenses).Error; err != nil {
		return 0, []domain.Expense{}, err
	}

	return totalRows, expenses, nil
}

func (r *ExpenseRepository) FindById(ctx context.Context, id uuid.UUID) (domain.Expense, error) {
	var expense domain.Expense

	if err := r.DB.WithContext(ctx).Preload("COA").Where("id = ?", id).First(&expense).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return domain.Expense{}, usecase_errors.NotFound
		}
		return domain.Expense{}, err
	}

	return expense, nil
}

func (r *ExpenseRepository) Store(ctx context.Context, expense *domain.Expense) (domain.Expense, error) {
	if err := r.DB.WithContext(ctx).Create(expense).Error; err != nil {
		return domain.Expense{}, err
	}

	return r.FindById(ctx, expense.ID)
}

func (r *ExpenseRepository) Update(ctx context.Context, id uuid.UUID, expense *domain.Expense) (domain.Expense, error) {
	var existing domain.Expense
	if err := r.DB.WithContext(ctx).Where("id = ?", id).First(&existing).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return domain.Expense{}, usecase_errors.NotFound
		}
		return domain.Expense{}, err
	}

	updateData := map[string]any{
		"coa_id":      expense.COAID,
		"amount":      expense.Amount,
		"description": expense.Description,
		"date":        expense.Date,
	}

	if err := r.DB.WithContext(ctx).Model(&existing).Updates(updateData).Error; err != nil {
		return domain.Expense{}, err
	}

	return r.FindById(ctx, id)
}

func (r *ExpenseRepository) Delete(ctx context.Context, id uuid.UUID) error {
	result := r.DB.WithContext(ctx).Delete(&domain.Expense{}, "id = ?", id)
	if result.RowsAffected == 0 {
		return usecase_errors.NotFound
	}
	return result.Error
}

func (r *ExpenseRepository) Restore(ctx context.Context, id uuid.UUID) error {
	result := r.DB.WithContext(ctx).
		Model(&domain.Expense{}).
		Unscoped().
		Where("id = ?", id).
		Update("deleted_at", nil)
	if result.RowsAffected == 0 {
		return usecase_errors.NotFound
	}
	return result.Error
}
