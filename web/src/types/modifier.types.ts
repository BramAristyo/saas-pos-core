import type { Product } from './product.types'

export interface ModifierOption {
  id?: string | null
  name: string
  priceAdjustment: string
  cogsAdjustment: string
  created_at: string
}

export interface ModifierGroup {
  id: string
  name: string
  isRequired: boolean
  created_at: string
  updated_at: string
}

export interface ModifierGroupDetail {
  id: string
  name: string
  isRequired: boolean
  created_at: string
  updated_at: string
  options: ModifierGroup[]
  productModifiers: Product[]
}

export interface CreateModifierOptionRequest {
  name: string
  priceAdjustment: string
  cogsAdjustment: string
}

export interface UpdateModifierOptionRequest extends CreateModifierOptionRequest {
  id?: string
}

export interface CreateModifierGroupRequest {
  name: string
  isRequired: boolean
  options: CreateModifierOptionRequest[]
  productModifiers: string[] | null
}

export interface UpdateModifierGroupRequest {
  id: string
  name: string
  isRequired: boolean
  options: UpdateModifierOptionRequest[]
  productModifiers: string[] | null
}
