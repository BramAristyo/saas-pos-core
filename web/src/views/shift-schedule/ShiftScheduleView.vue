<script setup lang="ts">
import { onMounted, ref, watch } from 'vue'
import { useShiftScheduleStore } from '@/stores/shiftSchedule.store'
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
import { MoreHorizontal, Plus, Search, CalendarClock } from 'lucide-vue-next'
import type { ShiftSchedule } from '@/types/shiftSchedule.types'
import ShiftScheduleFormDialog from './ShiftScheduleFormDialog.vue'
import ShiftScheduleDeleteDialog from './ShiftScheduleDeleteDialog.vue'
import { useSearch } from '@/composables/common/useSearch'
import { usePagination } from '@/composables/common/usePagination'
import { useFormatter } from '@/composables/common/useFormatter'
import { Input } from '@/components/ui/input'
import { CommonPagination } from '@/components/common/pagination'
import { SMALL_SIZE } from '@/constant/pagination.constant'
import { CommonEmpty } from '@/components/common/empty'
import { TableSkeleton } from '@/components/common/skeleton'

const shiftScheduleStore = useShiftScheduleStore()
const { formatRupiah } = useFormatter()

const isFormOpen = ref(false)
const isDeleteOpen = ref(false)
const selectedShift = ref<ShiftSchedule | null>(null)

const { page, pageSize, setMeta, goToPage } = usePagination(SMALL_SIZE)
const { search } = useSearch(() => {
  page.value = 1
  loadData()
})

async function loadData() {
  await shiftScheduleStore.fetchShiftSchedules({
    pageNumber: page.value,
    pageSize: pageSize.value,
    search: search.value || undefined,
  })
  if (shiftScheduleStore.meta) {
    setMeta(shiftScheduleStore.meta)
  }
}

watch(page, () => {
  loadData()
})

onMounted(() => {
  loadData()
})

function handleAdd() {
  selectedShift.value = null
  isFormOpen.value = true
}

function handleEdit(shift: ShiftSchedule) {
  selectedShift.value = shift
  isFormOpen.value = true
}

function handleDelete(shift: ShiftSchedule) {
  selectedShift.value = shift
  isDeleteOpen.value = true
}
</script>

<template>
  <AppLayout>
    <div class="flex flex-col sm:flex-row sm:items-center justify-between gap-4 mb-6">
      <h1 class="text-2xl font-bold">Shift Schedule</h1>
      <div class="flex items-center gap-2 flex-1 w-full sm:max-w-sm">
        <div class="relative w-full">
          <Search class="absolute left-2.5 top-2.5 h-4 w-4 text-muted-foreground" />
          <Input v-model="search" placeholder="Search shifts..." class="pl-9" />
        </div>
        <Button @click="handleAdd">
          <Plus class="size-4 mr-2" />
          Add
        </Button>
      </div>
    </div>

    <TableSkeleton
      v-if="shiftScheduleStore.loading && shiftScheduleStore.shiftSchedules.length === 0"
      :column-count="7"
    />

    <CommonEmpty
      v-else-if="shiftScheduleStore.shiftSchedules.length === 0"
      title="Shift Schedules"
      description="No shift schedules found. Start by adding a new shift schedule."
      :icon="CalendarClock"
      :search="search"
      add-button-text="Add Shift Schedule"
      @add="handleAdd"
      @clear-search="search = ''"
    />

    <div v-else class="space-y-4">
      <div class="overflow-x-auto">
        <Table>
          <TableHeader>
            <TableRow>
              <TableHead>Name</TableHead>
              <TableHead>Start Time</TableHead>
              <TableHead>End Time</TableHead>
              <TableHead>Tolerance</TableHead>
              <TableHead>Late Int.</TableHead>
              <TableHead>Deduction</TableHead>
              <TableHead class="w-12.5"></TableHead>
            </TableRow>
          </TableHeader>
          <TableBody>
            <TableRow v-for="shift in shiftScheduleStore.shiftSchedules" :key="shift.id" class="h-14">
              <TableCell class="font-medium">{{ shift.name }}</TableCell>
              <TableCell>{{ shift.startTime }}</TableCell>
              <TableCell>{{ shift.endTime }}</TableCell>
              <TableCell>{{ shift.toleranceMinutes }} Minutes</TableCell>
              <TableCell>{{ shift.lateIntervalMinutes }} Minutes</TableCell>
              <TableCell>{{ formatRupiah(shift.lateDeductionAmount) }}</TableCell>
              <TableCell>
                <DropdownMenu>
                  <DropdownMenuTrigger as-child>
                    <Button variant="ghost" size="icon">
                      <MoreHorizontal class="size-4" />
                    </Button>
                  </DropdownMenuTrigger>
                  <DropdownMenuContent align="end">
                    <DropdownMenuItem @click="handleEdit(shift)"> Edit </DropdownMenuItem>
                    <DropdownMenuItem
                      class="text-destructive focus:text-destructive"
                      @click="handleDelete(shift)"
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
        v-if="shiftScheduleStore.meta"
        :page="page"
        :page-size="pageSize"
        :total-rows="shiftScheduleStore.meta.totalRows"
        :total-pages="shiftScheduleStore.meta.totalPages"
        @update:page="goToPage"
      />
    </div>

    <ShiftScheduleFormDialog
      v-model:open="isFormOpen"
      :shift-schedule="selectedShift"
      @success="loadData"
    />

    <ShiftScheduleDeleteDialog
      v-model:open="isDeleteOpen"
      :shift-schedule="selectedShift"
      @success="loadData"
    />
  </AppLayout>
</template>
