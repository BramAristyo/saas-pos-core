package usecase

import (
	"context"

	"github.com/BramAristyo/go-pos-mawish/internal/api/dto"
	"github.com/BramAristyo/go-pos-mawish/internal/domain"
	"github.com/BramAristyo/go-pos-mawish/internal/repository"
	"github.com/BramAristyo/go-pos-mawish/pkg/filter"
	"github.com/BramAristyo/go-pos-mawish/pkg/helper"
	"github.com/BramAristyo/go-pos-mawish/pkg/usecase_errors"
	"github.com/google/uuid"
)

type SalesTypeUseCase struct {
	Repo       *repository.SalesTypeRepository
	LogUseCase *AuditLogUseCase
}

func NewSalesTypeUseCase(r *repository.SalesTypeRepository, log *AuditLogUseCase) *SalesTypeUseCase {
	return &SalesTypeUseCase{
		Repo:       r,
		LogUseCase: log,
	}
}

func (u *SalesTypeUseCase) Paginate(ctx context.Context, req filter.PaginationWithInputFilter) (dto.SalesTypeResponsePagination, error) {
	totalRows, salesTypes, err := u.Repo.Paginate(ctx, req)
	if err != nil {
		return dto.SalesTypeResponsePagination{}, err
	}

	responses := make([]dto.SalesTypeResponse, 0, len(salesTypes))
	for i := range salesTypes {
		responses = append(responses, dto.ToSalesTypeResponse(&salesTypes[i]))
	}

	return dto.ToSalesTypeResponsePagination(responses, req, totalRows), nil
}

func (u *SalesTypeUseCase) FindById(ctx context.Context, id uuid.UUID) (dto.SalesTypeResponse, error) {
	salesType, err := u.Repo.FindById(ctx, id)
	if err != nil {
		return dto.SalesTypeResponse{}, err
	}

	return dto.ToSalesTypeResponse(&salesType), nil
}

func (u *SalesTypeUseCase) Store(ctx context.Context, req dto.CreateSalesTypeRequest) (dto.SalesTypeResponse, error) {
	userId, _ := helper.ExtractUserID(ctx)
	salesType := dto.ToCreateSalesTypeModel(&req)

	created, err := u.Repo.Store(ctx, &salesType)
	if err != nil {
		if usecase_errors.IsUniqueViolation(err) {
			return dto.SalesTypeResponse{}, usecase_errors.DuplicateEntry
		}
		return dto.SalesTypeResponse{}, err
	}

	go u.LogUseCase.Log(context.Background(), domain.AuditLog{
		UserID:      userId,
		Action:      domain.ActionCreate,
		Entity:      domain.EntitySalesType,
		EntityID:    &created.ID,
		Description: "User created a new sales type: " + created.Name,
	})

	return dto.ToSalesTypeResponse(&created), nil
}

func (u *SalesTypeUseCase) Update(ctx context.Context, id uuid.UUID, req dto.UpdateSalesTypeRequest) (dto.SalesTypeResponse, error) {
	userId, _ := helper.ExtractUserID(ctx)
	salesType := dto.ToUpdateSalesTypeModel(&req)

	updated, err := u.Repo.SmartUpdate(ctx, id, &salesType)
	if err != nil {
		if usecase_errors.IsUniqueViolation(err) {
			return dto.SalesTypeResponse{}, usecase_errors.DuplicateEntry
		}
		return dto.SalesTypeResponse{}, err
	}

	go u.LogUseCase.Log(context.Background(), domain.AuditLog{
		UserID:      userId,
		Action:      domain.ActionUpdate,
		Entity:      domain.EntitySalesType,
		EntityID:    &updated.ID,
		Description: "User updated sales type: " + updated.Name,
	})

	return dto.ToSalesTypeResponse(&updated), nil
}

func (u *SalesTypeUseCase) UpdateStatus(ctx context.Context, id uuid.UUID, status bool) (dto.SalesTypeResponse, error) {
	userId, _ := helper.ExtractUserID(ctx)
	updated, err := u.Repo.UpdateStatus(ctx, id, status)
	if err != nil {
		return dto.SalesTypeResponse{}, err
	}

	action := domain.ActionActivate
	desc := "User activated sales type: " + updated.Name
	if !status {
		action = domain.ActionDeactivate
		desc = "User deactivated sales type: " + updated.Name
	}

	go u.LogUseCase.Log(context.Background(), domain.AuditLog{
		UserID:      userId,
		Action:      action,
		Entity:      domain.EntitySalesType,
		EntityID:    &updated.ID,
		Description: desc,
	})

	return dto.ToSalesTypeResponse(&updated), nil
}
