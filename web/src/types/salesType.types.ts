import type { AdjustmentType } from './common.types'

export interface SalesType {
  id: string
  name: string
  createdAt: string
}

export interface AdditionalCharge {
  id: string
  salesTypeId: string
  name: string
  type: AdjustmentType
  amount: string
  createdAt: string
}

export interface SalesTypeDetail extends SalesType {
  charges: AdditionalCharge[]
}

export interface AdditionalChargeRequest {
  name: string
  type: AdjustmentType
  amount: string
}

export interface UpdateAdditionalChargeRequest extends AdditionalChargeRequest {
  id?: string
}

export interface CreateSalesTypeRequest {
  name: string
  charges: AdditionalChargeRequest[]
}

export interface UpdateSalesTypeRequest {
  name: string
  charges: UpdateAdditionalChargeRequest[]
}
