package usecase

import (
	"context"

	"github.com/BramAristyo/saas-pos-core/server/internal/api/dto"
	"github.com/BramAristyo/saas-pos-core/server/internal/domain"
	"github.com/BramAristyo/saas-pos-core/server/internal/repository"
	"github.com/BramAristyo/saas-pos-core/server/pkg/filter"
	"github.com/BramAristyo/saas-pos-core/server/pkg/helper"
	"github.com/BramAristyo/saas-pos-core/server/pkg/usecase_errors"
	"github.com/google/uuid"
)

type BundlingUseCase struct {
	Repo       *repository.BundlingRepository
	LogUseCase *AuditLogUseCase
}

func NewBundlingUseCase(r *repository.BundlingRepository, log *AuditLogUseCase) *BundlingUseCase {
	return &BundlingUseCase{
		Repo:       r,
		LogUseCase: log,
	}
}

func (u *BundlingUseCase) Paginate(ctx context.Context, req filter.PaginationWithInputFilter) (dto.BundlingPackagePaginationResponse, error) {
	totalRows, bps, err := u.Repo.Paginate(ctx, req)
	if err != nil {
		return dto.BundlingPackagePaginationResponse{}, err
	}

	return dto.ToBundlingPackagePaginationResponse(bps, req, totalRows), nil
}

func (u *BundlingUseCase) FindById(ctx context.Context, id uuid.UUID) (dto.BundlingPackageResponse, error) {
	bp, err := u.Repo.FindById(ctx, id)
	if err != nil {
		return dto.BundlingPackageResponse{}, err
	}

	return dto.ToBundlingPackageResponse(&bp), nil
}

func (u *BundlingUseCase) Store(ctx context.Context, req dto.CreateBundlingPackageRequest) (dto.BundlingPackageResponse, error) {
	userId, _ := helper.ExtractUserID(ctx)
	bp := dto.ToBundlingPackageModel(&req)

	stored, err := u.Repo.Store(ctx, &bp)
	if err != nil {
		if usecase_errors.IsUniqueViolation(err) {
			return dto.BundlingPackageResponse{}, usecase_errors.DuplicateEntry
		}
		return dto.BundlingPackageResponse{}, err
	}

	go u.LogUseCase.Log(context.Background(), domain.AuditLog{
		UserID:      userId,
		Action:      domain.ActionCreate,
		Entity:      domain.EntityBundling,
		EntityID:    &stored.ID,
		Description: "User created a new bundling package: " + stored.Name,
	})

	return dto.ToBundlingPackageResponse(&stored), nil
}

func (u *BundlingUseCase) Update(ctx context.Context, id uuid.UUID, req dto.UpdateBundlingPackageRequest) (dto.BundlingPackageResponse, error) {
	userId, _ := helper.ExtractUserID(ctx)
	bp := dto.ToUpdateBundlingPackageModel(&req)
	updated, err := u.Repo.Update(ctx, id, &bp)
	if err != nil {
		return dto.BundlingPackageResponse{}, err
	}

	go u.LogUseCase.Log(context.Background(), domain.AuditLog{
		UserID:      userId,
		Action:      domain.ActionUpdate,
		Entity:      domain.EntityBundling,
		EntityID:    &updated.ID,
		Description: "User updated bundling package: " + updated.Name,
	})

	return dto.ToBundlingPackageResponse(&updated), nil
}

func (u *BundlingUseCase) Delete(ctx context.Context, id uuid.UUID) error {
	userId, _ := helper.ExtractUserID(ctx)

	bp, err := u.Repo.FindById(ctx, id)
	if err != nil {
		return err
	}

	if err := u.Repo.Delete(ctx, id); err != nil {
		return err
	}

	go u.LogUseCase.Log(context.Background(), domain.AuditLog{
		UserID:      userId,
		Action:      domain.ActionDelete,
		Entity:      domain.EntityBundling,
		EntityID:    &id,
		Description: "User deleted bundling package: " + bp.Name,
	})

	return nil
}

func (u *BundlingUseCase) Restore(ctx context.Context, id uuid.UUID) (dto.BundlingPackageResponse, error) {
	userId, _ := helper.ExtractUserID(ctx)

	if err := u.Repo.Restore(ctx, id); err != nil {
		return dto.BundlingPackageResponse{}, err
	}

	bp, err := u.Repo.FindById(ctx, id)
	if err != nil {
		return dto.BundlingPackageResponse{}, err
	}

	go u.LogUseCase.Log(context.Background(), domain.AuditLog{
		UserID:      userId,
		Action:      domain.ActionRestore,
		Entity:      domain.EntityBundling,
		EntityID:    &id,
		Description: "User restored bundling package: " + bp.Name,
	})

	return dto.ToBundlingPackageResponse(&bp), nil
}
