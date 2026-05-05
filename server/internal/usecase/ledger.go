package usecase

import (
	"context"

	"github.com/BramAristyo/saas-pos-core/server/internal/api/dto"
	"github.com/BramAristyo/saas-pos-core/server/internal/repository"
	"github.com/BramAristyo/saas-pos-core/server/pkg/filter"
)

type LedgerUseCase struct {
	Repo *repository.LedgerRepository
}

func NewLedgerUseCase(repo *repository.LedgerRepository) *LedgerUseCase {
	return &LedgerUseCase{
		Repo: repo,
	}
}

func (u *LedgerUseCase) TransactionList(ctx context.Context, req filter.PaginationWithInputFilter) (dto.TransactionSummaryResponse, error)
func (u *LedgerUseCase) CashFlowStatement(ctx context.Context, req filter.PaginationWithInputFilter) (dto.CashFlowReportResponse, error)
