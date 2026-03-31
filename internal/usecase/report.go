package usecase

import (
	"github.com/BramAristyo/go-pos-mawish/internal/repository"
)

type ReportUseCase struct {
	OrderRepo *repository.OrderRepository
	ShiftRepo *repository.ShiftRepository
}

func NewReportUseCase(
	orderRepo *repository.OrderRepository,
	shiftRepo *repository.ShiftRepository,
) *ReportUseCase {
	return &ReportUseCase{
		OrderRepo: orderRepo,
		ShiftRepo: shiftRepo,
	}
}

// func (u *ReportUseCase) SalesSummary(ctx context.Context, req filter.DynamicFilter) (dto.SalesReportResponse, error)
// func (u *ReportUseCase) GrossProfit(ctx context.Context, req filter.DynamicFilter) (dto.GrossProfitReportResponse, error)
// func (u *ReportUseCase) DiscountUsage(ctx context.Context, req filter.PaginationWithInputFilter) (dto.DiscountReportResponsePagination, error)
// func (u *ReportUseCase) ShiftReconciliation(ctx context.Context, req filter.PaginationWithInputFilter) (dto.ShiftReconciliationtResponsePagination, error)
