export interface SalesSummary {
  grossSales: number
  netSales: number
  grossProfit: number
  transactionCount: number
  averageSales: number
  grossMargin: number
}

export interface SalesSummaryFilter {
  'filter[created_at][type]': 'inRange'
  'filter[created_at][filterType]': 'date'
  'filter[created_at][from]': string
  'filter[created_at][to]': string
}
