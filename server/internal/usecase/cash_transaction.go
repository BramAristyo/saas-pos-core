package usecase

import (
	"context"

	"github.com/BramAristyo/saas-pos-core/server/internal/api/dto"
	"github.com/BramAristyo/saas-pos-core/server/internal/repository"
	"github.com/BramAristyo/saas-pos-core/server/pkg/filter"
)

type CashTransactionUseCase struct {
	Repo *repository.CashTransactionRepository
}

func (u *CashTransactionUseCase) Paginate(ctx context.Context, req filter.PaginationWithInputFilter) (dto.CashTransactionResponsePagination, error) {
	totalRows, cts, err := u.Repo.Paginate(ctx, req)
	if err != nil {
		return dto.CashTransactionResponsePagination{}, err
	}

	return dto.ToCashTransactionPagination(cts, req, totalRows), nil
}
