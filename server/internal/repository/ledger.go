package repository

import (
	"context"

	"github.com/BramAristyo/saas-pos-core/server/internal/domain"
	"github.com/BramAristyo/saas-pos-core/server/pkg/filter"
	"gorm.io/gorm"
)

type LedgerRepository struct {
	DB *gorm.DB
}

func NewLedgerRepository(db *gorm.DB) *LedgerRepository {
	return &LedgerRepository{
		DB: db,
	}
}

func (r *LedgerRepository) Paginate(ctx context.Context, req filter.PaginationWithInputFilter) (int64, []domain.Ledger, error)
func (r *LedgerRepository) TransactionSummary(ctx context.Context, req filter.DynamicFilter)
func (r *LedgerRepository) CashFlowStatement(ctx context.Context, req filter.DynamicFilter)
