package usecase

import (
	"context"

	"github.com/BramAristyo/go-pos-mawish/internal/api/dto"
	"github.com/BramAristyo/go-pos-mawish/internal/repository"
	"github.com/BramAristyo/go-pos-mawish/pkg/filter"
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

func (u *ReportUseCase) SalesSummary(ctx context.Context, req filter.DynamicFilter) (dto.SalesReportResponse, error) {
	summary, err := u.OrderRepo.SalesSummary(ctx, req)
	if err != nil {
		return dto.SalesReportResponse{}, err
	}

	return dto.ToSalesReportResponse(summary), nil
}

func (u *ReportUseCase) GrossProfit(ctx context.Context, req filter.DynamicFilter) (dto.GrossProfitReportResponse, error) {
	grossProfit, err := u.OrderRepo.GrossProfit(ctx, req)

	grossProfit.CalculateGrossProfit()

	if err != nil {
		return dto.GrossProfitReportResponse{}, err
	}
	return dto.ToGrossProfitReportResponse(grossProfit), nil
}

func (u *ReportUseCase) Transactions(ctx context.Context, req filter.PaginationWithInputFilter) (dto.TransactionReportResponsePagination, error) {
	totalRows, ts, err := u.OrderRepo.TransactionReport(ctx, req)
	if err != nil {
		return dto.TransactionReportResponsePagination{}, err
	}

	transactionResponses := make([]dto.TransactionReportResponse, 0, req.PaginationInput.PageSize)
	for i, t := range ts {
		transactionResponses[i] = dto.TransactionReportResponse{
			OrderNumber: t.OrderNumber,
			Time:        t.Time,
			Product:     t.Product,
			Price:       t.Price,
		}
	}

	return dto.TransactionReportResponsePagination{Data: transactionResponses, Meta: req.ToMeta(totalRows)}, nil
}

// func (u *ReportUseCase) DiscountUsage(ctx context.Context, req filter.PaginationWithInputFilter) (dto.DiscountReportResponsePagination, error)
// func (u *ReportUseCase) ShiftReconciliation(ctx context.Context, req filter.PaginationWithInputFilter) (dto.ShiftReconciliationtResponsePagination, error)
