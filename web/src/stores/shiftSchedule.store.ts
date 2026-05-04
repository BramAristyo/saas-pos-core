import { shiftScheduleApi } from '@/api/shiftSchedule.api'
import type { BaseFilterRequest, Meta } from '@/types/common.types'
import type {
  CreateShiftScheduleRequest,
  ShiftSchedule,
  UpdateShiftScheduleRequest,
} from '@/types/shiftSchedule.types'
import { defineStore } from 'pinia'
import { ref } from 'vue'

export const useShiftScheduleStore = defineStore('shiftSchedule', () => {
  const shiftSchedules = ref<ShiftSchedule[]>([])
  const meta = ref<Meta | null>(null)
  const selected = ref<ShiftSchedule | null>(null)
  const loading = ref(false)
  const error = ref<string | null>(null)

  async function fetchShiftSchedules(params: BaseFilterRequest) {
    loading.value = true
    error.value = null
    try {
      const res = await shiftScheduleApi.paginate(params)
      shiftSchedules.value = res.data
      meta.value = res.meta || null
    } catch (err: unknown) {
      const errorResponse = err as { message?: string }
      error.value = errorResponse.message || 'Unsuccessfully get shift schedules'
    } finally {
      loading.value = false
    }
  }

  async function fetchAllShiftSchedules() {
    loading.value = true
    error.value = null
    try {
      const res = await shiftScheduleApi.getAll()
      shiftSchedules.value = res.data
    } catch (err: unknown) {
      const errorResponse = err as { message?: string }
      error.value = errorResponse.message || 'Unsuccessfully get shift schedules'
    } finally {
      loading.value = false
    }
  }

  async function fetchShiftScheduleById(id: string) {
    loading.value = true
    error.value = null
    try {
      const res = await shiftScheduleApi.getById(id)
      selected.value = res.data
    } catch (err: unknown) {
      const errorResponse = err as { message?: string }
      error.value = errorResponse.message || 'Unsuccessfully get shift schedule'
    } finally {
      loading.value = false
    }
  }

  async function createShiftSchedule(payload: CreateShiftScheduleRequest) {
    loading.value = true
    error.value = null
    try {
      const res = await shiftScheduleApi.create(payload)
      shiftSchedules.value = [res.data, ...shiftSchedules.value]
    } catch (err: unknown) {
      const errorResponse = err as { message?: string }
      error.value = errorResponse.message || 'Unsuccessfully create shift schedule'
      throw err
    } finally {
      loading.value = false
    }
  }

  async function updateShiftSchedule(id: string, payload: UpdateShiftScheduleRequest) {
    loading.value = true
    error.value = null
    try {
      const res = await shiftScheduleApi.update(id, payload)
      const index = shiftSchedules.value.findIndex((s) => s.id === id)
      if (index !== -1) shiftSchedules.value[index] = res.data
      if (selected.value?.id === id) selected.value = res.data
    } catch (err: unknown) {
      const errorResponse = err as { message?: string }
      error.value = errorResponse.message || 'Unsuccessfully update shift schedule'
      throw err
    } finally {
      loading.value = false
    }
  }

  async function deleteShiftSchedule(id: string) {
    loading.value = true
    error.value = null
    try {
      await shiftScheduleApi.delete(id)
      shiftSchedules.value = shiftSchedules.value.filter((s) => s.id !== id)
    } catch (err: unknown) {
      const errorResponse = err as { message?: string }
      error.value = errorResponse.message || 'Unsuccessfully delete shift schedule'
      throw err
    } finally {
      loading.value = false
    }
  }

  async function restoreShiftSchedule(id: string) {
    loading.value = true
    error.value = null
    try {
      await shiftScheduleApi.restore(id)
    } catch (err: unknown) {
      const errorResponse = err as { message?: string }
      error.value = errorResponse.message || 'Unsuccessfully restore shift schedule'
      throw err
    } finally {
      loading.value = false
    }
  }

  async function ensureDataLoaded() {
    if (shiftSchedules.value.length === 0) await fetchAllShiftSchedules()
  }

  return {
    shiftSchedules,
    meta,
    selected,
    loading,
    error,
    fetchShiftSchedules,
    fetchAllShiftSchedules,
    fetchShiftScheduleById,
    create: createShiftSchedule,
    update: updateShiftSchedule,
    remove: deleteShiftSchedule,
    restore: restoreShiftSchedule,
    ensureDataLoaded,
  }
})
