export interface Employee {
  id: string
  code: string
  name: string
  phone: string
  baseSalary: number
  hasChangedPin: boolean
  createdAt: string
  updatedAt: string
}

export interface CreateEmployeeRequest {
  name: string
  phone: string
  baseSalary: number
  pin: string
}

export interface UpdateEmployeeRequest {
  name: string
  phone: string
  baseSalary: number
  pin?: string
}
