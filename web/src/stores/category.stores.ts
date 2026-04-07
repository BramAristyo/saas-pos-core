import { categoryApi } from '@/api/category.api'
import type { Category, CreateCategoryRequest, UpdateCategoryRequest } from '@/types/category.types'
import { defineStore } from 'pinia'
import { ref } from 'vue'

export const useCategoryStore = defineStore('category', () => {
  const categories = ref<Category[]>([])
  const selected = ref<Category | null>(null)
  const loading = ref(false)
  const error = ref<string | null>(null)

  async function fetchAll() {
    loading.value = true
    error.value = null
    try {
      const res = await categoryApi.getAll()
      categories.value = res.data
    } catch (err: any) {
      error.value = err?.message || 'Unsuccessfully get categories'
    } finally {
      loading.value = true
    }
  }

  async function ensureDataLoaded() {
    if (categories.value.length === 0) {
      await fetchAll()
    }
  }

  async function fetchById(id: string) {
    loading.value = true
    error.value = null
    try {
      const res = await categoryApi.getById(id)
      selected.value = res.data
    } catch (err: any) {
      error.value = err?.message || 'Unsuccessfully get category'
    } finally {
      loading.value = false
    }
  }

  async function create(payload: CreateCategoryRequest) {
    loading.value = true
    error.value = null
    try {
      const res = await categoryApi.create(payload)
      categories.value.push(res.data)
    } catch (err: any) {
      error.value = err?.message || 'Unsuccessfully create category'
      throw err
    } finally {
      loading.value = false
    }
  }

  async function update(id: string, payload: UpdateCategoryRequest) {
    loading.value = true
    error.value = null
    try {
      const res = await categoryApi.update(id, payload)
      const index = categories.value.findIndex((c) => c.id === id)
      if (index !== -1) categories.value[index] = res.data
    } catch (err: any) {
      error.value = err?.message || 'Unsuccessfully update category'
      throw err
    } finally {
      loading.value = false
    }
  }

  async function remove(id: string) {
    loading.value = true
    error.value = null
    try {
      await categoryApi.delete(id)
      categories.value = categories.value.filter((c) => c.id !== id)
    } catch (err: any) {
      error.value = err?.message || 'Unsuccessfully delete category'
      throw err
    } finally {
      loading.value = false
    }
  }

  return {
    categories,
    selected,
    loading,
    error,
    fetchAll,
    fetchById,
    create,
    update,
    remove,
  }
})
