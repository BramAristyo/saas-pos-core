package dto

import (
	"github.com/BramAristyo/saas-pos-core/backend/internal/domain"
	"github.com/BramAristyo/saas-pos-core/backend/pkg/filter"
	"github.com/google/uuid"
)

type ModifierGroupResponse struct {
	ID         uuid.UUID                `json:"id"`
	Name       string                   `json:"name"`
	IsRequired bool                     `json:"isRequired"`
	Options    []ModifierOptionResponse `json:"options,omitempty"`
	DeletedAt  *string                  `json:"deletedAt,omitempty"`
	CreatedAt  string                   `json:"createdAt"`
	UpdatedAt  string                   `json:"updatedAt"`
}

type ModifierGroupResponsePagination struct {
	Data []ModifierGroupResponse `json:"data"`
	Meta filter.Meta             `json:"meta"`
}

type CreateModifierGroupRequest struct {
	Name       string `json:"name" binding:"required,min=3,max=100"`
	IsRequired bool   `json:"isRequired"`
}

type UpdateModifierGroupRequest struct {
	Name       string `json:"name" binding:"required,min=3,max=100"`
	IsRequired bool   `json:"isRequired"`
}

func ToModifierGroupResponse(mg *domain.ModifierGroup) ModifierGroupResponse {
	var options []ModifierOptionResponse
	if len(mg.ModifierOptions) > 0 {
		for _, o := range mg.ModifierOptions {
			options = append(options, ToModifierOptionResponse(&o))
		}
	}

	resp := ModifierGroupResponse{
		ID:         mg.ID,
		Name:       mg.Name,
		IsRequired: mg.IsRequired,
		Options:    options,
		CreatedAt:  mg.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt:  mg.UpdatedAt.Format("2006-01-02 15:04:05"),
	}

	if mg.DeletedAt.Valid {
		at := mg.DeletedAt.Time.Format("2006-01-02 15:04:05")
		resp.DeletedAt = &at
	}

	return resp
}

func ToModifierGroupResponsePagination(mg []ModifierGroupResponse, f filter.PaginationWithInputFilter, totalRows int64) ModifierGroupResponsePagination {
	return ModifierGroupResponsePagination{
		Data: mg,
		Meta: f.ToMeta(totalRows),
	}
}

func ToModifierGroupModel(req *CreateModifierGroupRequest) domain.ModifierGroup {
	return domain.ModifierGroup{
		Name:       req.Name,
		IsRequired: req.IsRequired,
	}
}

func ToUpdateModifierGroupModel(req *UpdateModifierGroupRequest) domain.ModifierGroup {
	return domain.ModifierGroup{
		Name:       req.Name,
		IsRequired: req.IsRequired,
	}
}
