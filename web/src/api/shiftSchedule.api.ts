import http from '@/lib/http'
import type { BaseFilterRequest, BaseResponse } from '@/types/common.types'
import type {
  ShiftSchedule,
  CreateShiftScheduleRequest,
  UpdateShiftScheduleRequest,
} from '@/types/shiftSchedule.types'

export const shiftScheduleApi = {
  paginate: (params: BaseFilterRequest) =>
    http.get<any, BaseResponse<ShiftSchedule[]>>('/shift-schedules', { params }),

  getAll: () => http.get<any, BaseResponse<ShiftSchedule[]>>('/shift-schedules/all'),

  getById: (id: string) => http.get<any, BaseResponse<ShiftSchedule>>(`/shift-schedules/${id}`),

  create: (payload: CreateShiftScheduleRequest) =>
    http.post<any, BaseResponse<ShiftSchedule>>('/shift-schedules', payload),

  update: (id: string, payload: UpdateShiftScheduleRequest) =>
    http.put<any, BaseResponse<ShiftSchedule>>(`/shift-schedules/${id}`, payload),

  delete: (id: string) => http.delete<any, BaseResponse<null>>(`/shift-schedules/${id}`),

  restore: (id: string) =>
    http.patch<any, BaseResponse<ShiftSchedule>>(`/shift-schedules/${id}/restore`),
}
