package usecase

import (
	"context"

	"github.com/BramAristyo/saas-pos-core/backend/internal/api/dto"
	"github.com/BramAristyo/saas-pos-core/backend/internal/domain"
	"github.com/BramAristyo/saas-pos-core/backend/internal/repository"
	"github.com/BramAristyo/saas-pos-core/backend/pkg/filter"
	"github.com/BramAristyo/saas-pos-core/backend/pkg/helper"
	"github.com/google/uuid"
)

type ModifierOptionUseCase struct {
	Repo       *repository.ModifierOptionRepository
	LogUseCase *AuditLogUseCase
}

func NewModifierOptionUseCase(repo *repository.ModifierOptionRepository, log *AuditLogUseCase) *ModifierOptionUseCase {
	return &ModifierOptionUseCase{
		Repo:       repo,
		LogUseCase: log,
	}
}

func (u *ModifierOptionUseCase) Paginate(ctx context.Context, req filter.PaginationWithInputFilter) (dto.ModifierOptionResponsePagination, error) {
	totalRows, mos, err := u.Repo.Paginate(ctx, req)
	if err != nil {
		return dto.ModifierOptionResponsePagination{}, err
	}

	modifierOptions := make([]dto.ModifierOptionResponse, 0, len(mos))
	for i := range mos {
		modifierOptions = append(modifierOptions, dto.ToModifierOptionResponse(&mos[i]))
	}

	return dto.ToModifierOptionResponsePagination(modifierOptions, req, totalRows), nil
}

func (u *ModifierOptionUseCase) FindById(ctx context.Context, id uuid.UUID) (dto.ModifierOptionResponse, error) {
	mo, err := u.Repo.FindById(ctx, id)
	if err != nil {
		return dto.ModifierOptionResponse{}, err
	}

	return dto.ToModifierOptionResponse(&mo), nil
}

func (u *ModifierOptionUseCase) Store(ctx context.Context, req dto.CreateModifierOptionRequest) (dto.ModifierOptionResponse, error) {
	userId, _ := helper.ExtractUserID(ctx)
	mo := dto.ToModifierOptionModel(&req)
	stored, err := u.Repo.Store(ctx, &mo)
	if err != nil {
		return dto.ModifierOptionResponse{}, err
	}

	go u.LogUseCase.Log(context.Background(), domain.AuditLog{
		UserID:      userId,
		Action:      domain.ActionCreate,
		Entity:      domain.EntityModifierOption,
		EntityID:    &stored.ID,
		Description: "User created a new modifier option: " + stored.Name,
	})

	return dto.ToModifierOptionResponse(&stored), nil
}

func (u *ModifierOptionUseCase) Update(ctx context.Context, id uuid.UUID, req dto.UpdateModifierOptionRequest) (dto.ModifierOptionResponse, error) {
	userId, _ := helper.ExtractUserID(ctx)
	mo := dto.ToUpdateModifierOptionModel(&req)
	updated, err := u.Repo.Update(ctx, id, &mo)
	if err != nil {
		return dto.ModifierOptionResponse{}, err
	}

	go u.LogUseCase.Log(context.Background(), domain.AuditLog{
		UserID:      userId,
		Action:      domain.ActionUpdate,
		Entity:      domain.EntityModifierOption,
		EntityID:    &updated.ID,
		Description: "User updated modifier option: " + updated.Name,
	})

	return dto.ToModifierOptionResponse(&updated), nil
}

func (u *ModifierOptionUseCase) Delete(ctx context.Context, id uuid.UUID) error {
	userId, _ := helper.ExtractUserID(ctx)

	mo, err := u.Repo.FindById(ctx, id)
	if err != nil {
		return err
	}

	if err := u.Repo.Delete(ctx, id); err != nil {
		return err
	}

	go u.LogUseCase.Log(context.Background(), domain.AuditLog{
		UserID:      userId,
		Action:      domain.ActionDelete,
		Entity:      domain.EntityModifierOption,
		EntityID:    &id,
		Description: "User deleted modifier option: " + mo.Name,
	})

	return nil
}

func (u *ModifierOptionUseCase) Restore(ctx context.Context, id uuid.UUID) (dto.ModifierOptionResponse, error) {
	userId, _ := helper.ExtractUserID(ctx)

	if err := u.Repo.Restore(ctx, id); err != nil {
		return dto.ModifierOptionResponse{}, err
	}

	mo, err := u.Repo.FindById(ctx, id)
	if err != nil {
		return dto.ModifierOptionResponse{}, err
	}

	go u.LogUseCase.Log(context.Background(), domain.AuditLog{
		UserID:      userId,
		Action:      domain.ActionRestore,
		Entity:      domain.EntityModifierOption,
		EntityID:    &id,
		Description: "User restored modifier option: " + mo.Name,
	})

	return dto.ToModifierOptionResponse(&mo), nil
}
