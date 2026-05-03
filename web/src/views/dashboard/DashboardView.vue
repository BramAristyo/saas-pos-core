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
  Percent,
  ArrowUpRight,
  ArrowDownRight,
  Sparkles,
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

// Fantastic layout mapping
const secondaryMetrics = [
  {
    title: 'Gross Sales',
    key: 'grossSales',
    isMoney: true,
    icon: DollarSign,
    trend: '+12.5%',
    trendUp: true,
    bg: 'bg-blue-500/5',
    text: 'text-blue-600',
  },
  {
    title: 'Gross Profit',
    key: 'grossProfit',
    isMoney: true,
    icon: BarChart3,
    trend: '+8.2%',
    trendUp: true,
    bg: 'bg-violet-500/5',
    text: 'text-violet-600',
  },
  {
    title: 'Avg. Sales',
    key: 'averageSales',
    isMoney: true,
    icon: Receipt,
    trend: '-2.1%',
    trendUp: false,
    bg: 'bg-amber-500/5',
    text: 'text-amber-600',
  },
  {
    title: 'Gross Margin',
    key: 'grossMargin',
    isMoney: false,
    isPercent: true,
    icon: Percent,
    trend: '+0.5%',
    trendUp: true,
    bg: 'bg-rose-500/5',
    text: 'text-rose-600',
  },
]
</script>

<template>
  <AppLayout>
    <div class="flex flex-col gap-8 py-6 max-w-7xl mx-auto">
      <!-- High-End Header -->
      <div class="flex flex-col md:flex-row md:items-center justify-between gap-6 relative">
        <div class="space-y-1">
          <h1
            class="text-4xl font-black tracking-tight bg-linear-to-br from-foreground to-foreground/70 bg-clip-text text-transparent"
          >
            Hey, {{ authStore.user?.name.split(' ')[0] }}!
          </h1>
          <p class="text-muted-foreground text-sm">
            Everything looks good. You've had
            <span class="text-foreground font-bold">{{ summary?.transactionCount || 0 }}</span>
            transactions {{ filterType === 'day' ? 'today' : 'this month' }}.
          </p>
        </div>

        <div class="flex items-center gap-4 bg-background border shadow-sm p-1.5 rounded-2xl">
          <Toggle v-model="filterType" :options="filterOptions" />
        </div>
      </div>

      <!-- Fantastic Main Grid -->
      <div class="grid grid-cols-1 lg:grid-cols-12 gap-6">
        <!-- Hero Stat: Net Sales -->
        <Card
          class="lg:col-span-7 overflow-hidden border-none bg-primary shadow-2xl shadow-primary/20 relative group"
        >
          <!-- Glassmorphism Effect -->
          <div
            class="absolute inset-0 bg-[radial-gradient(circle_at_top_right,rgba(255,255,255,0.1),transparent)] pointer-events-none"
          ></div>
          <div
            class="absolute -right-20 -top-20 size-64 bg-white/5 rounded-full blur-3xl group-hover:bg-white/10 transition-all duration-700"
          ></div>

          <CardContent
            class="p-8 md:p-10 relative z-10 h-full flex flex-col justify-between min-h-70"
          >
            <div class="flex items-start justify-between">
              <div class="space-y-1">
                <p class="text-primary-foreground/60 text-sm font-bold uppercase tracking-widest">
                  Net Revenue
                </p>
                <h2 class="text-5xl md:text-6xl font-black text-white tracking-tighter">
                  <template v-if="loading">
                    <Skeleton class="h-16 w-64 bg-white/10" />
                  </template>
                  <template v-else-if="summary">
                    {{ formatRupiah(summary.netSales) }}
                  </template>
                </h2>
              </div>
              <div class="p-3 bg-white/10 backdrop-blur-md rounded-2xl border border-white/10">
                <TrendingUp class="size-8 text-white" />
              </div>
            </div>

            <div class="flex flex-wrap items-center gap-6 mt-8">
              <div class="flex flex-col gap-1">
                <span class="text-primary-foreground/50 text-[10px] font-bold uppercase"
                  >Growth</span
                >
                <div
                  class="flex items-center gap-1.5 px-2 py-1 bg-emerald-500/20 rounded-full border border-emerald-500/20 text-emerald-300 text-xs font-bold"
                >
                  <ArrowUpRight class="size-3" />
                  <span>+24.8%</span>
                </div>
              </div>
              <div class="w-px h-8 bg-white/10"></div>
              <div class="flex flex-col gap-1 text-white">
                <span class="text-primary-foreground/50 text-[10px] font-bold uppercase"
                  >Volume</span
                >
                <span class="text-lg font-bold"
                  >{{ summary?.transactionCount || 0 }}
                  <span class="text-[10px] text-primary-foreground/50">txns</span></span
                >
              </div>
            </div>
          </CardContent>
        </Card>

        <!-- Secondary Stats Grid -->
        <div class="lg:col-span-5 grid grid-cols-1 sm:grid-cols-2 gap-4">
          <template v-if="loading">
            <Card v-for="i in 4" :key="i" class="border shadow-none bg-muted/20">
              <CardContent class="p-5 space-y-3">
                <Skeleton class="h-4 w-16" />
                <Skeleton class="h-8 w-32" />
              </CardContent>
            </Card>
          </template>

          <template v-else-if="summary">
            <Card
              v-for="metric in secondaryMetrics"
              :key="metric.key"
              class="group hover:border-primary/50 hover:shadow-lg hover:shadow-primary/5 transition-all duration-300 border-border/60"
            >
              <CardContent class="p-5 flex flex-col justify-between h-full min-h-32.5">
                <div class="flex items-center justify-between mb-4">
                  <div :class="[metric.bg, 'p-2 rounded-lg']">
                    <component :is="metric.icon" :class="['size-4', metric.text]" />
                  </div>
                  <div
                    :class="[
                      metric.trendUp ? 'text-emerald-600' : 'text-rose-600',
                      'flex items-center gap-0.5 text-[10px] font-bold bg-muted/50 px-1.5 py-0.5 rounded-md border border-border/50',
                    ]"
                  >
                    <component
                      :is="metric.trendUp ? ArrowUpRight : ArrowDownRight"
                      class="size-3"
                    />
                    {{ metric.trend }}
                  </div>
                </div>

                <div class="space-y-0.5">
                  <p class="text-[10px] font-bold uppercase tracking-wider text-muted-foreground">
                    {{ metric.title }}
                  </p>
                  <div class="text-xl font-black tracking-tight">
                    <template v-if="metric.isMoney">
                      {{ formatRupiah(summary[metric.key as keyof typeof summary]) }}
                    </template>
                    <template v-else-if="metric.isPercent">
                      {{ summary[metric.key as keyof typeof summary] }}%
                    </template>
                  </div>
                </div>
              </CardContent>
            </Card>
          </template>
        </div>
      </div>

      <!-- Footer/Next Features Placeholder -->
      <div
        class="mt-4 p-8 rounded-4xl border-2 border-dashed border-muted flex flex-col md:flex-row items-center justify-between gap-6 opacity-80 hover:opacity-100 transition-opacity"
      >
        <div
          class="px-6 py-2 bg-muted text-muted-foreground text-xs font-bold rounded-full uppercase tracking-widest border border-border w-full"
        >
          Under Development
        </div>
      </div>
    </div>
  </AppLayout>
</template>
