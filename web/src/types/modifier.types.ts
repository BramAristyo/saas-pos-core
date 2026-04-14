export interface ModifierOption {
  id: string
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
}
