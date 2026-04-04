export interface Category {
  id: string
  name: string
  description: string
  updatedAt: string
  createdAt: string
}

export interface CreateCategoryRequest {
  name: string
  description: string
}

export interface UpdateCategoryRequest {
  name: string
  description: string
}
