import { taxApi } from '@/api/tax.api'
import type { Tax, CreateTaxRequest, UpdateTaxRequest } from '@/types/tax.types'
import type { BaseFilterRequest, Meta } from '@/types/common.types'
import { defineStore } from 'pinia'
import { ref } from 'vue'

export const useTaxStore = defineStore('tax', () => {
  const taxes = ref<Tax[]>([])
  const meta = ref<Meta | null>(null)
  const selected = ref<Tax | null>(null)
  const loading = ref(false)
  const error = ref<string | null>(null)

  async function fetchAll() {
    loading.value = true
    error.value = null
    try {
      const res = await taxApi.getAll()
      taxes.value = res.data
    } catch (err: any) {
      error.value = err?.message || 'Unsuccessfully get taxes'
    } finally {
      loading.value = false
    }
  }

  async function fetchPaginated(params: BaseFilterRequest) {
    loading.value = true
    error.value = null
    try {
      const res = await taxApi.paginate(params)
      taxes.value = res.data
      meta.value = res.meta || null
    } catch (err: any) {
      error.value = err?.message || 'Unsuccessfully get taxes'
    } finally {
      loading.value = false
    }
  }

  async function fetchById(id: string) {
    loading.value = true
    error.value = null
    try {
      const res = await taxApi.getById(id)
      selected.value = res.data
    } catch (err: any) {
      error.value = err?.message || 'Unsuccessfully get tax'
    } finally {
      loading.value = false
    }
  }

  async function create(payload: CreateTaxRequest) {
    loading.value = true
    error.value = null
    try {
      const res = await taxApi.create(payload)
      taxes.value = [res.data, ...taxes.value]
    } catch (err: any) {
      error.value = err?.message || 'Unsuccessfully create tax'
      throw err
    } finally {
      loading.value = false
    }
  }

  async function update(id: string, payload: UpdateTaxRequest) {
    loading.value = true
    error.value = null
    try {
      const res = await taxApi.update(id, payload)
      const index = taxes.value.findIndex((t) => t.id === id)
      if (index !== -1) taxes.value[index] = res.data
    } catch (err: any) {
      error.value = err?.message || 'Unsuccessfully update tax'
      throw err
    } finally {
      loading.value = false
    }
  }

  async function remove(id: string) {
    loading.value = true
    error.value = null
    try {
      await taxApi.delete(id)
      taxes.value = taxes.value.filter((t) => t.id !== id)
    } catch (err: any) {
      error.value = err?.message || 'Unsuccessfully delete tax'
      throw err
    } finally {
      loading.value = false
    }
  }

  return {
    taxes,
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
  }
})
