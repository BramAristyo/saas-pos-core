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
  <div class="flex items-center justify-between w-full mt-4">
    <p class="text-sm text-muted-foreground whitespace-nowrap">
      Showing
      {{ (page - 1) * pageSize + 1 }}
      to
      {{ Math.min(page * pageSize, totalRows) }}
      of
      {{ totalRows }} entries
    </p>

    <Pagination
      v-slot="{ page: currentPage }"
      :total="totalRows"
      :sibling-count="1"
      :items-per-page="pageSize"
      show-edges
      :page="page"
      @update:page="emit('update:page', $event)"
    >
      <PaginationContent v-slot="{ items }" class="flex items-center gap-2">
        <PaginationFirst />
        <PaginationPrevious />

        <template v-for="(item, index) in items" :key="index">
          <PaginationItem v-if="item.type === 'page'" :value="item.value" as-child>
            <Button
              class="h-8 w-8 p-0"
              :variant="item.value === currentPage ? 'default' : 'outline'"
            >
              {{ item.value }}
            </Button>
          </PaginationItem>
          <PaginationEllipsis v-else :index="index" />
        </template>

        <PaginationNext />
        <PaginationLast />
      </PaginationContent>
    </Pagination>
  </div>
</template>
