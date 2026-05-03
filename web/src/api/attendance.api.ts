import http from '@/lib/http'
import type { BaseFilterRequest, BaseResponse } from '@/types/common.types'
import type { Attendance, CreateAttendanceRequest } from '@/types/attendance.types'

export const attendanceApi = {
  paginate: (params: BaseFilterRequest) =>
    http.get<any, BaseResponse<Attendance[]>>('/attendances', { params }),

  create: (payload: CreateAttendanceRequest) =>
    http.post<any, BaseResponse<Attendance>>('/attendances', payload),
}
