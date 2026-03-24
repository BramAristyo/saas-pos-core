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

type DiscountUseCase struct {
	Repo       *repository.DiscountRepository
	LogUseCase *AuditLogUseCase
}

func NewDiscountUseCase(repo *repository.DiscountRepository, log *AuditLogUseCase) *DiscountUseCase {
	return &DiscountUseCase{
		Repo:       repo,
		LogUseCase: log,
	}
}

func (u *DiscountUseCase) Paginate(ctx context.Context, req filter.PaginationWithInputFilter) (dto.DiscountResponsePagination, error) {
	totalRows, discounts, err := u.Repo.Paginate(ctx, req)
	if err != nil {
		return dto.DiscountResponsePagination{}, err
	}

	discountResponses := dto.ToDiscountResponses(discounts)

	return dto.ToDiscountResponsePagination(discountResponses, req, totalRows), nil
}

func (u *DiscountUseCase) FindById(ctx context.Context, id uuid.UUID) (dto.DiscountResponse, error) {
	discount, err := u.Repo.FindById(ctx, id)
	if err != nil {
		return dto.DiscountResponse{}, err
	}

	return dto.ToDiscountResponse(&discount), nil
}

func (u *DiscountUseCase) Store(ctx context.Context, req dto.CreateDiscountRequest) (dto.DiscountResponse, error) {
	userId, _ := helper.ExtractUserID(ctx)
	discount := dto.ToCreateDiscountModel(&req)

	stored, err := u.Repo.Store(ctx, &discount)
	if err != nil {
		if usecase_errors.IsUniqueViolation(err) {
			return dto.DiscountResponse{}, usecase_errors.DuplicateEntry
		}
		return dto.DiscountResponse{}, err
	}

	go u.LogUseCase.Log(context.Background(), domain.AuditLog{
		UserID:      userId,
		Action:      domain.ActionCreate,
		Entity:      domain.EntityDiscount,
		EntityID:    &stored.ID,
		Description: "User created a new discount: " + stored.Name,
	})

	return dto.ToDiscountResponse(&stored), nil
}

func (u *DiscountUseCase) Update(ctx context.Context, id uuid.UUID, req dto.UpdateDiscountRequest) (dto.DiscountResponse, error) {
	userId, _ := helper.ExtractUserID(ctx)
	discount := dto.ToUpdateDiscountModel(&req)
	updated, err := u.Repo.Update(ctx, id, &discount)
	if err != nil {
		if usecase_errors.IsUniqueViolation(err) {
			return dto.DiscountResponse{}, usecase_errors.DuplicateEntry
		}
		return dto.DiscountResponse{}, err
	}

	go u.LogUseCase.Log(context.Background(), domain.AuditLog{
		UserID:      userId,
		Action:      domain.ActionUpdate,
		Entity:      domain.EntityDiscount,
		EntityID:    &updated.ID,
		Description: "User updated discount: " + updated.Name,
	})

	return dto.ToDiscountResponse(&updated), nil
}

func (u *DiscountUseCase) UpdateStatus(ctx context.Context, id uuid.UUID, status bool) (dto.DiscountResponse, error) {
	userId, _ := helper.ExtractUserID(ctx)
	discount, err := u.Repo.UpdateStatus(ctx, id, status)
	if err != nil {
		return dto.DiscountResponse{}, err
	}

	action := domain.ActionActivate
	desc := "User activated discount: " + discount.Name
	if !status {
		action = domain.ActionDeactivate
		desc = "User deactivated discount: " + discount.Name
	}

	go u.LogUseCase.Log(context.Background(), domain.AuditLog{
		UserID:      userId,
		Action:      action,
		Entity:      domain.EntityDiscount,
		EntityID:    &discount.ID,
		Description: desc,
	})

	return dto.ToDiscountResponse(&discount), nil
}
