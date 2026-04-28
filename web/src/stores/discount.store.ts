import { discountApi } from '@/api/discount.api'
import type { BaseFilterRequest, Meta } from '@/types/common.types'
import type { CreateDiscountRequest, Discount, UpdateDiscountRequest } from '@/types/discount.types'
import { defineStore } from 'pinia'
import { ref } from 'vue'

export const useDiscountStore = defineStore('discount', () => {
  const discounts = ref<Discount[]>([])
  const meta = ref<Meta | null>(null)
  const selected = ref<Discount | null>(null)
  const loading = ref(false)
  const error = ref<string | null>(null)

  async function fetchAll() {
    loading.value = true
    error.value = null
    try {
      const res = await discountApi.getAll()
      discounts.value = res.data
    } catch (err: any) {
      error.value = err?.message || 'Unsuccessfully get discounts'
    } finally {
      loading.value = false
    }
  }

  async function fetchPaginated(params: BaseFilterRequest) {
    loading.value = true
    error.value = null
    try {
      const res = await discountApi.paginate(params)
      discounts.value = res.data
      meta.value = res.meta || null
    } catch (err: any) {
      error.value = err?.message || 'Unsuccessfully get discounts'
    } finally {
      loading.value = false
    }
  }

  async function fetchById(id: string) {
    loading.value = true
    error.value = null
    try {
      const res = await discountApi.getById(id)
      selected.value = res.data
    } catch (err: any) {
      error.value = err?.message || 'Unsuccessfully get discount'
    } finally {
      loading.value = false
    }
  }

  async function create(payload: CreateDiscountRequest) {
    loading.value = true
    error.value = null
    try {
      const res = await discountApi.create(payload)
      discounts.value = [res.data, ...discounts.value]
    } catch (err: any) {
      error.value = err?.message || 'Unsuccessfully create discount'
      throw err
    } finally {
      loading.value = false
    }
  }

  async function update(id: string, payload: UpdateDiscountRequest) {
    loading.value = true
    error.value = null
    try {
      const res = await discountApi.update(id, payload)
      const index = discounts.value.findIndex((d) => d.id === id)
      if (index !== -1) discounts.value[index] = res.data
      if (selected.value?.id === id) {
        selected.value = res.data
      }
    } catch (err: any) {
      error.value = err?.message || 'Unsuccessfully update discount'
      throw err
    } finally {
      loading.value = false
    }
  }

  async function remove(id: string) {
    loading.value = true
    error.value = null
    try {
      await discountApi.delete(id)
      discounts.value = discounts.value.filter((d) => d.id !== id)
    } catch (err: any) {
      error.value = err?.message || 'Unsuccessfully delete discount'
      throw err
    } finally {
      loading.value = false
    }
  }

  async function ensureDataLoaded() {
    if (discounts.value.length === 0) await fetchAll()
  }

  return {
    discounts,
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
    ensureDataLoaded,
  }
})
