import { salesTypeApi } from '@/api/salesType.api'
import type { BaseFilterRequest, Meta } from '@/types/common.types'
import type {
  CreateSalesTypeRequest,
  SalesType,
  SalesTypeDetail,
  UpdateSalesTypeRequest,
} from '@/types/salesType.types'
import { ref } from 'vue'

const salesTypes = ref<SalesType[]>([])
const salesType = ref<SalesTypeDetail | null>(null)
const meta = ref<Meta | null>(null)
const loading = ref(false)
const error = ref<string | null>(null)

export const useSalesType = () => {
  async function fetchAll() {
    loading.value = true
    error.value = null
    try {
      const res = await salesTypeApi.getAll()
      salesTypes.value = res.data
    } catch (err: any) {
      error.value = err?.message || 'Unsuccessfully get sales types'
    } finally {
      loading.value = false
    }
  }

  async function fetchPaginated(params: BaseFilterRequest) {
    loading.value = true
    error.value = null
    try {
      const res = await salesTypeApi.paginate(params)
      salesTypes.value = res.data
      meta.value = res.meta || null
    } catch (err: any) {
      error.value = err?.message || 'Unsuccessfully get sales types'
    } finally {
      loading.value = false
    }
  }

  async function fetchById(id: string) {
    loading.value = true
    error.value = null
    try {
      const res = await salesTypeApi.getById(id)
      salesType.value = res.data
    } catch (err: any) {
      error.value = err?.message || 'Unsuccessfully get sales type'
      throw err
    } finally {
      loading.value = false
    }
  }

  async function create(payload: CreateSalesTypeRequest) {
    loading.value = true
    error.value = null
    try {
      const res = await salesTypeApi.create(payload)
      return res.data
    } catch (err: any) {
      error.value = err?.message || 'Unsuccessfully create sales type'
      throw err
    } finally {
      loading.value = false
    }
  }

  async function update(id: string, payload: UpdateSalesTypeRequest) {
    loading.value = true
    error.value = null
    try {
      const res = await salesTypeApi.update(id, payload)
      return res.data
    } catch (err: any) {
      error.value = err?.message || 'Unsuccessfully update sales type'
      throw err
    } finally {
      loading.value = false
    }
  }

  async function remove(id: string) {
    loading.value = true
    error.value = null
    try {
      await salesTypeApi.delete(id)
    } catch (err: any) {
      error.value = err?.message || 'Unsuccessfully delete sales type'
      throw err
    } finally {
      loading.value = false
    }
  }

  return {
    salesTypes,
    salesType,
    meta,
    loading,
    error,
    fetchAll,
    fetchPaginated,
    fetchById,
    create,
    update,
    remove,
  }
}
