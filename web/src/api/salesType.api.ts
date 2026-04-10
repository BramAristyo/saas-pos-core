import http from '@/lib/http'
import type { BaseFilterRequest, BaseResponse } from '@/types/common.types'
import type {
  CreateSalesTypeRequest,
  SalesType,
  SalesTypeDetail,
  UpdateSalesTypeRequest,
} from '@/types/salesType.types'

export const salesTypeApi = {
  paginate: (params: BaseFilterRequest) =>
    http.get<any, BaseResponse<SalesType[]>>('/sales-types', { params }),

  getAll: () => http.get<any, BaseResponse<SalesType[]>>('/sales-types/get-all'),

  getById: (id: string) => http.get<any, BaseResponse<SalesTypeDetail>>(`/sales-types/${id}`),

  create: (payload: CreateSalesTypeRequest) =>
    http.post<any, BaseResponse<SalesType>>('/sales-types', payload),

  update: (id: string, payload: UpdateSalesTypeRequest) =>
    http.put<any, BaseResponse<SalesType>>(`/sales-types/${id}`, payload),

  delete: (id: string) => http.delete<any, BaseResponse<null>>(`/sales-types/${id}`),

  restore: (id: string) => http.patch<any, BaseResponse<null>>(`/sales-types/${id}/restore`),
}
