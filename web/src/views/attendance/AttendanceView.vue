<script setup lang="ts">
import { onMounted, watch } from 'vue'
import { useAttendanceStore } from '@/stores/attendance.store'
import AppLayout from '@/layouts/AppLayout.vue'
import {
  Table,
  TableBody,
  TableCell,
  TableHead,
  TableHeader,
  TableRow,
} from '@/components/ui/table'
import { Search, ClipboardList } from 'lucide-vue-next'
import { useSearch } from '@/composables/common/useSearch'
import { usePagination } from '@/composables/common/usePagination'
import { useFormatter } from '@/composables/common/useFormatter'
import { Input } from '@/components/ui/input'
import { CommonPagination } from '@/components/common/pagination'
import { SMALL_SIZE } from '@/constant/pagination.constant'
import { CommonEmpty } from '@/components/common/empty'
import { TableSkeleton } from '@/components/common/skeleton'

const attendanceStore = useAttendanceStore()
const { formatDateOnly, formatRupiah, formatDuration } = useFormatter()

const { page, pageSize, setMeta, goToPage } = usePagination(SMALL_SIZE)
const { search } = useSearch(() => {
  page.value = 1
  loadData()
})

async function loadData() {
  await attendanceStore.fetchAttendances({
    pageNumber: page.value,
    pageSize: pageSize.value,
    search: search.value || undefined,
  })
  if (attendanceStore.meta) {
    setMeta(attendanceStore.meta)
  }
}

watch(page, () => {
  loadData()
})

onMounted(() => {
  loadData()
})
</script>

<template>
  <AppLayout>
    <div class="flex flex-col sm:flex-row sm:items-center justify-between gap-4 mb-6">
      <h1 class="text-2xl font-bold">Attendance</h1>
      <div class="flex items-center gap-2 flex-1 w-full sm:max-w-sm">
        <div class="relative w-full">
          <Search class="absolute left-2.5 top-2.5 h-4 w-4 text-muted-foreground" />
          <Input v-model="search" placeholder="Search attendance..." class="pl-9" />
        </div>
      </div>
    </div>

    <TableSkeleton
      v-if="attendanceStore.loading && attendanceStore.attendances.length === 0"
      :column-count="9"
    />

    <CommonEmpty
      v-else-if="attendanceStore.attendances.length === 0"
      title="Attendance"
      description="No attendance records found."
      :icon="ClipboardList"
      :search="search"
      @clear-search="search = ''"
    />

    <div v-else class="space-y-4">
      <div class="overflow-x-auto">
        <Table>
          <TableHeader>
            <TableRow>
              <TableHead>Employee</TableHead>
              <TableHead>Date</TableHead>
              <TableHead>Shift</TableHead>
              <TableHead>Check In</TableHead>
              <TableHead>Check Out</TableHead>
              <TableHead>Late</TableHead>
              <TableHead>Total Work</TableHead>
              <TableHead>Deduction</TableHead>
              <TableHead class="w-12.5"></TableHead>
            </TableRow>
          </TableHeader>
          <TableBody>
            <TableRow
              v-for="attendance in attendanceStore.attendances"
              :key="attendance.id"
              class="h-14"
            >
              <TableCell class="max-w-30.5 truncate" :title="attendance.employeeName">
                {{ attendance.employeeName }}
              </TableCell>
              <TableCell>{{ formatDateOnly(attendance.date) }}</TableCell>
              <TableCell>{{ attendance.shiftScheduleName }}</TableCell>
              <TableCell>{{ attendance.checkIn }}</TableCell>
              <TableCell>{{ attendance.checkOut || '-' }}</TableCell>
              <TableCell>{{ formatDuration(attendance.lateMinutes) }}</TableCell>
              <TableCell>{{ formatDuration(attendance.totalWorkMinutes) }}</TableCell>
              <TableCell>{{ formatRupiah(attendance.deductionAmount) }}</TableCell>
              <TableCell></TableCell>
            </TableRow>
          </TableBody>
        </Table>
      </div>

      <CommonPagination
        v-if="attendanceStore.meta"
        :page="page"
        :page-size="pageSize"
        :total-rows="attendanceStore.meta.totalRows"
        :total-pages="attendanceStore.meta.totalPages"
        @update:page="goToPage"
      />
    </div>
  </AppLayout>
</template>
