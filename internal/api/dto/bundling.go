package dto

import (
	"time"

	"github.com/BramAristyo/go-pos-mawish/internal/domain"
	"github.com/BramAristyo/go-pos-mawish/pkg/filter"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type BundlingItemResponse struct {
	ID                uuid.UUID       `json:"id"`
	BundlingPackageID uuid.UUID       `json:"bundlingPackageId"`
	ProductID         uuid.UUID       `json:"productId"`
	Product           ProductResponse `json:"product"`
	Qty               int             `json:"qty"`
	CreatedAt         time.Time       `json:"createdAt"`
}

type BundlingPackageResponse struct {
	ID            uuid.UUID              `json:"id"`
	Name          string                 `json:"name"`
	Description   *string                `json:"description,omitempty"`
	Price         decimal.Decimal        `json:"price"`
	Cogs          decimal.Decimal        `json:"cogs"`
	ImageURL      *string                `json:"imageUrl"`
	UpdatedAt     time.Time              `json:"updatedAt"`
	CreatedAt     time.Time              `json:"createdAt"`
	BundlingItems []BundlingItemResponse `json:"bundlingItems,omitempty"`
}

type BundlingPackagePaginationResponse struct {
	Data []BundlingPackageResponse `json:"data"`
	Meta filter.Meta               `json:"meta"`
}

type BundlingItemRequest struct {
	BundlingPackageID *uuid.UUID `json:"bundlingPackageId"`
	ProductID         uuid.UUID  `json:"productId" binding:"required"`
	Qty               int        `json:"qty" binding:"required,min=1"`
}

type CreateBundlingPackageRequest struct {
	Name          string                `json:"name" binding:"required"`
	Description   *string               `json:"description"`
	Price         decimal.Decimal       `json:"price" binding:"required"`
	Cogs          decimal.Decimal       `json:"cogs" binding:"required"`
	ImageURL      *string               `json:"imageUrl"`
	BundlingItems []BundlingItemRequest `json:"bundlingItems" binding:"required"`
}

type UpdateBundlingPackageRequest struct {
	Name          string                `json:"name" binding:"required"`
	Description   *string               `json:"description"`
	Price         decimal.Decimal       `json:"price" binding:"required"`
	Cogs          decimal.Decimal       `json:"cogs" binding:"required"`
	ImageURL      *string               `json:"imageUrl"`
	IsActive      bool                  `json:"isActive"`
	BundlingItems []BundlingItemRequest `json:"bundlingItems" binding:"required"`
}

func toBundlingItemResponses(bis []domain.BundlingItem) []BundlingItemResponse {
	bisRes := make([]BundlingItemResponse, len(bis))
	for i, bi := range bis {
		bisRes[i] = BundlingItemResponse{
			ID:                bi.ID,
			BundlingPackageID: bi.BundlingPackageID,
			ProductID:         bi.ProductID,
			Product:           ToProductResponse(&bi.Product),
			Qty:               bi.Qty,
		}
	}
	return bisRes
}

func ToBundlingPackageResponse(bp *domain.BundlingPackage) BundlingPackageResponse {
	return BundlingPackageResponse{
		ID:            bp.ID,
		Name:          bp.Name,
		Description:   bp.Description,
		Price:         bp.Price,
		Cogs:          bp.Cogs,
		ImageURL:      bp.ImageURL,
		UpdatedAt:     bp.UpdatedAt,
		CreatedAt:     bp.CreatedAt,
		BundlingItems: toBundlingItemResponses(bp.BundlingItems),
	}
}

func toBundlingPackageResponses(bps []domain.BundlingPackage) []BundlingPackageResponse {
	bpsRes := make([]BundlingPackageResponse, len(bps))
	for i, bp := range bps {
		bpsRes[i] = ToBundlingPackageResponse(&bp)
	}

	return bpsRes
}

func ToBundlingPackagePaginationResponse(bps []domain.BundlingPackage, f filter.PaginationWithInputFilter, totalRows int64) BundlingPackagePaginationResponse {
	return BundlingPackagePaginationResponse{
		Data: toBundlingPackageResponses(bps),
		Meta: f.ToMeta(totalRows),
	}
}

func toBundlingItemModels(req []BundlingItemRequest) []domain.BundlingItem {
	bis := make([]domain.BundlingItem, len(req))
	for i, bi := range req {
		bis[i] = domain.BundlingItem{
			ProductID: bi.ProductID,
			Qty:       bi.Qty,
		}
	}
	return bis
}

func ToBundlingPackageModel(req *CreateBundlingPackageRequest) domain.BundlingPackage {
	return domain.BundlingPackage{
		Name:          req.Name,
		Description:   req.Description,
		Price:         req.Price,
		Cogs:          req.Cogs,
		ImageURL:      req.ImageURL,
		BundlingItems: toBundlingItemModels(req.BundlingItems),
	}
}

func ToUpdateBundlingPackageModel(req *UpdateBundlingPackageRequest) domain.BundlingPackage {
	return domain.BundlingPackage{
		Name:          req.Name,
		Description:   req.Description,
		Price:         req.Price,
		Cogs:          req.Cogs,
		ImageURL:      req.ImageURL,
		IsActive:      req.IsActive,
		BundlingItems: toBundlingItemModels(req.BundlingItems),
	}
}
