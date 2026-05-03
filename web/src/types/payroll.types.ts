export interface Payroll {
  id: string
  employeeName: string
  employeeCode: string
  periodStart: string
  periodEnd: string
  baseSalary: string | number
  totalDeduction: string | number
  netSalary: string | number
  notes: string | null
  createdAt: string
}

export interface CreatePayrollRequest {
  employeeID: string
  periodStart: string
  periodEnd: string
}
