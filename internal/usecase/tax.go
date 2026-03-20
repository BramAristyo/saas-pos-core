package usecase

import (
	"context"

	"github.com/BramAristyo/go-pos-mawish/internal/api/dto"
	"github.com/BramAristyo/go-pos-mawish/internal/repository"
	"github.com/BramAristyo/go-pos-mawish/pkg/filter"
	"github.com/BramAristyo/go-pos-mawish/pkg/usecase_errors"
	"github.com/google/uuid"
)

type TaxUseCase struct {
	Repo *repository.TaxRepository
}

func NewTaxUseCase(repo *repository.TaxRepository) *TaxUseCase {
	return &TaxUseCase{
		Repo: repo,
	}
}

func (u *TaxUseCase) Paginate(ctx context.Context, req filter.PaginationWithInputFilter) (dto.TaxResponsePagination, error) {
	totalRows, taxes, err := u.Repo.Paginate(ctx, req)
	if err != nil {
		return dto.TaxResponsePagination{}, err
	}

	taxResponses := dto.ToTaxResponses(taxes)

	return dto.ToTaxResponsePagination(taxResponses, req, totalRows), nil
}

func (u *TaxUseCase) FindById(ctx context.Context, id uuid.UUID) (dto.TaxResponse, error) {
	tax, err := u.Repo.FindById(ctx, id)
	if err != nil {
		return dto.TaxResponse{}, err
	}

	return dto.ToTaxResponse(tax), nil
}

func (u *TaxUseCase) Store(ctx context.Context, req dto.CreateTaxRequest) (dto.TaxResponse, error) {
	tax := dto.ToCreateTaxModel(req)

	stored, err := u.Repo.Store(ctx, &tax)
	if err != nil {
		if usecase_errors.IsUniqueViolation(err) {
			return dto.TaxResponse{}, usecase_errors.DuplicateEntry
		}
		return dto.TaxResponse{}, err
	}

	return dto.ToTaxResponse(stored), nil
}

func (u *TaxUseCase) Update(ctx context.Context, id uuid.UUID, req dto.UpdateTaxRequest) (dto.TaxResponse, error) {
	tax := dto.ToUpdateTaxModel(req)
	updated, err := u.Repo.Update(ctx, id, &tax)
	if err != nil {
		if usecase_errors.IsUniqueViolation(err) {
			return dto.TaxResponse{}, usecase_errors.DuplicateEntry
		}
		return dto.TaxResponse{}, err
	}

	return dto.ToTaxResponse(updated), nil
}

func (u *TaxUseCase) UpdateStatus(ctx context.Context, id uuid.UUID, status bool) (dto.TaxResponse, error) {
	if status {
		if err := u.Repo.DeactiveAll(ctx); err != nil {
			return dto.TaxResponse{}, err
		}
	}

	tax, err := u.Repo.UpdateStatus(ctx, id, status)
	if err != nil {
		return dto.TaxResponse{}, err
	}

	return dto.ToTaxResponse(tax), nil
}
