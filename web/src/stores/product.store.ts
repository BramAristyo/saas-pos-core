import { productApi } from '@/api/product.api'
import type { BaseFilterRequest, Meta } from '@/types/common.types'
import type { CreateProductRequest, Product, UpdateProductRequest } from '@/types/product.types'
import { defineStore } from 'pinia'
import { ref } from 'vue'

export const useProductStore = defineStore('product', () => {
  const products = ref<Product[]>([])
  const meta = ref<Meta | null>(null)
  const selected = ref<Product | null>(null)
  const loading = ref(false)
  const error = ref<string | null>(null)

  async function fetchAll() {
    loading.value = true
    error.value = null
    try {
      const res = await productApi.getAll()
      products.value = res.data
    } catch (err: any) {
      error.value = err?.message || 'Unsuccessfully get products'
    } finally {
      loading.value = false
    }
  }

  async function fetchPaginated(params: BaseFilterRequest) {
    loading.value = true
    error.value = null
    try {
      const res = await productApi.paginate(params)
      products.value = res.data
      meta.value = res.meta || null
    } catch (err: any) {
      error.value = err?.message || 'Unsuccessfully get products'
    } finally {
      loading.value = false
    }
  }

  async function fetchById(id: string) {
    loading.value = true
    error.value = null
    try {
      const res = await productApi.getById(id)
      selected.value = res.data
    } catch (err: any) {
      error.value = err?.message || 'Unsuccessfully get product'
    } finally {
      loading.value = false
    }
  }

  async function create(payload: CreateProductRequest) {
    loading.value = true
    error.value = null
    try {
      const res = await productApi.create(payload)
      products.value = [res.data, ...products.value]
    } catch (err: any) {
      error.value = err?.message || 'Unsuccessfully create product'
      throw err
    } finally {
      loading.value = false
    }
  }

  async function update(id: string, payload: UpdateProductRequest) {
    loading.value = true
    error.value = null
    try {
      const res = await productApi.update(id, payload)
      const index = products.value.findIndex((p) => p.id === id)
      if (index !== -1) products.value[index] = res.data
      if (selected.value?.id === id) selected.value = res.data
    } catch (err: any) {
      error.value = err?.message || 'Unsuccessfully update product'
      throw err
    } finally {
      loading.value = false
    }
  }

  async function remove(id: string) {
    loading.value = true
    error.value = null
    try {
      await productApi.delete(id)
      products.value = products.value.filter((p) => p.id !== id)
    } catch (err: any) {
      error.value = err?.message || 'Unsuccessfully delete product'
      throw err
    } finally {
      loading.value = false
    }
  }

  async function restore(id: string) {
    loading.value = true
    error.value = null
    try {
      await productApi.restore(id)
      // Since it was removed from the list on delete, we might want to re-fetch or just handle the list update
    } catch (err: any) {
      error.value = err?.message || 'Unsuccessfully restore product'
      throw err
    } finally {
      loading.value = false
    }
  }

  async function ensureDataLoaded() {
    if (products.value.length === 0) await fetchAll()
  }

  return {
    products,
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
