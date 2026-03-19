package usecase

import (
	"context"

	"github.com/BramAristyo/go-pos-mawish/internal/api/dto"
	"github.com/BramAristyo/go-pos-mawish/internal/repository"
	"github.com/BramAristyo/go-pos-mawish/pkg/filter"
	"github.com/google/uuid"
)

type ModifierOptionUseCase struct {
	Repo *repository.ModifierOptionRepository
}

func NewModifierOptionUseCase(repo *repository.ModifierOptionRepository) *ModifierOptionUseCase {
	return &ModifierOptionUseCase{Repo: repo}
}

func (u *ModifierOptionUseCase) Paginate(ctx context.Context, req filter.PaginationWithInputFilter) (dto.ModifierOptionResponsePagination, error) {
	totalRows, mos, err := u.Repo.Paginate(ctx, req)
	if err != nil {
		return dto.ModifierOptionResponsePagination{}, err
	}

	modifierOptions := make([]dto.ModifierOptionResponse, 0, len(mos))
	for _, mo := range mos {
		modifierOptions = append(modifierOptions, dto.ToModifierOptionResponse(mo))
	}

	return dto.ToModifierOptionResponsePagination(modifierOptions, req, totalRows), nil
}

func (u *ModifierOptionUseCase) FindById(ctx context.Context, id uuid.UUID) (dto.ModifierOptionResponse, error) {
	mo, err := u.Repo.FindById(ctx, id)
	if err != nil {
		return dto.ModifierOptionResponse{}, err
	}

	return dto.ToModifierOptionResponse(mo), nil
}

func (u *ModifierOptionUseCase) Store(ctx context.Context, req dto.CreateModifierOptionRequest) (dto.ModifierOptionResponse, error) {
	mo := dto.ToModifierOptionModel(req)
	stored, err := u.Repo.Store(ctx, &mo)
	if err != nil {
		return dto.ModifierOptionResponse{}, err
	}

	return dto.ToModifierOptionResponse(stored), nil
}

func (u *ModifierOptionUseCase) Update(ctx context.Context, id uuid.UUID, req dto.UpdateModifierOptionRequest) (dto.ModifierOptionResponse, error) {
	mo := dto.ToUpdateModifierOptionModel(req)
	updated, err := u.Repo.Update(ctx, id, &mo)
	if err != nil {
		return dto.ModifierOptionResponse{}, err
	}

	return dto.ToModifierOptionResponse(updated), nil
}

func (u *ModifierOptionUseCase) UpdateStatus(ctx context.Context, id uuid.UUID, status bool) (dto.ModifierOptionResponse, error) {
	updated, err := u.Repo.UpdateStatus(ctx, id, status)
	if err != nil {
		return dto.ModifierOptionResponse{}, err
	}

	return dto.ToModifierOptionResponse(updated), nil
}
