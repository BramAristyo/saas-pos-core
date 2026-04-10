<script setup lang="ts">
import {
  Empty,
  EmptyContent,
  EmptyDescription,
  EmptyHeader,
  EmptyMedia,
  EmptyTitle,
} from '@/components/ui/empty'
import { Button } from '@/components/ui/button'
import { Plus } from 'lucide-vue-next'
import type { Component } from 'vue'

interface Props {
  title?: string
  description?: string
  icon?: Component
  search?: string
  addButtonText?: string
}

const props = withDefaults(defineProps<Props>(), {
  title: 'No items found',
  description: 'No items found. Adjust your filters or create a new one to get started.',
  addButtonText: 'Create New',
})

const emit = defineEmits<{
  (e: 'add'): void
  (e: 'clearSearch'): void
}>()
</script>

<template>
  <Empty>
    <EmptyContent>
      <EmptyMedia v-if="icon" variant="icon">
        <component :is="icon" class="size-6" />
      </EmptyMedia>
      <EmptyHeader>
        <EmptyTitle>{{ search ? `No ${title.toLowerCase()} match` : title }}</EmptyTitle>
        <EmptyDescription>
          {{
            search
              ? `We couldn't find any ${title.toLowerCase()} matching your search. Try a different term or clear the search.`
              : description
          }}
        </EmptyDescription>
      </EmptyHeader>
      <div class="flex items-center gap-2">
        <Button variant="outline" @click="emit('add')">
          <Plus class="size-4 mr-2" />
          {{ addButtonText }}
        </Button>
        <Button v-if="search" variant="ghost" @click="emit('clearSearch')"> Clear search </Button>
      </div>
    </EmptyContent>
  </Empty>
</template>
