import http from '@/lib/http'
import type { BaseFilterRequest, BaseResponse } from '@/types/common.types'

export function createCrudApi<
  TModel,
  TCreate = Partial<TModel>,
  TUpdate = Partial<TModel>,
  TFilter = BaseFilterRequest,
>(
  endpoint: string,
  options?: {
    hasGetAll?: boolean
  },
) {
  const api = {
    paginate: (params: TFilter) =>
      http.get<any, BaseResponse<TModel[]>>(`/${endpoint}`, { params }),

    getById: (id: string) => http.get<any, BaseResponse<TModel>>(`/${endpoint}/${id}`),

    create: (payload: TCreate) => http.post<any, BaseResponse<TModel>>(`/${endpoint}`, payload),

    update: (id: string, payload: TUpdate) =>
      http.put<any, BaseResponse<TModel>>(`/${endpoint}/${id}`, payload),

    delete: (id: string) => http.delete<any, BaseResponse<null>>(`/${endpoint}/${id}`),
  }

  if (options?.hasGetAll) {
    return {
      ...api,
      getAll: () => http.get<any, BaseResponse<TModel[]>>(`/${endpoint}/get-all`),
    }
  }

  return api
}

// export const categoryApi = createCrudApi<
//   Category,
//   CreateCategoryRequest,
//   UpdateCategoryRequest
// >('categories')

// export const discountExtendedApi = {
//   ...discountApi,
//   applyCoupon: (code: string) => http.post(`/discounts/apply`, { code })
// }
