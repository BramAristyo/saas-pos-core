import type { User } from './user'

export interface LoginRequest {
  email: string
  password: string
}

export interface RegisterRequest {
  name: string
  email: string
  password: string
  passwordConfirmation: string
}

export interface LoginResponse {
  user: User
  token: string
}
