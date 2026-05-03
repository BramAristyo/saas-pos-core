import { coaApi } from '@/api/coa.api'
import type { Coa, CreateCoaRequest, UpdateCoaRequest } from '@/types/coa.types'
import type { BaseFilterRequest, Meta } from '@/types/common.types'
import { defineStore } from 'pinia'
import { ref } from 'vue'

export const useCoaStore = defineStore('coa', () => {
  const coas = ref<Coa[]>([])
  const meta = ref<Meta | null>(null)
  const selected = ref<Coa | null>(null)
  const loading = ref(false)
  const error = ref<string | null>(null)

  async function fetchAll() {
    loading.value = true
    error.value = null
    try {
      const res = await coaApi.getAll()
      coas.value = res.data || []
    } catch (err: any) {
      error.value = err?.message || 'Unsuccessfully get COAs'
      coas.value = []
    } finally {
      loading.value = false
    }
  }

  async function fetchPaginated(params: BaseFilterRequest) {
    loading.value = true
    error.value = null
    try {
      const res = await coaApi.paginate(params)
      coas.value = res.data || []
      meta.value = res.meta || null
    } catch (err: any) {
      error.value = err?.message || 'Unsuccessfully get COAs'
      coas.value = []
    } finally {
      loading.value = false
    }
  }

  async function fetchById(id: string) {
    loading.value = true
    error.value = null
    try {
      const res = await coaApi.getById(id)
      selected.value = res.data
    } catch (err: any) {
      error.value = err?.message || 'Unsuccessfully get COA'
    } finally {
      loading.value = false
    }
  }

  async function create(payload: CreateCoaRequest) {
    loading.value = true
    error.value = null
    try {
      const res = await coaApi.create(payload)
      coas.value = [res.data, ...coas.value]
    } catch (err: any) {
      error.value = err?.message || 'Unsuccessfully create COA'
      throw err
    } finally {
      loading.value = false
    }
  }

  async function update(id: string, payload: UpdateCoaRequest) {
    loading.value = true
    error.value = null
    try {
      const res = await coaApi.update(id, payload)
      const index = coas.value.findIndex((c) => c.id === id)
      if (index !== -1) coas.value[index] = res.data
    } catch (err: any) {
      error.value = err?.message || 'Unsuccessfully update COA'
      throw err
    } finally {
      loading.value = false
    }
  }

  async function remove(id: string) {
    loading.value = true
    error.value = null
    try {
      await coaApi.delete(id)
      coas.value = coas.value.filter((c) => c.id !== id)
    } catch (err: any) {
      error.value = err?.message || 'Unsuccessfully delete COA'
      throw err
    } finally {
      loading.value = false
    }
  }

  async function restore(id: string) {
    loading.value = true
    error.value = null
    try {
      const res = await coaApi.restore(id)
      const index = coas.value.findIndex((c) => c.id === id)
      if (index !== -1) coas.value[index] = res.data
    } catch (err: any) {
      error.value = err?.message || 'Unsuccessfully restore COA'
      throw err
    } finally {
      loading.value = false
    }
  }

  async function ensureDataLoaded() {
    if (!coas.value || coas.value.length === 0) {
      await fetchAll()
    }
  }

  return {
    coas,
    meta,
    selected,
    loading,
    error,
    fetchAll,
    fetchPaginated,
    fetchById,
    create,
    update,
    remove,
    restore,
    ensureDataLoaded,
  }
})
