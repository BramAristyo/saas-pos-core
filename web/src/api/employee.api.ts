import http from '@/lib/http'
import type { BaseFilterRequest, BaseResponse } from '@/types/common.types'
import type {
  Employee,
  CreateEmployeeRequest,
  UpdateEmployeeRequest,
} from '@/types/employee.types'

export const employeeApi = {
  paginate: (params: BaseFilterRequest) =>
    http.get<any, BaseResponse<Employee[]>>('/employees', { params }),

  getAll: () => http.get<any, BaseResponse<Employee[]>>('/employees/all'),

  getById: (id: string) => http.get<any, BaseResponse<Employee>>(`/employees/${id}`),

  create: (payload: CreateEmployeeRequest) =>
    http.post<any, BaseResponse<Employee>>('/employees', payload),

  update: (id: string, payload: UpdateEmployeeRequest) =>
    http.put<any, BaseResponse<Employee>>(`/employees/${id}`, payload),

  delete: (id: string) => http.delete<any, BaseResponse<null>>(`/employees/${id}`),

  restore: (id: string) => http.patch<any, BaseResponse<Employee>>(`/employees/${id}/restore`),
}
