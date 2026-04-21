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

func (r *LedgerRepository) Paginate(ctx context.Context, startDate string, endDate string, limit int, offset int) (int64, []domain.LedgerWithBalance, error) {
	var totalRows int64
	var results []domain.LedgerWithBalance

	if err := r.DB.WithContext(ctx).Model(&domain.Ledger{}).
		Where("transaction_date BETWEEN ? AND ?", startDate, endDate).
		Count(&totalRows).Error; err != nil {
		return 0, nil, err
	}

	q := `
		WITH GlobalOpening AS (
			SELECT COALESCE(SUM(
				CASE
					WHEN ca.type = 'in' THEN l.amount
					WHEN ca.type = 'out' THEN -l.amount
					ELSE 0
				END
			),0) AS balance
			FROM ledgers l
			JOIN chart_of_accounts ca ON ca.id = l.coa_id
			WHERE l.transaction_date < ?
		),
		RunningData AS (
		SELECT
				l.*,
				(SELECT balance FROM GlobalOpening) +
				SUM(
					CASE
						WHEN ca.type = 'in' THEN l.amount
						WHEN ca.type = 'out' THEN -l.amount
						ELSE 0
					END
				) OVER (ORDER BY l.transaction_date ASC, l.created_at ASC) AS running_balance
			FROM ledgers l
			JOIN chart_of_accounts ca ON ca.id = l.coa_id
			WHERE l.transaction_date BETWEEN ? AND ?
		)
		SELECT * FROM RunningData
		ORDER BY transaction_date ASC, created_at ASC
		LIMIT ? OFFSET ?
	`

	err := r.DB.WithContext(ctx).Raw(q,
		startDate,
		startDate,
		endDate,
		limit,
		offset,
	).Scan(&results).Error

	if err != nil {
		return 0, nil, err
	}

	return totalRows, results, nil
}
func (r *LedgerRepository) TransactionSummary(ctx context.Context, req filter.DynamicFilter) (domain.TransactionSummary, error) {

}
func (r *LedgerRepository) CashFlowStatement(ctx context.Context, req filter.DynamicFilter)
