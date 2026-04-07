import http from '@/lib/http'
import type { LoginRequest, LoginResponse } from '@/types/auth.types'
import type { BaseResponse } from '@/types/common.types'
import type { User } from '@/types/user.types'

export const authApi = {
  login: (payload: LoginRequest) => http.post<any, BaseResponse<LoginResponse>>('/', payload),
  me: () => http.get<any, BaseResponse<User>>('/me'),
}
