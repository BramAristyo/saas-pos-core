import type { Category } from './category.types'

export interface Product {
  id: string
  name: string
  description: string
  price: string
  cogs: string
  created_at: string
  category: Category
}
