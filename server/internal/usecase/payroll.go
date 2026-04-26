package usecase

import (
	"context"

	"github.com/BramAristyo/saas-pos-core/server/internal/api/dto"
	"github.com/BramAristyo/saas-pos-core/server/internal/repository"
	"github.com/BramAristyo/saas-pos-core/server/pkg/filter"
)

type PayrollUseCase struct {
	Repo *repository.PayrollRepository
}

func NewPayrollUseCase(repo *repository.PayrollRepository) *PayrollUseCase {
	return &PayrollUseCase{
		Repo: repo,
	}
}

func (u *PayrollUseCase) Paginate(ctx context.Context, req filter.PaginationWithInputFilter) (dto.PayrollResponsePagination, error) {
	totalRows, ps, err := u.Repo.Paginate(ctx, req)
	if err != nil {
		return dto.PayrollResponsePagination{}, err
	}

	return dto.ToPayrollResponsePagination(ps, req, totalRows), nil
}

func (u *PayrollUseCase) Store(ctx context.Context, req dto.CreatePayrollRequest) (dto.PayrollResponse, error) {
	payrollDomain, err := dto.ToCreatePayrollModel(&req)
	if err != nil {
		return dto.PayrollResponse{}, err
	}

	_, err = u.Repo.Store(ctx, payrollDomain)
	if err != nil {
		return dto.PayrollResponse{}, err
	}

	return dto.ToPayrollResponse(payrollDomain), nil
}
