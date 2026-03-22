package usecase

import (
	"context"

	"github.com/BramAristyo/go-pos-mawish/internal/api/dto"
	"github.com/BramAristyo/go-pos-mawish/internal/repository"
	"github.com/BramAristyo/go-pos-mawish/pkg/filter"
	"github.com/BramAristyo/go-pos-mawish/pkg/usecase_errors"
	"github.com/google/uuid"
)

type SalesTypeUseCase struct {
	Repo *repository.SalesTypeRepository
}

func NewSalesTypeUseCase(r *repository.SalesTypeRepository) *SalesTypeUseCase {
	return &SalesTypeUseCase{Repo: r}
}

func (u *SalesTypeUseCase) Paginate(ctx context.Context, req filter.PaginationWithInputFilter) (dto.SalesTypeResponsePagination, error) {
	totalRows, salesTypes, err := u.Repo.Paginate(ctx, req)
	if err != nil {
		return dto.SalesTypeResponsePagination{}, err
	}

	responses := make([]dto.SalesTypeResponse, 0, len(salesTypes))
	for _, s := range salesTypes {
		responses = append(responses, dto.ToSalesTypeResponse(s))
	}

	return dto.ToSalesTypeResponsePagination(responses, req, totalRows), nil
}

func (u *SalesTypeUseCase) FindById(ctx context.Context, id uuid.UUID) (dto.SalesTypeResponse, error) {
	salesType, err := u.Repo.FindById(ctx, id)
	if err != nil {
		return dto.SalesTypeResponse{}, err
	}

	return dto.ToSalesTypeResponse(salesType), nil
}

func (u *SalesTypeUseCase) Store(ctx context.Context, req dto.CreateSalesTypeRequest) (dto.SalesTypeResponse, error) {
	salesType := dto.ToCreateSalesTypeModel(req)

	created, err := u.Repo.Store(ctx, &salesType)
	if err != nil {
		if usecase_errors.IsUniqueViolation(err) {
			return dto.SalesTypeResponse{}, usecase_errors.DuplicateEntry
		}
		return dto.SalesTypeResponse{}, err
	}

	return dto.ToSalesTypeResponse(created), nil
}

func (u *SalesTypeUseCase) Update(ctx context.Context, id uuid.UUID, req dto.UpdateSalesTypeRequest) (dto.SalesTypeResponse, error) {
	salesType := dto.ToUpdateSalesTypeModel(req)

	updated, err := u.Repo.SmartUpdate(ctx, id, &salesType)
	if err != nil {
		if usecase_errors.IsUniqueViolation(err) {
			return dto.SalesTypeResponse{}, usecase_errors.DuplicateEntry
		}
		return dto.SalesTypeResponse{}, err
	}

	return dto.ToSalesTypeResponse(updated), nil
}

func (u *SalesTypeUseCase) UpdateStatus(ctx context.Context, id uuid.UUID, status bool) (dto.SalesTypeResponse, error) {
	updated, err := u.Repo.UpdateStatus(ctx, id, status)
	if err != nil {
		return dto.SalesTypeResponse{}, err
	}

	return dto.ToSalesTypeResponse(updated), nil
}
