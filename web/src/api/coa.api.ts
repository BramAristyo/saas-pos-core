import type { Coa, CreateCoaRequest, UpdateCoaRequest } from '@/types/coa.types'
import type { BaseFilterRequest, BaseResponse } from '@/types/common.types'
import http from '@/lib/http'

export const coaApi = {
  paginate: (params: BaseFilterRequest) =>
    http.get<any, BaseResponse<Coa[]>>('/coa', { params }),

  getAll: () => http.get<any, BaseResponse<Coa[]>>('/coa/get-all'),

  getById: (id: string) => http.get<any, BaseResponse<Coa>>(`/coa/${id}`),

  create: (payload: CreateCoaRequest) =>
    http.post<any, BaseResponse<Coa>>('/coa', payload),

  update: (id: string, payload: UpdateCoaRequest) =>
    http.put<any, BaseResponse<Coa>>(`/coa/${id}`, payload),

  delete: (id: string) => http.delete<any, BaseResponse<null>>(`/coa/${id}`),

  restore: (id: string) => http.patch<any, BaseResponse<Coa>>(`/coa/${id}/restore`),
}
