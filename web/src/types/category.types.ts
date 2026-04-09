export interface Category {
  id: string
  name: string
  description: string
  updatedAt: string
  createdAt: string
}

export type CreateCategoryRequest = Omit<Category, 'id' | 'updatedAt' | 'createdAt'>
export type UpdateCategoryRequest = CreateCategoryRequest
