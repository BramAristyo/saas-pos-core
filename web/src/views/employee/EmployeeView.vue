<script setup lang="ts">
import { onMounted, ref, watch } from 'vue'
import { useEmployeeStore } from '@/stores/employee.store'
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
import { MoreHorizontal, Plus, Search, Users } from 'lucide-vue-next'
import type { Employee } from '@/types/employee.types'
import EmployeeFormDialog from './EmployeeFormDialog.vue'
import EmployeeDeleteDialog from './EmployeeDeleteDialog.vue'
import { useSearch } from '@/composables/common/useSearch'
import { usePagination } from '@/composables/common/usePagination'
import { useFormatter } from '@/composables/common/useFormatter'
import { Input } from '@/components/ui/input'
import { CommonPagination } from '@/components/common/pagination'
import { SMALL_SIZE } from '@/constant/pagination.constant'
import { CommonEmpty } from '@/components/common/empty'
import { TableSkeleton } from '@/components/common/skeleton'
import { Badge } from '@/components/ui/badge'

const employeeStore = useEmployeeStore()
const { formatDate, formatRupiah } = useFormatter()

const isFormOpen = ref(false)
const isDeleteOpen = ref(false)
const selectedEmployee = ref<Employee | null>(null)

const { page, pageSize, setMeta, goToPage } = usePagination(SMALL_SIZE)
const { search } = useSearch(() => {
  page.value = 1
  loadData()
})

async function loadData() {
  await employeeStore.fetchPaginated({
    pageNumber: page.value,
    pageSize: pageSize.value,
    search: search.value || undefined,
  })
  if (employeeStore.meta) {
    setMeta(employeeStore.meta)
  }
}

watch(page, () => {
  loadData()
})

onMounted(() => {
  loadData()
})

function handleAdd() {
  selectedEmployee.value = null
  isFormOpen.value = true
}

function handleEdit(employee: Employee) {
  selectedEmployee.value = employee
  isFormOpen.value = true
}

function handleDelete(employee: Employee) {
  selectedEmployee.value = employee
  isDeleteOpen.value = true
}
</script>

<template>
  <AppLayout>
    <div class="flex flex-col sm:flex-row sm:items-center justify-between gap-4 mb-6">
      <h1 class="text-2xl font-bold">Employee</h1>
      <div class="flex items-center gap-2 flex-1 w-full sm:max-w-sm">
        <div class="relative w-full">
          <Search class="absolute left-2.5 top-2.5 h-4 w-4 text-muted-foreground" />
          <Input v-model="search" placeholder="Search employees..." class="pl-9" />
        </div>
        <Button @click="handleAdd">
          <Plus class="size-4 mr-2" />
          Add
        </Button>
      </div>
    </div>

    <TableSkeleton
      v-if="employeeStore.loading && employeeStore.employees.length === 0"
      :column-count="7"
    />

    <CommonEmpty
      v-else-if="employeeStore.employees.length === 0"
      title="Employees"
      description="Start by adding your first employee to manage your team."
      :icon="Users"
      :search="search"
      add-button-text="Add Employee"
      @add="handleAdd"
      @clear-search="search = ''"
    />

    <div v-else class="space-y-4">
      <div class="overflow-x-auto">
        <Table>
          <TableHeader>
            <TableRow>
              <TableHead>Code</TableHead>
              <TableHead>Name</TableHead>
              <TableHead>Phone</TableHead>
              <TableHead>Base Salary</TableHead>
              <TableHead>Created At</TableHead>
              <TableHead class="w-12.5"></TableHead>
            </TableRow>
          </TableHeader>
          <TableBody>
            <TableRow v-for="employee in employeeStore.employees" :key="employee.id">
              <TableCell class="font-medium">{{ employee.code }}</TableCell>
              <TableCell>{{ employee.name }}</TableCell>
              <TableCell>{{ employee.phone }}</TableCell>
              <TableCell>{{ formatRupiah(employee.baseSalary) }}</TableCell>
              <TableCell>{{ formatDate(employee.createdAt) }}</TableCell>
              <TableCell>
                <DropdownMenu>
                  <DropdownMenuTrigger as-child>
                    <Button variant="ghost" size="icon">
                      <MoreHorizontal class="size-4" />
                    </Button>
                  </DropdownMenuTrigger>
                  <DropdownMenuContent align="end">
                    <DropdownMenuItem @click="handleEdit(employee)"> Edit </DropdownMenuItem>
                    <DropdownMenuItem
                      class="text-destructive focus:text-destructive"
                      @click="handleDelete(employee)"
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
        v-if="employeeStore.meta"
        :page="page"
        :page-size="pageSize"
        :total-rows="employeeStore.meta.totalRows"
        :total-pages="employeeStore.meta.totalPages"
        @update:page="goToPage"
      />
    </div>

    <EmployeeFormDialog
      v-model:open="isFormOpen"
      :employee="selectedEmployee"
      @success="loadData"
    />

    <EmployeeDeleteDialog
      v-model:open="isDeleteOpen"
      :employee="selectedEmployee"
      @success="loadData"
    />
  </AppLayout>
</template>
