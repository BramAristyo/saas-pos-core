import type { AdjustmentType } from './common.types'

export interface Discount {
  id: string
  name: string
  type: AdjustmentType
  value: string
  startDate: string | null
  endDate: string | null
  createdAt: string
}

export type CreateDiscountRequest = Omit<Discount, 'id' | 'createdAt'>
export type UpdateDiscountRequest = CreateDiscountRequest
