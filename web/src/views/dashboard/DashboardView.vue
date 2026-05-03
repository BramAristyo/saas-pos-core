<script setup lang="ts">
import { onMounted, watch } from 'vue'
import { useAuthStore } from '@/stores/auth.stores'
import { useDashboard } from '@/composables/useDashboard'
import { useFormatter } from '@/composables/common/useFormatter'
import AppLayout from '@/layouts/AppLayout.vue'
import { Card, CardContent, CardHeader, CardTitle } from '@/components/ui/card'
import { Toggle } from '@/components/common/form'
import { Skeleton } from '@/components/ui/skeleton'

const authStore = useAuthStore()
const { summary, loading, filterType, fetchSummary } = useDashboard()
const { formatRupiah } = useFormatter()

const filterOptions = [
  { label: 'Day', value: 'day' },
  { label: 'Month', value: 'month' },
]

onMounted(fetchSummary)

watch(filterType, fetchSummary)

const metrics = [
  { title: 'Gross Sales', key: 'grossSales', isMoney: true },
  { title: 'Net Sales', key: 'netSales', isMoney: true },
  { title: 'Gross Profit', key: 'grossProfit', isMoney: true },
  { title: 'Transactions', key: 'transactionCount', isMoney: false },
  { title: 'Average Sales', key: 'averageSales', isMoney: true },
  { title: 'Gross Margin', key: 'grossMargin', isMoney: false, isPercent: true },
]
</script>

<template>
  <AppLayout>
    <div class="flex flex-col gap-6">
      <div class="flex flex-col sm:flex-row sm:items-center justify-between gap-4">
        <h1 class="text-2xl font-bold">Hello, {{ authStore.user?.name }}!</h1>
        <Toggle v-model="filterType" :options="filterOptions" />
      </div>

      <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-4">
        <template v-if="loading">
          <Card v-for="i in 6" :key="i" class="border-none shadow-none bg-muted/20">
            <CardHeader class="pb-2">
              <Skeleton class="h-4 w-24" />
            </CardHeader>
            <CardContent>
              <Skeleton class="h-8 w-32" />
            </CardContent>
          </Card>
        </template>

        <template v-else-if="summary">
          <Card v-for="metric in metrics" :key="metric.key" class="border-none shadow-none bg-muted/20">
            <CardHeader class="pb-2 text-muted-foreground text-sm font-medium uppercase tracking-wider">
              {{ metric.title }}
            </CardHeader>
            <CardContent class="text-2xl font-bold">
              <template v-if="metric.isMoney">
                {{ formatRupiah(summary[metric.key as keyof typeof summary]) }}
              </template>
              <template v-else-if="metric.isPercent">
                {{ (Number(summary[metric.key as keyof typeof summary]) * 100).toFixed(1) }}%
              </template>
              <template v-else>
                {{ summary[metric.key as keyof typeof summary] }}
              </template>
            </CardContent>
          </Card>
        </template>
      </div>
    </div>
  </AppLayout>
</template>
