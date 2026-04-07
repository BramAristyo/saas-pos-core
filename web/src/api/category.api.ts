import http from '@/lib/http'
import type { Category, CreateCategoryRequest, UpdateCategoryRequest } from '@/types/category.types'
import type { BaseResponse } from '@/types/common.types'

export const categoryApi = {
  getAll: () => http.get<any, BaseResponse<Category[]>>('/categories/get-all'),
  getById: (id: string) => http.get<any, BaseResponse<Category>>(`/categories/${id}`),
  create: (payload: CreateCategoryRequest) =>
    http.post<any, BaseResponse<Category>>('/categories', payload),
  update: (id: string, payload: UpdateCategoryRequest) =>
    http.put<any, BaseResponse<Category>>(`/categories/${id}`, payload),
  delete: (id: string) => http.delete<any, BaseResponse<null>>(`categories/${id}`),
}
