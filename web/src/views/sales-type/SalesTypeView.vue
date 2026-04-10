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
import { MoreHorizontal, Plus, Search } from 'lucide-vue-next'
import type { SalesType } from '@/types/salesType.types'
import SalesTypeFormDialog from './SalesTypeFormDialog.vue'
import SalesTypeDeleteDialog from './SalesTypeDeleteDialog.vue'
import { useSearch } from '@/composables/common/useSearch'
import { usePagination } from '@/composables/common/usePagination'
import { Input } from '@/components/ui/input'
import {
  Pagination,
  PaginationContent,
  PaginationEllipsis,
  PaginationFirst,
  PaginationItem,
  PaginationLast,
  PaginationNext,
  PaginationPrevious,
} from '@/components/ui/pagination'
import { SMALL_SIZE } from '@/constant/pagination.constant'

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

    <div v-if="loading && salesTypes.length === 0" class="flex justify-center py-8">
      <p>Loading...</p>
    </div>

    <div
      v-else-if="salesTypes.length === 0"
      class="flex flex-col items-center justify-center py-12 border rounded-lg bg-muted/20"
    >
      <p class="text-muted-foreground mb-4">
        {{ search ? 'No sales types match your search' : 'No sales types found' }}
      </p>
      <Button v-if="!search" variant="outline" @click="handleAdd"
        >Create your first sales type</Button
      >
      <Button v-else variant="outline" @click="search = ''">Clear search</Button>
    </div>

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

      <div v-if="meta && meta.totalPages > 1" class="flex items-center justify-between w-full mt-4">
        <div></div>
        <Pagination
          v-slot="{ page: currentPage }"
          :total="meta.totalRows"
          :sibling-count="1"
          :items-per-page="pageSize"
          show-edges
          :page="page"
          @update:page="goToPage"
        >
          <PaginationContent v-slot="{ items }">
            <PaginationFirst />
            <PaginationPrevious />

            <template v-for="(item, index) in items">
              <PaginationItem v-if="item.type === 'page'" :key="index" :value="item.value" as-child>
                <Button
                  class="w-8 h-8 p-0"
                  :variant="item.value === currentPage ? 'default' : 'outline'"
                >
                  {{ item.value }}
                </Button>
              </PaginationItem>
              <PaginationEllipsis v-else :key="item.type" :index="index" />
            </template>

            <PaginationNext />
            <PaginationLast />
          </PaginationContent>
        </Pagination>
      </div>
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
