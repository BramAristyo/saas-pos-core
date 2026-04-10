<script setup lang="ts">
import { onMounted, ref, watch } from 'vue'
import { useSalesType } from '@/composables/useSalesType'
import AppLayout from '@/layouts/AppLayout.vue'
import { Button } from '@/components/ui/button'
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
import { MoreHorizontal, Plus, Search, Tag } from 'lucide-vue-next'
import type { SalesType } from '@/types/salesType.types'
import SalesTypeFormDialog from './SalesTypeFormDialog.vue'
import SalesTypeDeleteDialog from './SalesTypeDeleteDialog.vue'
import { useSearch } from '@/composables/common/useSearch'
import { usePagination } from '@/composables/common/usePagination'
import { Input } from '@/components/ui/input'
import { CommonPagination } from '@/components/common/pagination'
import { SMALL_SIZE } from '@/constant/pagination.constant'
import { Skeleton } from '@/components/ui/skeleton'
import { CommonEmpty } from '@/components/common/empty'

const { salesTypes, loading, meta, fetchPaginated } = useSalesType()

const isFormOpen = ref(false)
const isDeleteOpen = ref(false)
const selectedSalesType = ref<SalesType | null>(null)

const { page, pageSize, setMeta, goToPage } = usePagination(SMALL_SIZE)
const { search } = useSearch(() => {
  page.value = 1
  loadData()
})

async function loadData() {
  await fetchPaginated({
    pageNumber: page.value,
    pageSize: pageSize.value,
    search: search.value || undefined,
  })
  if (meta.value) {
    setMeta(meta.value)
  }
}

watch(page, () => {
  loadData()
})

onMounted(() => {
  loadData()
})

function handleAdd() {
  selectedSalesType.value = null
  isFormOpen.value = true
}

function handleEdit(st: SalesType) {
  selectedSalesType.value = st
  isFormOpen.value = true
}

function handleDelete(st: SalesType) {
  selectedSalesType.value = st
  isDeleteOpen.value = true
}
</script>

<template>
  <AppLayout>
    <div class="flex flex-col md:flex-row md:items-center justify-between gap-4 mb-6">
      <h1 class="text-2xl font-bold">Sales Type</h1>
      <div class="flex items-center gap-2 flex-1 md:max-w-sm">
        <div class="relative w-full">
          <Search class="absolute left-2.5 top-2.5 h-4 w-4 text-muted-foreground" />
          <Input v-model="search" placeholder="Search sales types..." class="pl-9" />
        </div>
        <Button @click="handleAdd">
          <Plus class="size-4 mr-2" />
          Add
        </Button>
      </div>
    </div>

    <TableSkeleton v-if="loading && salesTypes.length === 0" :column-count="3" />

    <CommonEmpty
      v-else-if="salesTypes.length === 0"
      title="Sales Types"
      description="Create your first sales type to define how orders are handled (e.g., Dine In, Takeaway)."
      :icon="Tag"
      :search="search"
      add-button-text="Create Sales Type"
      @add="handleAdd"
      @clear-search="search = ''"
    />

    <div v-else class="space-y-4">
      <div class="overflow-hidden">
        <Table>
          <TableHeader>
            <TableRow>
              <TableHead>Name</TableHead>
              <TableHead>Created At</TableHead>
              <TableHead class="w-12.5"></TableHead>
            </TableRow>
          </TableHeader>
          <TableBody>
            <TableRow v-for="st in salesTypes" :key="st.id">
              <TableCell class="font-medium">{{ st.name }}</TableCell>
              <TableCell>{{ st.createdAt }}</TableCell>
              <TableCell>
                <DropdownMenu>
                  <DropdownMenuTrigger as-child>
                    <Button variant="ghost" size="icon">
                      <MoreHorizontal class="size-4" />
                    </Button>
                  </DropdownMenuTrigger>
                  <DropdownMenuContent align="end">
                    <DropdownMenuItem @click="handleEdit(st)"> Edit </DropdownMenuItem>
                    <DropdownMenuItem
                      class="text-destructive focus:text-destructive"
                      @click="handleDelete(st)"
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
        v-if="meta"
        :page="page"
        :page-size="pageSize"
        :total-rows="meta.totalRows"
        :total-pages="meta.totalPages"
        @update:page="goToPage"
      />
    </div>

    <SalesTypeFormDialog
      v-model:open="isFormOpen"
      :sales-type="selectedSalesType"
      @success="loadData"
    />

    <SalesTypeDeleteDialog
      v-model:open="isDeleteOpen"
      :sales-type="selectedSalesType"
      @success="loadData"
    />
  </AppLayout>
</template>
