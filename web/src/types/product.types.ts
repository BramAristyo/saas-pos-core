import type { Category } from './category.types'

export interface Product {
  id: string
  name: string
  description: string
  price: string
  cogs: string
  created_at: string
  category: Category
  deleted_at?: string
}

export interface CreateProductRequest {
  name: string
  description?: string
  categoryId: string
  price: string
  cogs: string
  imageUrl?: string
}

export interface UpdateProductRequest {
  name: string
  description?: string
  categoryId: string
  price: string
  cogs: string
  imageUrl?: string
}

export interface ProductResponse {
  id: string
  name: string
  description: string
  price: string
  cogs: string
  createdAt: string
  updatedAt: string
  deletedAt?: string
  category: Category
}
