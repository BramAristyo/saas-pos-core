package usecase

import (
	"context"
	"fmt"

	"github.com/BramAristyo/saas-pos-core/server/internal/api/dto"
	"github.com/BramAristyo/saas-pos-core/server/internal/domain"
	"github.com/BramAristyo/saas-pos-core/server/internal/repository"
	"github.com/BramAristyo/saas-pos-core/server/pkg/filter"
	"github.com/BramAristyo/saas-pos-core/server/pkg/helper"
	"github.com/BramAristyo/saas-pos-core/server/pkg/usecase_errors"
	"github.com/google/uuid"
)

type ExpenseUseCase struct {
	Repo       *repository.ExpenseRepository
	LedgerRepo *repository.LedgerRepository
	LogUseCase *AuditLogUseCase
}

func NewExpenseUseCase(repo *repository.ExpenseRepository, ledgerRepo *repository.LedgerRepository, log *AuditLogUseCase) *ExpenseUseCase {
	return &ExpenseUseCase{
		Repo:       repo,
		LedgerRepo: ledgerRepo,
		LogUseCase: log,
	}
}

func (u *ExpenseUseCase) Paginate(ctx context.Context, req filter.PaginationWithInputFilter) (dto.ExpenseResponsePagination, error) {
	totalRows, expenses, err := u.Repo.Paginate(ctx, req)
	if err != nil {
		return dto.ExpenseResponsePagination{}, err
	}

	responses := make([]dto.ExpenseResponse, 0, len(expenses))
	for i := range expenses {
		responses = append(responses, dto.ToExpenseResponse(&expenses[i]))
	}

	return dto.ToExpenseResponsePagination(responses, req, totalRows), nil
}

func (u *ExpenseUseCase) FindById(ctx context.Context, id uuid.UUID) (dto.ExpenseResponse, error) {
	expense, err := u.Repo.FindById(ctx, id)
	if err != nil {
		return dto.ExpenseResponse{}, err
	}

	return dto.ToExpenseResponse(&expense), nil
}

func (u *ExpenseUseCase) Store(ctx context.Context, req dto.CreateExpenseRequest) (dto.ExpenseResponse, error) {
	userId, err := helper.ExtractUserID(ctx)
	if err != nil {
		return dto.ExpenseResponse{}, err
	}

	expense, err := dto.ToExpenseModel(&req)
	if err != nil {
		return dto.ExpenseResponse{}, err
	}

	stored, err := u.Repo.Store(ctx, &expense)
	if err != nil {
		return dto.ExpenseResponse{}, err
	}

	_, err = u.LedgerRepo.Store(ctx, domain.ToLedgerModel(stored))
	if err != nil {
		return dto.ExpenseResponse{}, usecase_errors.LedgerRecordFailed
	}

	go u.LogUseCase.Log(context.Background(), domain.AuditLog{
		UserID:      userId,
		Action:      domain.ActionCreate,
		Entity:      domain.EntityExpense,
		EntityID:    &stored.ID,
		Description: fmt.Sprintf("User created a new expense: %s (Amount: %s)", stored.Description, stored.Amount),
	})

	return dto.ToExpenseResponse(&stored), nil
}

func (u *ExpenseUseCase) Update(ctx context.Context, id uuid.UUID, req dto.UpdateExpenseRequest) (dto.ExpenseResponse, error) {
	userId, err := helper.ExtractUserID(ctx)
	if err != nil {
		return dto.ExpenseResponse{}, err
	}

	expense, err := dto.ToUpdateExpenseModel(&req)
	if err != nil {
		return dto.ExpenseResponse{}, err
	}

	updated, err := u.Repo.Update(ctx, id, &expense)
	if err != nil {
		return dto.ExpenseResponse{}, err
	}

	_, err = u.LedgerRepo.ExpenseUpdate(ctx, id, domain.ToLedgerModel(updated))
	if err != nil {
		return dto.ExpenseResponse{}, usecase_errors.LedgerRecordFailed
	}

	go u.LogUseCase.Log(context.Background(), domain.AuditLog{
		UserID:      userId,
		Action:      domain.ActionUpdate,
		Entity:      domain.EntityExpense,
		EntityID:    &updated.ID,
		Description: fmt.Sprintf("User updated expense: %s (Amount: %s)", updated.Description, updated.Amount),
	})

	return dto.ToExpenseResponse(&updated), nil
}

func (u *ExpenseUseCase) Delete(ctx context.Context, id uuid.UUID) error {
	userId, err := helper.ExtractUserID(ctx)
	if err != nil {
		return err
	}

	expense, err := u.Repo.FindById(ctx, id)
	if err != nil {
		return err
	}

	if err := u.Repo.Delete(ctx, id); err != nil {
		return err
	}

	err = u.LedgerRepo.ExpenseDelete(ctx, id)
	if err != nil {
		return usecase_errors.LedgerRecordFailed
	}

	go u.LogUseCase.Log(context.Background(), domain.AuditLog{
		UserID:      userId,
		Action:      domain.ActionDelete,
		Entity:      domain.EntityExpense,
		EntityID:    &id,
		Description: fmt.Sprintf("User deleted expense: %s", expense.Description),
	})

	return nil
}

func (u *ExpenseUseCase) Restore(ctx context.Context, id uuid.UUID) (dto.ExpenseResponse, error) {
	userId, err := helper.ExtractUserID(ctx)
	if err != nil {
		return dto.ExpenseResponse{}, err
	}

	if err := u.Repo.Restore(ctx, id); err != nil {
		return dto.ExpenseResponse{}, err
	}

	expense, err := u.Repo.FindById(ctx, id)
	if err != nil {
		return dto.ExpenseResponse{}, err
	}

	go u.LogUseCase.Log(context.Background(), domain.AuditLog{
		UserID:      userId,
		Action:      domain.ActionRestore,
		Entity:      domain.EntityExpense,
		EntityID:    &id,
		Description: fmt.Sprintf("User restored expense: %s", expense.Description),
	})

	return dto.ToExpenseResponse(&expense), nil
}
