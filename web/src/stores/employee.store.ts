import { employeeApi } from '@/api/employee.api'
import type { BaseFilterRequest, Meta } from '@/types/common.types'
import type { CreateEmployeeRequest, Employee, UpdateEmployeeRequest } from '@/types/employee.types'
import { defineStore } from 'pinia'
import { ref } from 'vue'

export const useEmployeeStore = defineStore('employee', () => {
  const employees = ref<Employee[]>([])
  const meta = ref<Meta | null>(null)
  const selected = ref<Employee | null>(null)
  const loading = ref(false)
  const error = ref<string | null>(null)

  async function fetchEmployees(params: BaseFilterRequest) {
    loading.value = true
    error.value = null
    try {
      const res = await employeeApi.paginate(params)
      employees.value = res.data
      meta.value = res.meta || null
    } catch (err: unknown) {
      const errorResponse = err as { message?: string }
      error.value = errorResponse.message || 'Unsuccessfully get employees'
    } finally {
      loading.value = false
    }
  }

  async function fetchAllEmployees() {
    loading.value = true
    error.value = null
    try {
      const res = await employeeApi.getAll()
      employees.value = res.data
    } catch (err: unknown) {
      const errorResponse = err as { message?: string }
      error.value = errorResponse.message || 'Unsuccessfully get employees'
    } finally {
      loading.value = false
    }
  }

  async function fetchEmployeeById(id: string) {
    loading.value = true
    error.value = null
    try {
      const res = await employeeApi.getById(id)
      selected.value = res.data
    } catch (err: unknown) {
      const errorResponse = err as { message?: string }
      error.value = errorResponse.message || 'Unsuccessfully get employee'
    } finally {
      loading.value = false
    }
  }

  async function createEmployee(payload: CreateEmployeeRequest) {
    loading.value = true
    error.value = null
    try {
      const res = await employeeApi.create(payload)
      employees.value = [res.data, ...employees.value]
    } catch (err: unknown) {
      const errorResponse = err as { message?: string }
      error.value = errorResponse.message || 'Unsuccessfully create employee'
      throw err
    } finally {
      loading.value = false
    }
  }

  async function updateEmployee(id: string, payload: UpdateEmployeeRequest) {
    loading.value = true
    error.value = null
    try {
      const res = await employeeApi.update(id, payload)
      const index = employees.value.findIndex((e) => e.id === id)
      if (index !== -1) employees.value[index] = res.data
      if (selected.value?.id === id) selected.value = res.data
    } catch (err: unknown) {
      const errorResponse = err as { message?: string }
      error.value = errorResponse.message || 'Unsuccessfully update employee'
      throw err
    } finally {
      loading.value = false
    }
  }

  async function deleteEmployee(id: string) {
    loading.value = true
    error.value = null
    try {
      await employeeApi.delete(id)
      employees.value = employees.value.filter((e) => e.id !== id)
    } catch (err: unknown) {
      const errorResponse = err as { message?: string }
      error.value = errorResponse.message || 'Unsuccessfully delete employee'
      throw err
    } finally {
      loading.value = false
    }
  }

  async function restoreEmployee(id: string) {
    loading.value = true
    error.value = null
    try {
      await employeeApi.restore(id)
    } catch (err: unknown) {
      const errorResponse = err as { message?: string }
      error.value = errorResponse.message || 'Unsuccessfully restore employee'
      throw err
    } finally {
      loading.value = false
    }
  }

  async function ensureDataLoaded() {
    if (employees.value.length === 0) await fetchAllEmployees()
  }

  return {
    employees,
    meta,
    selected,
    loading,
    error,
    fetchEmployees: fetchEmployees, // Keep consistent with UI usage if needed
    fetchPaginated: fetchEmployees, // Map to what UI expects
    fetchAllEmployees,
    fetchEmployeeById,
    create: createEmployee,
    update: updateEmployee,
    remove: deleteEmployee,
    restoreEmployee,
    ensureDataLoaded,
  }
})
