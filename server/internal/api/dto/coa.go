package dto

import (
	"github.com/BramAristyo/saas-pos-core/server/internal/domain"
	"github.com/BramAristyo/saas-pos-core/server/pkg/filter"
	"github.com/google/uuid"
)

type ChartOfAccountResponse struct {
	ID        uuid.UUID      `json:"id"`
	Name      string         `json:"name"`
	Type      domain.COAType `json:"type"`
	IsSystem  bool           `json:"isSystem"`
	CreatedAt string         `json:"createdAt"`
	DeletedAt *string        `json:"deletedAt,omitempty"`
}

type ChartOfAccountResponsePagination struct {
	Data []ChartOfAccountResponse `json:"data"`
	Meta filter.Meta              `json:"meta"`
}

type CreateCOARequest struct {
	Name string         `json:"name" binding:"required,min=3,max=100"`
	Type domain.COAType `json:"type" binding:"required,oneof=in out"`
}

type UpdateCOARequest struct {
	Name string         `json:"name" binding:"required,min=3,max=100"`
	Type domain.COAType `json:"type" binding:"required,oneof=in out"`
}

func ToCOAModel(req *CreateCOARequest) domain.ChartOfAccount {
	return domain.ChartOfAccount{
		Name: req.Name,
		Type: req.Type,
	}
}

func ToUpdateCOAModel(req *UpdateCOARequest) domain.ChartOfAccount {
	return domain.ChartOfAccount{
		Name: req.Name,
		Type: req.Type,
	}
}

func ToCOAResponse(coa *domain.ChartOfAccount) ChartOfAccountResponse {
	if coa == nil {
		return ChartOfAccountResponse{}
	}
	resp := ChartOfAccountResponse{
		ID:        coa.ID,
		Name:      coa.Name,
		Type:      coa.Type,
		IsSystem:  coa.IsSystem,
		CreatedAt: coa.CreatedAt.Format("2006-01-02 15:04:05"),
	}

	if coa.DeletedAt.Valid {
		t := coa.DeletedAt.Time.Format("2006-01-02 15:04:05")
		resp.DeletedAt = &t
	}

	return resp
}

func ToCOAResponses(coas []domain.ChartOfAccount) []ChartOfAccountResponse {
	res := make([]ChartOfAccountResponse, len(coas))
	for i, coa := range coas {
		res[i] = ToCOAResponse(&coa)
	}
	return res
}

func ToCOAResponsePagination(coas []ChartOfAccountResponse, p filter.PaginationWithInputFilter, totalRows int64) ChartOfAccountResponsePagination {
	return ChartOfAccountResponsePagination{
		Data: coas,
		Meta: p.ToMeta(totalRows),
	}
}
