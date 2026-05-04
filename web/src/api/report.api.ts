import http from '@/lib/http'
import type { BaseResponse } from '@/types/common.types'
import type { SalesSummary, SalesSummaryFilter } from '@/types/report.types'

export const reportApi = {
  getSalesSummary: (params: SalesSummaryFilter) =>
    http.get<any, BaseResponse<SalesSummary>>('/dashboard/sales-summary', { params }),
}
