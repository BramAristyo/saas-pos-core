export interface Coa {
  id: string
  name: string
  type: string
  isSystem: boolean
  IsOperational: boolean
  createdAt: string
  updatedAt: string
  deletedAt?: string
}

export interface CreateCoaRequest {
  name: string
  type: string
  isOperational: boolean
}

export interface UpdateCoaRequest {
  name: string
  type: string
  isOperational: boolean
}
