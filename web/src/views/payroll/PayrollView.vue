<script setup lang="ts">
import { onMounted, ref, watch } from 'vue'
import { usePayrollStore } from '@/stores/payroll.store'
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
import { Plus, Search, Wallet } from 'lucide-vue-next'
import PayrollFormDialog from './PayrollFormDialog.vue'
import { useSearch } from '@/composables/common/useSearch'
import { usePagination } from '@/composables/common/usePagination'
import { useFormatter } from '@/composables/common/useFormatter'
import { Input } from '@/components/ui/input'
import { CommonPagination } from '@/components/common/pagination'
import { MEDIUM_SIZE } from '@/constant/pagination.constant'
import { CommonEmpty } from '@/components/common/empty'
import { TableSkeleton } from '@/components/common/skeleton'

const payrollStore = usePayrollStore()
const { formatRupiah, formatDateOnly } = useFormatter()

const isFormOpen = ref(false)

const { page, pageSize, setMeta, goToPage } = usePagination(MEDIUM_SIZE)
const { search } = useSearch(() => {
  page.value = 1
  loadData()
})

async function loadData() {
  await payrollStore.fetchPaginated({
    pageNumber: page.value,
    pageSize: pageSize.value,
    search: search.value || undefined,
  })
  if (payrollStore.meta) {
    setMeta(payrollStore.meta)
  }
}

watch(page, () => {
  loadData()
})

onMounted(() => {
  loadData()
})

function handleAdd() {
  isFormOpen.value = true
}
</script>

<template>
  <AppLayout>
    <div class="flex flex-col sm:flex-row sm:items-center justify-between gap-4 mb-6">
      <h1 class="text-2xl font-bold">Payroll</h1>
      <div class="flex items-center gap-2 w-full sm:max-w-sm">
        <div class="relative flex-1">
          <Search class="absolute left-2.5 top-2.5 h-4 w-4 text-muted-foreground" />
          <Input v-model="search" placeholder="Search payrolls..." class="pl-9" />
        </div>
        <Button @click="handleAdd">
          <Plus class="size-4 mr-2" />
          Add
        </Button>
      </div>
    </div>

    <TableSkeleton v-if="payrollStore.loading && payrollStore.payrolls.length === 0" :column-count="5" />

    <CommonEmpty
      v-else-if="payrollStore.payrolls.length === 0"
      title="Payrolls"
      description="No payroll records found. Start by creating your first payroll record."
      :icon="Wallet"
      :search="search"
      add-button-text="Create Payroll"
      @add="handleAdd"
      @clear-search="search = ''"
    />

    <div v-else class="space-y-4">
      <div class="overflow-x-auto border rounded-lg">
        <Table>
          <TableHeader>
            <TableRow>
              <TableHead>Employee</TableHead>
              <TableHead>Period</TableHead>
              <TableHead class="text-right">Base Salary</TableHead>
              <TableHead class="text-right">Total Deduction</TableHead>
              <TableHead class="text-right font-bold">Net Salary</TableHead>
            </TableRow>
          </TableHeader>
          <TableBody>
            <TableRow v-for="payroll in payrollStore.payrolls" :key="payroll.id">
              <TableCell>
                <div class="flex flex-col">
                  <span class="font-medium">{{ payroll.employeeName }}</span>
                  <span class="text-xs text-muted-foreground">{{ payroll.employeeCode }}</span>
                </div>
              </TableCell>
              <TableCell>
                <div class="text-sm">
                  {{ formatDateOnly(payroll.periodStart) }} - {{ formatDateOnly(payroll.periodEnd) }}
                </div>
              </TableCell>
              <TableCell class="text-right">{{ formatRupiah(payroll.baseSalary) }}</TableCell>
              <TableCell class="text-right text-destructive">
                -{{ formatRupiah(payroll.totalDeduction) }}
              </TableCell>
              <TableCell class="text-right font-bold text-primary">
                {{ formatRupiah(payroll.netSalary) }}
              </TableCell>
            </TableRow>
          </TableBody>
        </Table>
      </div>

      <CommonPagination
        v-if="payrollStore.meta"
        :page="page"
        :page-size="pageSize"
        :total-rows="payrollStore.meta.totalRows"
        :total-pages="payrollStore.meta.totalPages"
        @update:page="goToPage"
      />
    </div>

    <PayrollFormDialog
      v-model:open="isFormOpen"
      @success="loadData"
    />
  </AppLayout>
</template>
