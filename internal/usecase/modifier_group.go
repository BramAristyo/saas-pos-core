package usecase

import (
	"context"

	"github.com/BramAristyo/go-pos-mawish/internal/api/dto"
	"github.com/BramAristyo/go-pos-mawish/internal/repository"
	"github.com/BramAristyo/go-pos-mawish/pkg/filter"
	"github.com/google/uuid"
)

type ModifierGroupUseCase struct {
	Repo *repository.ModifierGroupRepository
}

func NewModifierGroupUseCase(repo *repository.ModifierGroupRepository) *ModifierGroupUseCase {
	return &ModifierGroupUseCase{
		Repo: repo,
	}
}

func (s *ModifierGroupUseCase) Paginate(ctx context.Context, req filter.PaginationWithInputFilter) (dto.ModifierGroupResponsePagination, error) {
	totalRows, mgs, err := s.Repo.Paginate(ctx, req)
	if err != nil {
		return dto.ModifierGroupResponsePagination{}, err
	}

	modifierGroups := make([]dto.ModifierGroupResponse, 0, len(mgs))
	for _, mg := range mgs {
		modifierGroups = append(modifierGroups, dto.ToModifierGroupResponse(mg))
	}

	return dto.ToModifierGroupResponsePagination(modifierGroups, req, totalRows), nil
}

func (s *ModifierGroupUseCase) FindById(ctx context.Context, id uuid.UUID) (dto.ModifierGroupResponse, error) {
	mg, err := s.Repo.FindById(ctx, id)
	if err != nil {
		return dto.ModifierGroupResponse{}, err
	}

	return dto.ToModifierGroupResponse(*mg), nil
}

func (s *ModifierGroupUseCase) Store(ctx context.Context, req dto.CreateModifierGroupRequest) (dto.ModifierGroupResponse, error) {
	mg := dto.ToModifierGroupModel(req)
	if _, err := s.Repo.Store(ctx, &mg); err != nil {
		return dto.ModifierGroupResponse{}, nil
	}

	return dto.ToModifierGroupResponse(mg), nil
}

func (s *ModifierGroupUseCase) Update(ctx context.Context, id uuid.UUID, req dto.UpdateModifierGroupRequest) (dto.ModifierGroupResponse, error) {
	mg := dto.ToUpdateModifierGroupModel(req)
	updated, err := s.Repo.Update(ctx, id, &mg)
	if err != nil {
		return dto.ModifierGroupResponse{}, nil
	}

	return dto.ToModifierGroupResponse(*updated), nil
}

func (s *ModifierGroupUseCase) UpdateStatus(ctx context.Context, id uuid.UUID, status bool) (dto.ModifierGroupResponse, error) {
	updated, err := s.Repo.UpdateStatus(ctx, id, status)
	if err != nil {
		return dto.ModifierGroupResponse{}, err
	}

	return dto.ToModifierGroupResponse(*updated), nil
}
