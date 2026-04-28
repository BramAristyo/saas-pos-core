export interface SalesSummary {
  grossSales: string
  discounts: string
  netSales: string
  gratuity: string
  tax: string
  total: string
}

export interface grossProfit {
  grossSales: string
  discounts: string
  netSales: string
  cogs: string
  grossProfit: string
}

export interface discountReport {
  name: string
  count: number
  grossDiscount: string
  discount: string
}

export interface shiftRecociliation {
  cashierName: string
  startTime: string
  endTime: string | null
  totalExpected: string
  totalActual: string
  difference: string
}
