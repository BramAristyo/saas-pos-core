import { MEDIUM_SIZE } from '@/constant/pagination.constant'
import type { Meta } from '@/types/common.types'
import { computed, ref } from 'vue'

export function usePagination(defaultPageSize = MEDIUM_SIZE) {
  const page = ref(1)
  const pageSize = ref(defaultPageSize)
  const totalRows = ref(0)
  const totalPages = ref(0)

  const hasNext = computed(() => page.value < totalPages.value)
  const hasPrev = computed(() => page.value > 1)

  function nextPage() {
    if (hasNext) page.value++
  }

  function prevPage() {
    if (hasPrev) page.value--
  }

  function goToPage(target: number) {
    if (target > 0 && target <= totalPages.value) page.value = target
  }

  function setMeta(meta: Meta) {
    totalRows.value = meta.totalRows
    totalPages.value = meta.totalPages
  }

  function reset() {
    page.value
  }

  return {
    page,
    pageSize,
    totalRows,
    totalPages,
    hasNext,
    hasPrev,
    nextPage,
    prevPage,
    goToPage,
    setMeta,
    reset,
  }
}
