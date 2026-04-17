import http from '@/lib/http'
import type { BaseFilterRequest, BaseResponse } from '@/types/common.types'
import type {
  CreateModifierGroupRequest,
  ModifierGroup,
  ModifierGroupDetail,
  UpdateModifierGroupRequest,
} from '@/types/modifier.types'

export const modifierApi = {
  paginate: (params: BaseFilterRequest) =>
    http.get<any, BaseResponse<ModifierGroup[]>>('/modifier-groups', { params }),

  getAll: () => http.get<any, BaseResponse<ModifierGroup[]>>('/modifier-groups/get-all'),

  getById: (id: string) =>
    http.get<any, BaseResponse<ModifierGroupDetail>>(`/modifier-groups/${id}`),

  create: (payload: CreateModifierGroupRequest) =>
    http.post<any, BaseResponse<ModifierGroup>>('/modifier-groups', payload),

  update: (id: string, payload: UpdateModifierGroupRequest) =>
    http.put<any, BaseResponse<ModifierGroup>>(`/modifier-groups/${id}`, payload),

  delete: (id: string) => http.delete<any, BaseResponse<null>>(`/modifier-groups/${id}`),

  restore: (id: string) =>
    http.patch<any, BaseResponse<ModifierGroup>>(`/modifier-groups/${id}/restore`),
}
