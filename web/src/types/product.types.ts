import type { BaseResponse } from './common.types'
import type { Category } from './category.types'

export interface Product {
  id: string
  name: string
  description?: string
  price: number
  cogs: number
  categoryId: string
  category: Category
  imageUrl?: string
  createdAt: string
  updatedAt: string
  deletedAt?: string
}

export interface StoreProductRequest {
  name: string
  description?: string
  price: number
  cogs: number
  categoryId: string
  imageUrl?: string
}

export interface UpdateProductRequest {
  name: string
  description?: string
  price: number
  cogs: number
  categoryId: string
  imageUrl?: string
}

export type PaginatedProductResponse = BaseResponse<Product[]>
