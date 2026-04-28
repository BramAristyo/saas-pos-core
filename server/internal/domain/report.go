package domain

import "github.com/shopspring/decimal"

type SalesSummary struct {
	GrossSales decimal.Decimal
	Discounts  decimal.Decimal
	NetSales   decimal.Decimal
	Gratuity   decimal.Decimal
	Tax        decimal.Decimal
	Total      decimal.Decimal
}

type GrossProfit struct {
	GrossSales            decimal.Decimal
	Discounts             decimal.Decimal
	NetSales              decimal.Decimal
	NetSalesPercentage    decimal.Decimal
	Cogs                  decimal.Decimal
	CogsPercentage        decimal.Decimal
	GrossProfit           decimal.Decimal
	GrossProfitPercentage decimal.Decimal
}

func (gp *GrossProfit) CalculateGrossProfit() {
	gp.GrossProfit = gp.NetSales.Sub(gp.Cogs)
}

func (gp *GrossProfit) PercentageCalculation() {
	if gp.GrossSales.IsZero() {
		return
	}

	hundred := decimal.NewFromInt(100)

	gp.NetSalesPercentage = gp.NetSales.Div(gp.GrossSales).Mul(hundred)
	gp.CogsPercentage = gp.Cogs.Div(gp.GrossSales).Mul(hundred)
	gp.GrossProfitPercentage = gp.GrossProfit.Div(gp.GrossSales).Mul(hundred)
}

type DiscountReportWihFooter struct {
	TotalCount         int64
	TotalGrossDiscount decimal.Decimal
	Discounts          []DiscountReport
}

func (dr *DiscountReportWihFooter) CalculateTotal() {
	totalCount := int64(0)
	totalGrossDiscount := decimal.NewFromInt(0)

	for _, disc := range dr.Discounts {
		totalCount = totalCount + disc.Count
		totalGrossDiscount = totalGrossDiscount.Add(disc.GrossDiscount)
	}

	dr.TotalCount = totalCount
	dr.TotalGrossDiscount = totalGrossDiscount
}

type DiscountReport struct {
	Name          string
	Count         int64
	GrossDiscount decimal.Decimal
	Discount      decimal.Decimal
}

type Transaction struct {
	OrderNumber string
	Time        string
	Product     string
	Price       decimal.Decimal
}

type ShiftReconciliaton struct {
	CashierName   string
	StartTime     string
	EndTime       *string
	TotalExpected decimal.Decimal
	TotalActual   decimal.Decimal
	Difference    decimal.Decimal
}

func (sr *ShiftReconciliaton) CalculateDiff() {
	sr.Difference = sr.TotalExpected.Sub(sr.TotalActual)
}
