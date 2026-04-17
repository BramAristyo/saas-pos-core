<script setup lang="ts">
import { onMounted, ref, watch } from 'vue'
import { useRouter } from 'vue-router'
import { useDiscountStore } from '@/stores/discount.store'
import AppLayout from '@/layouts/AppLayout.vue'
import { Button } from '@/components/ui/button'
import { Badge } from '@/components/ui/badge'
import {
  Table,
  TableBody,
  TableCell,
  TableHead,
  TableHeader,
  TableRow,
} from '@/components/ui/table'
import {
  DropdownMenu,
  DropdownMenuContent,
  DropdownMenuItem,
  DropdownMenuTrigger,
} from '@/components/ui/dropdown-menu'
import { MoreHorizontal, Plus, Search, Percent } from 'lucide-vue-next'
import type { Discount } from '@/types/discount.types'
import DiscountDeleteDialog from './DiscountDeleteDialog.vue'
import { useSearch } from '@/composables/common/useSearch'
import { usePagination } from '@/composables/common/usePagination'
import { useFormatter } from '@/composables/common/useFormatter'
import { Input } from '@/components/ui/input'
import { CommonPagination } from '@/components/common/pagination'
import { SMALL_SIZE } from '@/constant/pagination.constant'
import { CommonEmpty } from '@/components/common/empty'
import { TableSkeleton } from '@/components/common/skeleton'

const discountStore = useDiscountStore()
const router = useRouter()
const { formatDate, formatDateOnly, formatRupiah, formatPercent } = useFormatter()

const isDeleteOpen = ref(false)
const selectedDiscount = ref<Discount | null>(null)

const { page, pageSize, setMeta, goToPage } = usePagination(SMALL_SIZE)
const { search } = useSearch(() => {
  page.value = 1
  loadData()
})

async function loadData() {
  await discountStore.fetchPaginated({
    pageNumber: page.value,
    pageSize: pageSize.value,
    search: search.value || undefined,
  })
  if (discountStore.meta) {
    setMeta(discountStore.meta)
  }
}

watch(page, () => {
  loadData()
})

onMounted(() => {
  loadData()
})

function handleAdd() {
  router.push({ name: 'discount-create' })
}

function handleEdit(discount: Discount) {
  router.push({ name: 'discount-edit', params: { id: discount.id } })
}

function handleDelete(discount: Discount) {
  selectedDiscount.value = discount
  isDeleteOpen.value = true
}

function formatValue(discount: Discount) {
  if (discount.type === 'percentage') {
    return formatPercent(discount.value)
  }
  return formatRupiah(discount.value)
}
</script>

<template>
  <AppLayout>
    <div class="flex flex-col md:flex-row md:items-center justify-between gap-4 mb-6">
      <h1 class="text-2xl font-bold">Discounts</h1>
      <div class="flex items-center gap-2 flex-1 md:max-w-sm">
        <div class="relative w-full">
          <Search class="absolute left-2.5 top-2.5 h-4 w-4 text-muted-foreground" />
          <Input v-model="search" placeholder="Search discounts..." class="pl-9" />
        </div>
        <Button @click="handleAdd">
          <Plus class="size-4 mr-2" />
          Add
        </Button>
      </div>
    </div>

    <TableSkeleton
      v-if="discountStore.loading && discountStore.discounts.length === 0"
      :column-count="4"
    />

    <CommonEmpty
      v-else-if="!discountStore.loading && discountStore.discounts.length === 0"
      title="Discounts"
      description="Start by creating your first discount for your customers."
      :icon="Percent"
      :search="search"
      add-button-text="Create Discount"
      @add="handleAdd"
      @clear-search="search = ''"
    />

    <div v-else class="space-y-4">
      <div class="overflow-hidden">
        <Table>
          <TableHeader>
            <TableRow>
              <TableHead>Name</TableHead>
              <TableHead>Value</TableHead>
              <TableHead>Period</TableHead>
              <TableHead>Created At</TableHead>
              <TableHead class="w-12.5"></TableHead>
            </TableRow>
          </TableHeader>
          <TableBody>
            <TableRow v-for="discount in discountStore.discounts" :key="discount.id">
              <TableCell class="font-medium">
                {{ discount.name }}
              </TableCell>
              <TableCell>
                <div class="flex items-center gap-2">
                  <Badge variant="secondary" class="capitalize">
                    {{ discount.type }}
                  </Badge>
                  <span class="font-medium">{{ formatValue(discount) }}</span>
                </div>
              </TableCell>
              <TableCell>
                <div class="flex flex-col gap-0.5">
                  <span v-if="!discount.startDate && !discount.endDate" class="text-sm">Always active</span>
                  <span v-else-if="discount.startDate && !discount.endDate" class="text-sm">
                    From {{ formatDateOnly(discount.startDate) }}
                  </span>
                  <span v-else-if="!discount.startDate && discount.endDate" class="text-sm">
                    Until {{ formatDateOnly(discount.endDate) }}
                  </span>
                  <span v-else class="text-sm">
                    {{ formatDateOnly(discount.startDate) }} - {{ formatDateOnly(discount.endDate) }}
                  </span>
                </div>
              </TableCell>
              <TableCell>{{ formatDate(discount.createdAt) }}</TableCell>
              <TableCell>
                <DropdownMenu>
                  <DropdownMenuTrigger as-child>
                    <Button variant="ghost" size="icon">
                      <MoreHorizontal class="size-4" />
                    </Button>
                  </DropdownMenuTrigger>
                  <DropdownMenuContent align="end">
                    <DropdownMenuItem @click="handleEdit(discount)"> Edit </DropdownMenuItem>
                    <DropdownMenuItem
                      class="text-destructive focus:text-destructive"
                      @click="handleDelete(discount)"
                    >
                      Delete
                    </DropdownMenuItem>
                  </DropdownMenuContent>
                </DropdownMenu>
              </TableCell>
            </TableRow>
          </TableBody>
        </Table>
      </div>

      <CommonPagination
        v-if="discountStore.meta"
        :page="page"
        :page-size="pageSize"
        :total-rows="discountStore.meta.totalRows"
        :total-pages="discountStore.meta.totalPages"
        @update:page="goToPage"
      />
    </div>

    <DiscountDeleteDialog
      v-model:open="isDeleteOpen"
      :discount="selectedDiscount"
      @success="loadData"
    />
  </AppLayout>
</template>
