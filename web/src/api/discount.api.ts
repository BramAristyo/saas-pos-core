import type { CreateDiscountRequest, Discount, UpdateDiscountRequest } from '@/types/discount.types'
import { createCrudApi } from './base.api'

export const discountApi = createCrudApi<Discount, CreateDiscountRequest, UpdateDiscountRequest>(
  'discounts',
  {
    hasGetAll: false,
  },
)
