import { modifierApi } from '@/api/modifier.api'
import type { BaseFilterRequest, Meta } from '@/types/common.types'
import type {
  CreateModifierGroupRequest,
  ModifierGroup,
  ModifierGroupDetail,
  UpdateModifierGroupRequest,
} from '@/types/modifier.types'
import { defineStore } from 'pinia'
import { ref } from 'vue'

export const useModifierStore = defineStore('modifier', () => {
  const modifiers = ref<ModifierGroup[]>([])
  const meta = ref<Meta | null>(null)
  const selected = ref<ModifierGroupDetail | null>(null)
  const loading = ref(false)
  const error = ref<string | null>(null)

  async function fetchAll() {
    loading.value = true
    error.value = null
    try {
      const res = await modifierApi.getAll()
      modifiers.value = res.data
    } catch (err: any) {
      error.value = err?.message || 'Unsuccessfully get modifiers'
    } finally {
      loading.value = false
    }
  }

  async function fetchPaginated(params: BaseFilterRequest) {
    loading.value = true
    error.value = null
    try {
      const res = await modifierApi.paginate(params)
      modifiers.value = res.data
      meta.value = res.meta || null
    } catch (err: any) {
      error.value = err?.message || 'Unsuccessfully get modifiers'
    } finally {
      loading.value = false
    }
  }

  async function fetchById(id: string) {
    loading.value = true
    error.value = null
    try {
      const res = await modifierApi.getById(id)
      selected.value = res.data
    } catch (err: any) {
      error.value = err?.message || 'Unsuccessfully get modifier group'
    } finally {
      loading.value = false
    }
  }

  async function create(payload: CreateModifierGroupRequest) {
    loading.value = true
    error.value = null
    try {
      const res = await modifierApi.create(payload)
      modifiers.value = [res.data, ...modifiers.value]
    } catch (err: any) {
      error.value = err?.message || 'Unsuccessfully create modifier group'
      throw err
    } finally {
      loading.value = false
    }
  }

  async function update(id: string, payload: UpdateModifierGroupRequest) {
    loading.value = true
    error.value = null
    try {
      const res = await modifierApi.update(id, payload)
      const index = modifiers.value.findIndex((m) => m.id === id)
      if (index !== -1) modifiers.value[index] = res.data
      // If we update, we might want to refresh selected too if it's the same
      if (selected.value?.id === id) {
        await fetchById(id)
      }
    } catch (err: any) {
      error.value = err?.message || 'Unsuccessfully update modifier group'
      throw err
    } finally {
      loading.value = false
    }
  }

  async function remove(id: string) {
    loading.value = true
    error.value = null
    try {
      await modifierApi.delete(id)
      modifiers.value = modifiers.value.filter((m) => m.id !== id)
    } catch (err: any) {
      error.value = err?.message || 'Unsuccessfully delete modifier group'
      throw err
    } finally {
      loading.value = false
    }
  }

  async function restore(id: string) {
    loading.value = true
    error.value = null
    try {
      await modifierApi.restore(id)
    } catch (err: any) {
      error.value = err?.message || 'Unsuccessfully restore modifier group'
      throw err
    } finally {
      loading.value = false
    }
  }

  async function ensureDataLoaded() {
    if (modifiers.value.length === 0) await fetchAll()
  }

  return {
    modifiers,
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
