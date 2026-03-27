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
	GrossSales  decimal.Decimal
	Discounts   decimal.Decimal
	NetSales    decimal.Decimal
	Cogs        decimal.Decimal
	GrossProfit decimal.Decimal
}

func (gp *GrossProfit) CalculateGrossProfit() {
	gp.GrossProfit = gp.NetSales.Sub(gp.Cogs)
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
