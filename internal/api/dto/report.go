package dto

import (
	"github.com/BramAristyo/go-pos-mawish/internal/domain"
	"github.com/BramAristyo/go-pos-mawish/pkg/filter"
	"github.com/shopspring/decimal"
)

type SalesReportResponse struct {
	GrossSales decimal.Decimal `json:"grossSales"`
	Discounts  decimal.Decimal `json:"discounts"`
	NetSales   decimal.Decimal `json:"netSales"`
	Gratuity   decimal.Decimal `json:"gratuity"`
	Tax        decimal.Decimal `json:"tax"`
	Total      decimal.Decimal `json:"total"`
}

func ToSalesReportResponse(s domain.SalesSummary) SalesReportResponse {
	return SalesReportResponse{
		GrossSales: s.GrossSales,
		Discounts:  s.Discounts,
		NetSales:   s.NetSales,
		Gratuity:   s.Gratuity,
		Tax:        s.Tax,
		Total:      s.Total,
	}
}

type GrossProfitReportResponse struct {
	GrossSales  decimal.Decimal `json:"grossSales"`
	Discounts   decimal.Decimal `json:"discounts"`
	NetSales    decimal.Decimal `json:"netSales"`
	Cogs        decimal.Decimal `json:"cogs"`
	GrossProfit decimal.Decimal `json:"grossProfit"`
}

func ToGrossProfitReportResponse(gp domain.GrossProfit) GrossProfitReportResponse {
	return GrossProfitReportResponse{
		GrossSales:  gp.GrossSales,
		Discounts:   gp.Discounts,
		NetSales:    gp.NetSales,
		Cogs:        gp.Cogs,
		GrossProfit: gp.GrossProfit,
	}
}

type DiscountReportResponse struct {
	Name          string          `json:"name"`
	Count         int64           `json:"count"`
	GrossDiscount decimal.Decimal `json:"grossDiscount"`
	Discount      decimal.Decimal `json:"discount"`
}
type DiscountReportResponsePagination struct {
	Data []DiscountReportResponse `json:"data"`
	Meta filter.Meta              `json:"meta"`
}

type TransactionReportResponse struct {
	OrderNumber string          `json:"orderNumber"`
	Time        string          `json:"time"`
	Product     string          `json:"product"`
	Price       decimal.Decimal `json:"price"`
}
type TransactionReportResponsePagination struct {
	Data []TransactionReportResponse `json:"data"`
	Meta filter.Meta                 `json:"meta"`
}

type ShiftReconciliationResponse struct {
	CashierName   string          `json:"cashierName"`
	StartTime     string          `json:"startTime"`
	EndTime       *string         `json:"endTime,omitempty"`
	TotalExpected decimal.Decimal `json:"totalExpected"`
	TotalActual   decimal.Decimal `json:"totalActual"`
	Difference    decimal.Decimal `json:"difference"`
}
type ShiftReconciliationtResponsePagination struct {
	Data []ShiftReconciliationResponse `json:"data"`
	Meta filter.Meta                   `json:"meta"`
}
