package usecase

import (
	"context"

	"github.com/BramAristyo/go-pos-mawish/internal/api/dto"
	"github.com/BramAristyo/go-pos-mawish/internal/repository"
	"github.com/BramAristyo/go-pos-mawish/pkg/filter"
	"github.com/BramAristyo/go-pos-mawish/pkg/usecase_errors"
	"github.com/google/uuid"
)

type BundlingUseCase struct {
	Repo *repository.BundlingRepository
}

func NewBundlingUseCase(r *repository.BundlingRepository) *BundlingUseCase {
	return &BundlingUseCase{Repo: r}
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

	return dto.ToBundlingPackageResponse(bp), nil
}

func (u *BundlingUseCase) Store(ctx context.Context, req dto.CreateBundlingPackageRequest) (dto.BundlingPackageResponse, error) {
	bp := dto.ToBundlingPackageModel(req)

	if _, err := u.Repo.Store(ctx, &bp); err != nil {
		if usecase_errors.IsUniqueViolation(err) {
			return dto.BundlingPackageResponse{}, usecase_errors.DuplicateEntry
		}
		return dto.BundlingPackageResponse{}, nil
	}

	return dto.ToBundlingPackageResponse(bp), nil
}

func (u *BundlingUseCase) Update(ctx context.Context, id uuid.UUID, req dto.UpdateBundlingPackageRequest) (dto.BundlingPackageResponse, error) {
	bp := dto.ToUpdateBundlingPackageModel(req)
	updated, err := u.Repo.Update(ctx, id, &bp)
	if err != nil {
		return dto.BundlingPackageResponse{}, err
	}

	return dto.ToBundlingPackageResponse(updated), nil
}

func (u *BundlingUseCase) UpdateStatus(ctx context.Context, id uuid.UUID, status bool) (dto.BundlingPackageResponse, error) {
	bp, err := u.Repo.UpdateStatus(ctx, id, status)
	if err != nil {
		return dto.BundlingPackageResponse{}, err
	}

	return dto.ToBundlingPackageResponse(bp), nil
}
