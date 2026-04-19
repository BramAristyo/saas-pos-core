package usecase

import (
	"context"

	"github.com/BramAristyo/saas-pos-core/server/internal/api/dto"
	"github.com/BramAristyo/saas-pos-core/server/internal/repository"
	"github.com/BramAristyo/saas-pos-core/server/pkg/filter"
	"github.com/google/uuid"
)

type ShiftExpensesUseCase struct {
	Repo *repository.ShiftExpensesRepository
}

func NewShiftExpensesUseCase(repo *repository.ShiftExpensesRepository) *ShiftExpensesUseCase {
	return &ShiftExpensesUseCase{
		Repo: repo,
	}
}

func (u *ShiftExpensesUseCase) Paginate(ctx context.Context, req filter.PaginationWithInputFilter) (dto.ShiftExpenseResponsePagination, error) {
	totalRows, expenses, err := u.Repo.Paginate(ctx, req)
	if err != nil {
		return dto.ShiftExpenseResponsePagination{}, err
	}

	responses := make([]dto.ShiftExpenseResponse, 0, len(expenses))
	for i := range expenses {
		responses = append(responses, dto.ToShiftExpenseResponse(&expenses[i]))
	}

	return dto.ToShiftExpenseResponsePagination(responses, req, totalRows), nil
}

func (u *ShiftExpensesUseCase) FindById(ctx context.Context, id uuid.UUID) (dto.ShiftExpenseResponse, error) {
	expense, err := u.Repo.FindById(ctx, id)
	if err != nil {
		return dto.ShiftExpenseResponse{}, err
	}

	return dto.ToShiftExpenseResponse(&expense), nil
}
