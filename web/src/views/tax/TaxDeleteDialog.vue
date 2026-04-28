<script setup lang="ts">
import { computed } from 'vue'
import { useTaxStore } from '@/stores/tax.stores'
import type { Tax } from '@/types/tax.types'
import {
  Dialog,
  DialogContent,
  DialogDescription,
  DialogFooter,
  DialogHeader,
  DialogTitle,
} from '@/components/ui/dialog'
import { Button } from '@/components/ui/button'
import { toast } from 'vue-sonner'

const props = defineProps<{
  open: boolean
  tax: Tax | null
}>()

const emit = defineEmits<{
  (e: 'update:open', value: boolean): void
  (e: 'success'): void
}>()

const taxStore = useTaxStore()

const isOpen = computed({
  get: () => props.open,
  set: (value) => emit('update:open', value),
})

async function handleDelete() {
  if (!props.tax) return
  try {
    await taxStore.remove(props.tax.id)
    toast.success('Tax deleted successfully')
    emit('success')
    isOpen.value = false
  } catch (error: any) {
    toast.error(error?.message || 'Failed to delete tax')
  }
}
</script>

<template>
  <Dialog v-model:open="isOpen">
    <DialogContent class="sm:max-w-106.25">
      <DialogHeader>
        <DialogTitle>Delete Tax</DialogTitle>
        <DialogDescription>
          Are you sure you want to delete <strong>{{ tax?.name }}</strong
          >? This action cannot be undone.
        </DialogDescription>
      </DialogHeader>

      <DialogFooter class="flex gap-2 justify-end">
        <Button type="button" variant="outline" @click="isOpen = false"> Cancel </Button>
        <Button variant="destructive" :disabled="taxStore.loading" @click="handleDelete">
          {{ taxStore.loading ? 'Deleting...' : 'Delete' }}
        </Button>
      </DialogFooter>
    </DialogContent>
  </Dialog>
</template>
