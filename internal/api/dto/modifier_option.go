package dto

import (
	"github.com/BramAristyo/go-pos-mawish/internal/domain"
	"github.com/BramAristyo/go-pos-mawish/pkg/filter"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type ModifierOptionResponse struct {
	ID              uuid.UUID              `json:"id"`
	ModifierGroupID uuid.UUID              `json:"modifierGroupId"`
	ModifierGroup   *ModifierGroupResponse `json:"modifierGroup,omitempty"`
	Name            string                 `json:"name"`
	PriceAdjustment decimal.Decimal        `json:"priceAdjustment"`
	CogsAdjustment  decimal.Decimal        `json:"cogsAdjustment"`
	DeletedAt       *string                `json:"deletedAt,omitempty"`
	CreatedAt       string                 `json:"createdAt"`
	UpdatedAt       string                 `json:"updatedAt"`
}

type ModifierOptionResponsePagination struct {
	Data []ModifierOptionResponse `json:"data"`
	Meta filter.Meta              `json:"meta"`
}

type CreateModifierOptionRequest struct {
	ModifierGroupID uuid.UUID       `json:"modifierGroupId" binding:"required"`
	Name            string          `json:"name" binding:"required,min=3,max=100"`
	PriceAdjustment decimal.Decimal `json:"priceAdjustment" binding:"required"`
	CogsAdjustment  decimal.Decimal `json:"cogsAdjustment" binding:"required"`
}

type UpdateModifierOptionRequest struct {
	ModifierGroupID uuid.UUID       `json:"modifierGroupId" binding:"required"`
	Name            string          `json:"name" binding:"required,min=3,max=100"`
	PriceAdjustment decimal.Decimal `json:"priceAdjustment" binding:"required"`
	CogsAdjustment  decimal.Decimal `json:"cogsAdjustment" binding:"required"`
}

func ToModifierOptionResponse(mo *domain.ModifierOption) ModifierOptionResponse {
	var mg *ModifierGroupResponse
	if mo.ModifierGroup != nil {
		c := ToModifierGroupResponse(mo.ModifierGroup)
		mg = &c
	}

	resp := ModifierOptionResponse{
		ID:              mo.ID,
		ModifierGroupID: mo.ModifierGroupID,
		ModifierGroup:   mg,
		Name:            mo.Name,
		PriceAdjustment: mo.PriceAdjustment,
		CogsAdjustment:  mo.CogsAdjustment,
		CreatedAt:       mo.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt:       mo.UpdatedAt.Format("2006-01-02 15:04:05"),
	}

	if mo.DeletedAt.Valid {
		at := mo.DeletedAt.Time.Format("2006-01-02 15:04:05")
		resp.DeletedAt = &at
	}

	return resp
}

func ToModifierOptionResponsePagination(mo []ModifierOptionResponse, f filter.PaginationWithInputFilter, totalRows int64) ModifierOptionResponsePagination {
	return ModifierOptionResponsePagination{
		Data: mo,
		Meta: f.ToMeta(totalRows),
	}
}

func ToModifierOptionModel(req *CreateModifierOptionRequest) domain.ModifierOption {
	return domain.ModifierOption{
		ModifierGroupID: req.ModifierGroupID,
		Name:            req.Name,
		PriceAdjustment: req.PriceAdjustment,
		CogsAdjustment:  req.CogsAdjustment,
	}
}

func ToUpdateModifierOptionModel(req *UpdateModifierOptionRequest) domain.ModifierOption {
	return domain.ModifierOption{
		ModifierGroupID: req.ModifierGroupID,
		Name:            req.Name,
		PriceAdjustment: req.PriceAdjustment,
		CogsAdjustment:  req.CogsAdjustment,
	}
}
