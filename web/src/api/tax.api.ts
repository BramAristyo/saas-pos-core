import http from '@/lib/http'
import type { BaseFilterRequest, BaseResponse } from '@/types/common.types'
import type { CreateTaxRequest, Tax, UpdateTaxRequest } from '@/types/tax.types'

export const taxApi = {
  paginate: (params: BaseFilterRequest) => http.get<any, BaseResponse<Tax[]>>('/taxes', { params }),

  getAll: () => http.get<any, BaseResponse<Tax[]>>('/taxes/get-all'),

  getById: (id: string) => http.get<any, BaseResponse<Tax>>(`/taxes/${id}`),

  create: (payload: CreateTaxRequest) => http.post<any, BaseResponse<Tax>>('/taxes', payload),

  update: (id: string, payload: UpdateTaxRequest) =>
    http.put<any, BaseResponse<Tax>>(`/taxes/${id}`, payload),

  delete: (id: string) => http.delete<any, BaseResponse<null>>(`/taxes/${id}`),
}
