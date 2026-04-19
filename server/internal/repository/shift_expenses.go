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

type ShiftExpensesRepository struct {
	DB *gorm.DB
}

func NewShiftExpensesRepository(db *gorm.DB) *ShiftExpensesRepository {
	return &ShiftExpensesRepository{
		DB: db,
	}
}

func (r *ShiftExpensesRepository) Paginate(ctx context.Context, req filter.PaginationWithInputFilter) (int64, []domain.ShiftExpenses, error) {
	var totalRows int64
	expenses := make([]domain.ShiftExpenses, 0, req.PaginationInput.PageSize)

	allowedFields := map[string]string{
		"shift_id":    "shift_id",
		"type":        "type",
		"amount":      "amount",
		"created_at":  "created_at",
	}

	q := database.BuildQuery(r.DB.WithContext(ctx).Model(&domain.ShiftExpenses{}), req.DynamicFilter, []string{"description"}, allowedFields)

	if err := q.Count(&totalRows).Error; err != nil {
		return 0, []domain.ShiftExpenses{}, err
	}

	if totalRows == 0 {
		return 0, []domain.ShiftExpenses{}, nil
	}

	if err := q.Preload("Shift").Offset(req.Offset()).Limit(req.PaginationInput.PageSize).Find(&expenses).Error; err != nil {
		return 0, []domain.ShiftExpenses{}, err
	}

	return totalRows, expenses, nil
}

func (r *ShiftExpensesRepository) FindById(ctx context.Context, id uuid.UUID) (domain.ShiftExpenses, error) {
	var expense domain.ShiftExpenses

	if err := r.DB.WithContext(ctx).Preload("Shift").Where("id = ?", id).First(&expense).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return domain.ShiftExpenses{}, usecase_errors.NotFound
		}
		return domain.ShiftExpenses{}, err
	}

	return expense, nil
}
