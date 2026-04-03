package domain

import "github.com/shopspring/decimal"

type SalesSummaryDashboard struct {
	GrossSales       decimal.Decimal
	NetSales         decimal.Decimal
	GrossProfit      decimal.Decimal
	TransactionCount int64
	AverageSales     decimal.Decimal
	GrossMargin      float32
}

func (s *SalesSummaryDashboard) Calculate() {
	if s.TransactionCount > 0 {
		s.AverageSales = s.NetSales.Div(decimal.NewFromInt(s.TransactionCount))
	}
	if !s.NetSales.IsZero() {
		margin, _ := s.GrossProfit.Div(s.NetSales).Float64()
		s.GrossMargin = float32(margin * 100)
	}
}
