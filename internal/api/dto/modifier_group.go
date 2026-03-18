package dto

import (
	"github.com/BramAristyo/go-pos-mawish/internal/domain"
	"github.com/BramAristyo/go-pos-mawish/pkg/filter"
	"github.com/google/uuid"
)

type ModifierGroupResponse struct {
	ID         uuid.UUID                `json:"id"`
	Name       string                   `json:"name"`
	IsRequired bool                     `json:"is_required"`
	IsActive   bool                     `json:"is_active"`
	Options    []ModifierOptionResponse `json:"options,omitempty"`
	CreatedAt  string                   `json:"created_at"`
	UpdatedAt  string                   `json:"updated_at"`
}

type ModifierGroupResponsePagination struct {
	Data []ModifierGroupResponse `json:"data"`
	Meta filter.Meta             `json:"meta"`
}

type CreateModifierGroupRequest struct {
	Name       string `json:"name" binding:"required,min=3,max=100"`
	IsRequired bool   `json:"is_required"`
}

type UpdateModifierGroupRequest struct {
	Name       string `json:"name" binding:"required,min=3,max=100"`
	IsRequired bool   `json:"is_required"`
	IsActive   bool   `json:"is_active"`
}

func ToModifierGroupResponse(mg domain.ModifierGroup) ModifierGroupResponse {
	var options []ModifierOptionResponse
	if len(mg.ModifierOptions) > 0 {
		for _, o := range mg.ModifierOptions {
			options = append(options, ToModifierOptionResponse(o))
		}
	}

	return ModifierGroupResponse{
		ID:         mg.ID,
		Name:       mg.Name,
		IsRequired: mg.IsRequired,
		IsActive:   mg.IsActive,
		Options:    options,
		CreatedAt:  mg.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt:  mg.UpdatedAt.Format("2006-01-02 15:04:05"),
	}
}

func ToModifierGroupResponsePagination(mg []ModifierGroupResponse, f filter.PaginationWithInputFilter, totalRows int64) ModifierGroupResponsePagination {
	return ModifierGroupResponsePagination{
		Data: mg,
		Meta: f.ToMeta(totalRows),
	}
}

func ToModifierGroupModel(req CreateModifierGroupRequest) domain.ModifierGroup {
	return domain.ModifierGroup{
		Name:       req.Name,
		IsRequired: req.IsRequired,
		IsActive:   true,
	}
}

func ToUpdateModifierGroupModel(req UpdateModifierGroupRequest) domain.ModifierGroup {
	return domain.ModifierGroup{
		Name:       req.Name,
		IsRequired: req.IsRequired,
		IsActive:   req.IsActive,
	}
}
