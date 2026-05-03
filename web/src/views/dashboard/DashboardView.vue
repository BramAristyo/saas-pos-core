<script setup lang="ts">
import { onMounted, watch } from 'vue'
import { useAuthStore } from '@/stores/auth.stores'
import { useDashboard } from '@/composables/useDashboard'
import { useFormatter } from '@/composables/common/useFormatter'
import AppLayout from '@/layouts/AppLayout.vue'
import { Card, CardContent } from '@/components/ui/card'
import { Toggle } from '@/components/common/form'
import { Skeleton } from '@/components/ui/skeleton'
import { 
  TrendingUp, 
  Receipt, 
  DollarSign, 
  Hash, 
  BarChart3, 
  Percent 
} from 'lucide-vue-next'

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
  { 
    title: 'Gross Sales', 
    key: 'grossSales', 
    isMoney: true, 
    icon: DollarSign,
    color: 'border-blue-500' 
  },
  { 
    title: 'Net Sales', 
    key: 'netSales', 
    isMoney: true, 
    icon: TrendingUp,
    color: 'border-emerald-500' 
  },
  { 
    title: 'Gross Profit', 
    key: 'grossProfit', 
    isMoney: true, 
    icon: BarChart3,
    color: 'border-violet-500' 
  },
  { 
    title: 'Transactions', 
    key: 'transactionCount', 
    isMoney: false, 
    icon: Hash,
    color: 'border-amber-500' 
  },
  { 
    title: 'Average Sales', 
    key: 'averageSales', 
    isMoney: true, 
    icon: Receipt,
    color: 'border-indigo-500' 
  },
  { 
    title: 'Gross Margin', 
    key: 'grossMargin', 
    isMoney: false, 
    isPercent: true, 
    icon: Percent,
    color: 'border-rose-500' 
  },
]
</script>

<template>
  <AppLayout>
    <div class="flex flex-col gap-8 py-4">
      <!-- Header Section -->
      <div class="flex flex-col md:flex-row md:items-end justify-between gap-6">
        <div class="space-y-1">
          <h1 class="text-3xl font-extrabold tracking-tight">
            Welcome back, {{ authStore.user?.name }}
          </h1>
          <p class="text-muted-foreground">
            Here's what's happening with your sales for 
            <span class="font-medium text-foreground underline decoration-primary/30 decoration-2 underline-offset-4">
              {{ filterType === 'day' ? 'today' : 'this month' }}
            </span>.
          </p>
        </div>
        <div class="flex items-center gap-3 bg-muted/30 p-1.5 rounded-xl border border-border/50 backdrop-blur-sm">
          <span class="text-xs font-semibold px-2 text-muted-foreground uppercase tracking-wider">Range</span>
          <Toggle v-model="filterType" :options="filterOptions" />
        </div>
      </div>

      <!-- Stats Grid -->
      <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-5">
        <template v-if="loading">
          <Card 
            v-for="i in 6" 
            :key="i" 
            class="border-none bg-muted/20 animate-pulse"
          >
            <CardContent class="p-6 space-y-3">
              <Skeleton class="h-4 w-24 opacity-50" />
              <Skeleton class="h-9 w-40" />
            </CardContent>
          </Card>
        </template>

        <template v-else-if="summary">
          <Card 
            v-for="metric in metrics" 
            :key="metric.key" 
            class="group relative overflow-hidden border-none shadow-sm transition-all hover:shadow-md hover:-translate-y-0.5"
            :class="[metric.color, 'border-l-4 bg-card']"
          >
            <CardContent class="p-6">
              <div class="flex items-center justify-between mb-4">
                <span class="text-[11px] font-bold uppercase tracking-[0.1em] text-muted-foreground/80">
                  {{ metric.title }}
                </span>
                <component 
                  :is="metric.icon" 
                  class="size-5 text-muted-foreground/40 group-hover:text-foreground transition-colors" 
                />
              </div>
              
              <div class="flex flex-col gap-1">
                <div class="text-2xl font-bold tracking-tight">
                  <template v-if="metric.isMoney">
                    {{ formatRupiah(summary[metric.key as keyof typeof summary]) }}
                  </template>
                  <template v-else-if="metric.isPercent">
                    {{ summary[metric.key as keyof typeof summary] }}%
                  </template>
                  <template v-else>
                    {{ summary[metric.key as keyof typeof summary] }}
                  </template>
                </div>
                <div class="h-1 w-8 rounded-full bg-muted/40 transition-all group-hover:w-16 group-hover:bg-primary/40"></div>
              </div>
            </CardContent>
          </Card>
        </template>
      </div>
    </div>
  </AppLayout>
</template>
