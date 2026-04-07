import { authApi } from '@/api/auth.api'
import type { LoginRequest } from '@/types/auth.types'
import type { User } from '@/types/user.types'
import { defineStore } from 'pinia'
import { computed, ref } from 'vue'

export const useAuthStore = defineStore('auth', () => {
  const user = ref<User | null>(null)
  const token = ref<string | null>(localStorage.getItem('token'))
  const loading = ref(false)
  const error = ref<string | null>(null)

  const isAuthenticated = computed(() => !!token.value)
  const isAdmin = computed(() => user.value?.role === 'admin')

  async function login(payload: LoginRequest) {
    loading.value = true
    error.value = null
    try {
      const res = await authApi.login(payload)
      console.log(res)
      token.value = res.data.token
      user.value = res.data.user

      localStorage.setItem('token', res.data.token)

      return res
    } catch (err: any) {
      error.value = err.message || err.error || 'Something went wrong'
      throw err
    } finally {
      loading.value = false
    }
  }

  async function fetchMe() {
    try {
      const res = await authApi.me()
      user.value = res.data
    } catch {}
  }

  function logout() {
    user.value = null
    token.value = null
    localStorage.removeItem('token')
  }

  return {
    user,
    token,
    loading,
    error,

    isAuthenticated,
    isAdmin,

    login,
    fetchMe,
    logout,
  }
})
