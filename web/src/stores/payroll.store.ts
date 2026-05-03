import { payrollApi } from '@/api/payroll.api'
import type { Payroll, CreatePayrollRequest } from '@/types/payroll.types'
import type { BaseFilterRequest, Meta } from '@/types/common.types'
import { defineStore } from 'pinia'
import { ref } from 'vue'

export const usePayrollStore = defineStore('payroll', () => {
  const payrolls = ref<Payroll[]>([])
  const meta = ref<Meta | null>(null)
  const loading = ref(false)
  const error = ref<string | null>(null)

  async function fetchPaginated(params: BaseFilterRequest) {
    loading.value = true
    error.value = null
    try {
      const res = await payrollApi.paginate(params)
      payrolls.value = res.data || []
      meta.value = res.meta || null
    } catch (err: any) {
      error.value = err?.message || 'Unsuccessfully get payrolls'
      payrolls.value = []
    } finally {
      loading.value = false
    }
  }

  async function create(payload: CreatePayrollRequest) {
    loading.value = true
    error.value = null
    try {
      const res = await payrollApi.create(payload)
      payrolls.value = [res.data, ...payrolls.value]
    } catch (err: any) {
      error.value = err?.message || 'Unsuccessfully create payroll'
      throw err
    } finally {
      loading.value = false
    }
  }

  return {
    payrolls,
    meta,
    loading,
    error,
    fetchPaginated,
    create,
  }
})
