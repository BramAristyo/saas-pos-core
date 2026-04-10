<script setup lang="ts">
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
import { Button } from '@/components/ui/button'

interface Props {
  totalRows: number
  pageSize: number
  page: number
  totalPages: number
}

defineProps<Props>()

const emit = defineEmits<{
  (e: 'update:page', value: number): void
}>()
</script>

<template>
  <div v-if="totalPages > 1" class="flex items-center justify-between w-full mt-4">
    <!-- <div>
      <p class="text-sm text-muted-foreground">
        Showing {{ (page - 1) * pageSize + 1 }} to
        {{ Math.min(page * pageSize, totalRows) }} of {{ totalRows }} entries
      </p>
    </div> -->
    <Pagination
      v-slot="{ page: currentPage }"
      :total="totalRows"
      :sibling-count="1"
      :items-per-page="pageSize"
      show-edges
      :page="page"
      @update:page="emit('update:page', $event)"
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
</template>
