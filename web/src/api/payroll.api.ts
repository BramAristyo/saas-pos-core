import type { Payroll, CreatePayrollRequest } from '@/types/payroll.types'
import type { BaseFilterRequest, BaseResponse } from '@/types/common.types'
import http from '@/lib/http'

export const payrollApi = {
  paginate: (params: BaseFilterRequest) =>
    http.get<any, BaseResponse<Payroll[]>>('/payrolls', { params }),

  create: (payload: CreatePayrollRequest) =>
    http.post<any, BaseResponse<Payroll>>('/payrolls', payload),
}
