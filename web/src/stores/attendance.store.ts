import { attendanceApi } from '@/api/attendance.api'
import type { BaseFilterRequest, Meta } from '@/types/common.types'
import type { Attendance, CreateAttendanceRequest } from '@/types/attendance.types'
import { defineStore } from 'pinia'
import { ref } from 'vue'

export const useAttendanceStore = defineStore('attendance', () => {
  const attendances = ref<Attendance[]>([])
  const meta = ref<Meta | null>(null)
  const loading = ref(false)
  const error = ref<string | null>(null)

  async function fetchAttendances(params: BaseFilterRequest) {
    loading.value = true
    error.value = null
    try {
      const res = await attendanceApi.paginate(params)
      attendances.value = res.data
      meta.value = res.meta || null
    } catch (err: unknown) {
      const errorResponse = err as { message?: string }
      error.value = errorResponse.message || 'Unsuccessfully get attendances'
    } finally {
      loading.value = false
    }
  }

  async function createAttendance(payload: CreateAttendanceRequest) {
    loading.value = true
    error.value = null
    try {
      const res = await attendanceApi.create(payload)
      attendances.value = [res.data, ...attendances.value]
    } catch (err: unknown) {
      const errorResponse = err as { message?: string }
      error.value = errorResponse.message || 'Unsuccessfully create attendance'
      throw err
    } finally {
      loading.value = false
    }
  }

  return {
    attendances,
    meta,
    loading,
    error,
    fetchAttendances,
    create: createAttendance,
  }
})
