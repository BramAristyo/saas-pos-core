import { ref, computed } from 'vue'
import { reportApi } from '@/api/report.api'
import type { SalesSummary } from '@/types/report.types'

export function useDashboard() {
  const summary = ref<SalesSummary | null>(null)
  const loading = ref(false)
  const error = ref<string | null>(null)
  const filterType = ref<'day' | 'month'>('day')

  const dateRange = computed(() => {
    const now = new Date()
    const from = new Date()
    const to = now.toISOString().split('T')[0]

    if (filterType.value === 'day') {
      // Current day: from and to are today
      return { from: to, to }
    } else {
      // Current month: from is 1st of month, to is today
      from.setDate(1)
      return { from: from.toISOString().split('T')[0], to }
    }
  })

  async function fetchSummary() {
    loading.value = true
    error.value = null
    try {
      const { from, to } = dateRange.value
      const res = await reportApi.getSalesSummary({
        'filter[created_at][type]': 'inRange',
        'filter[created_at][filterType]': 'date',
        'filter[created_at][from]': from as string,
        'filter[created_at][to]': to as string,
      })
      summary.value = res.data
    } catch (err: any) {
      error.value = err?.message || 'Failed to fetch sales summary'
    } finally {
      loading.value = false
    }
  }

  return {
    summary,
    loading,
    error,
    filterType,
    fetchSummary,
  }
}
