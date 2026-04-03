package usecase

import (
	"context"

	"github.com/BramAristyo/go-pos-mawish/internal/api/dto"
	"github.com/BramAristyo/go-pos-mawish/internal/repository"
	"github.com/BramAristyo/go-pos-mawish/pkg/filter"
)

type DashboardUseCase struct {
	OrderRepo *repository.OrderRepository
}

func NewDashboardUseCase(orderRepo *repository.OrderRepository) *DashboardUseCase {
	return &DashboardUseCase{
		OrderRepo: orderRepo,
	}
}

func (u *DashboardUseCase) SalesSummary(ctx context.Context, req filter.DynamicFilter) (dto.SalesSummaryDashboardResponse, error) {
	summary, err := u.OrderRepo.SalesSummaryDashboard(ctx, req)
	if err != nil {
		return dto.SalesSummaryDashboardResponse{}, err
	}

	return dto.ToSalesSummaryDashboardResponse(summary), nil
}
