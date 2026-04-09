import type { CreateDiscountRequest, Discount, UpdateDiscountRequest } from '@/types/discount.types'
import type { BaseFilterRequest, BaseResponse } from '@/types/common.types'
import http from '@/lib/http'

export const discountApi = {
  paginate: (params: BaseFilterRequest) =>
    http.get<any, BaseResponse<Discount[]>>('/discounts', { params }),

  getById: (id: string) => http.get<any, BaseResponse<Discount>>(`/discounts/${id}`),

  create: (payload: CreateDiscountRequest) =>
    http.post<any, BaseResponse<Discount>>('/discounts', payload),

  update: (id: string, payload: UpdateDiscountRequest) =>
    http.put<any, BaseResponse<Discount>>(`/discounts/${id}`, payload),

  delete: (id: string) => http.delete<any, BaseResponse<null>>(`/discounts/${id}`),
}
