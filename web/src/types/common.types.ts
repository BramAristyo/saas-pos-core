export type AdjustmentType = 'fixed' | 'percentage'

export interface BaseResponse<T> {
  success: boolean
  message?: string
  data: T
  error?: string | ValidationError[]
  meta?: Meta
}

export interface Meta {
  page: number
  pageSize: number
  totalRows: number
  totalPages: number
  hasNext: boolean
  hasPrev: boolean
}

export interface ValidationError {
  property: string
  tag: string
  value: string
  message?: string
}

export interface PaginationInput {
  pageSize: number
  pageNumber: number
}

export interface BaseFilterRequest extends PaginationInput {
  search?: string
}
