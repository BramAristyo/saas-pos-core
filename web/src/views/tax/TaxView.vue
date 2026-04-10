<script setup lang="ts">
import { onMounted, ref, watch } from 'vue'
import { useTaxStore } from '@/stores/tax.stores'
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
import type { Tax } from '@/types/tax.types'
import TaxFormDialog from './TaxFormDialog.vue'
import TaxDeleteDialog from './TaxDeleteDialog.vue'
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

const taxStore = useTaxStore()

const isFormOpen = ref(false)
const isDeleteOpen = ref(false)
const selectedTax = ref<Tax | null>(null)

const { page, pageSize, setMeta, goToPage } = usePagination(SMALL_SIZE)
const { search } = useSearch(() => {
  page.value = 1
  loadData()
})

async function loadData() {
  await taxStore.fetchPaginated({
    pageNumber: page.value,
    pageSize: pageSize.value,
    search: search.value || undefined,
  })
  if (taxStore.meta) {
    setMeta(taxStore.meta)
  }
}

watch(page, () => {
  loadData()
})

onMounted(() => {
  loadData()
})

function handleAdd() {
  selectedTax.value = null
  isFormOpen.value = true
}

function handleEdit(tax: Tax) {
  selectedTax.value = tax
  isFormOpen.value = true
}

function handleDelete(tax: Tax) {
  selectedTax.value = tax
  isDeleteOpen.value = true
}
</script>

<template>
  <AppLayout>
    <div class="flex flex-col md:flex-row md:items-center justify-between gap-4 mb-6">
      <h1 class="text-2xl font-bold">Tax</h1>
      <div class="flex items-center gap-2 flex-1 md:max-w-sm">
        <div class="relative w-full">
          <Search class="absolute left-2.5 top-2.5 h-4 w-4 text-muted-foreground" />
          <Input v-model="search" placeholder="Search taxes..." class="pl-9" />
        </div>
        <Button @click="handleAdd">
          <Plus class="size-4 mr-2" />
          Add
        </Button>
      </div>
    </div>

    <div v-if="taxStore.loading && taxStore.taxes.length === 0" class="flex justify-center py-8">
      <p>Loading...</p>
    </div>

    <div
      v-else-if="taxStore.taxes.length === 0"
      class="flex flex-col items-center justify-center py-12 border rounded-lg bg-muted/20"
    >
      <p class="text-muted-foreground mb-4">
        {{ search ? 'No taxes match your search' : 'No taxes found' }}
      </p>
      <Button v-if="!search" variant="outline" @click="handleAdd">Create your first tax</Button>
      <Button v-else variant="outline" @click="search = ''">Clear search</Button>
    </div>

    <div v-else class="space-y-4">
      <div class="overflow-hidden">
        <Table>
          <TableHeader>
            <TableRow>
              <TableHead>Name</TableHead>
              <TableHead>Percentage</TableHead>
              <TableHead class="w-12.5"></TableHead>
            </TableRow>
          </TableHeader>
          <TableBody>
            <TableRow v-for="tax in taxStore.taxes" :key="tax.id">
              <TableCell class="font-medium">{{ tax.name }}</TableCell>
              <TableCell>{{ tax.percentage }}%</TableCell>
              <TableCell>
                <DropdownMenu>
                  <DropdownMenuTrigger as-child>
                    <Button variant="ghost" size="icon">
                      <MoreHorizontal class="size-4" />
                    </Button>
                  </DropdownMenuTrigger>
                  <DropdownMenuContent align="end">
                    <DropdownMenuItem @click="handleEdit(tax)"> Edit </DropdownMenuItem>
                    <DropdownMenuItem
                      class="text-destructive focus:text-destructive"
                      @click="handleDelete(tax)"
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

      <div
        v-if="taxStore.meta && taxStore.meta.totalPages > 1"
        class="flex items-center justify-between w-full mt-4"
      >
        <div></div>
        <Pagination
          v-slot="{ page: currentPage }"
          :total="taxStore.meta.totalRows"
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

    <TaxFormDialog v-model:open="isFormOpen" :tax="selectedTax" @success="loadData" />

    <TaxDeleteDialog v-model:open="isDeleteOpen" :tax="selectedTax" @success="loadData" />
  </AppLayout>
</template>
