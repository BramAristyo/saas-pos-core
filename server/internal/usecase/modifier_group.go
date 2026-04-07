package usecase

import (
	"context"

	"github.com/BramAristyo/saas-pos-core/server/internal/api/dto"
	"github.com/BramAristyo/saas-pos-core/server/internal/domain"
	"github.com/BramAristyo/saas-pos-core/server/internal/repository"
	"github.com/BramAristyo/saas-pos-core/server/pkg/filter"
	"github.com/BramAristyo/saas-pos-core/server/pkg/helper"
	"github.com/google/uuid"
)

type ModifierGroupUseCase struct {
	Repo       *repository.ModifierGroupRepository
	LogUseCase *AuditLogUseCase
}

func NewModifierGroupUseCase(repo *repository.ModifierGroupRepository, log *AuditLogUseCase) *ModifierGroupUseCase {
	return &ModifierGroupUseCase{
		Repo:       repo,
		LogUseCase: log,
	}
}

func (u *ModifierGroupUseCase) Paginate(ctx context.Context, req filter.PaginationWithInputFilter) (dto.ModifierGroupResponsePagination, error) {
	totalRows, mgs, err := u.Repo.Paginate(ctx, req)
	if err != nil {
		return dto.ModifierGroupResponsePagination{}, err
	}

	modifierGroups := make([]dto.ModifierGroupResponse, 0, len(mgs))
	for i := range mgs {
		modifierGroups = append(modifierGroups, dto.ToModifierGroupResponse(&mgs[i]))
	}

	return dto.ToModifierGroupResponsePagination(modifierGroups, req, totalRows), nil
}

func (u *ModifierGroupUseCase) FindById(ctx context.Context, id uuid.UUID) (dto.ModifierGroupResponse, error) {
	mg, err := u.Repo.FindById(ctx, id)
	if err != nil {
		return dto.ModifierGroupResponse{}, err
	}

	return dto.ToModifierGroupResponse(&mg), nil
}

func (u *ModifierGroupUseCase) Store(ctx context.Context, req dto.CreateModifierGroupRequest) (dto.ModifierGroupResponse, error) {
	userId, err := helper.ExtractUserID(ctx)
	if err != nil {
		return dto.ModifierGroupResponse{}, err
	}

	mg := dto.ToModifierGroupModel(&req)
	stored, err := u.Repo.Store(ctx, &mg)
	if err != nil {
		return dto.ModifierGroupResponse{}, err
	}

	go u.LogUseCase.Log(context.Background(), domain.AuditLog{
		UserID:      userId,
		Action:      domain.ActionCreate,
		Entity:      domain.EntityModifierGroup,
		EntityID:    &stored.ID,
		Description: "User created a new modifier group: " + stored.Name,
	})

	return dto.ToModifierGroupResponse(&stored), nil
}

func (u *ModifierGroupUseCase) Update(ctx context.Context, id uuid.UUID, req dto.UpdateModifierGroupRequest) (dto.ModifierGroupResponse, error) {
	userId, err := helper.ExtractUserID(ctx)
	if err != nil {
		return dto.ModifierGroupResponse{}, err
	}

	mg := dto.ToUpdateModifierGroupModel(&req)
	updated, err := u.Repo.Update(ctx, id, &mg)
	if err != nil {
		return dto.ModifierGroupResponse{}, err
	}

	go u.LogUseCase.Log(context.Background(), domain.AuditLog{
		UserID:      userId,
		Action:      domain.ActionUpdate,
		Entity:      domain.EntityModifierGroup,
		EntityID:    &updated.ID,
		Description: "User updated modifier group: " + updated.Name,
	})

	return dto.ToModifierGroupResponse(&updated), nil
}

func (u *ModifierGroupUseCase) Delete(ctx context.Context, id uuid.UUID) error {
	userId, err := helper.ExtractUserID(ctx)
	if err != nil {
		return err
	}

	mg, err := u.Repo.FindById(ctx, id)
	if err != nil {
		return err
	}

	if err := u.Repo.Delete(ctx, id); err != nil {
		return err
	}

	go u.LogUseCase.Log(context.Background(), domain.AuditLog{
		UserID:      userId,
		Action:      domain.ActionDelete,
		Entity:      domain.EntityModifierGroup,
		EntityID:    &id,
		Description: "User deleted modifier group: " + mg.Name,
	})

	return nil
}

func (u *ModifierGroupUseCase) Restore(ctx context.Context, id uuid.UUID) (dto.ModifierGroupResponse, error) {
	userId, err := helper.ExtractUserID(ctx)
	if err != nil {
		return dto.ModifierGroupResponse{}, err
	}

	if err := u.Repo.Restore(ctx, id); err != nil {
		return dto.ModifierGroupResponse{}, err
	}

	mg, err := u.Repo.FindById(ctx, id)
	if err != nil {
		return dto.ModifierGroupResponse{}, err
	}

	go u.LogUseCase.Log(context.Background(), domain.AuditLog{
		UserID:      userId,
		Action:      domain.ActionRestore,
		Entity:      domain.EntityModifierGroup,
		EntityID:    &id,
		Description: "User restored modifier group: " + mg.Name,
	})

	return dto.ToModifierGroupResponse(&mg), nil
}
