<script setup lang="ts">
import { computed } from 'vue'
import { useSalesType } from '@/composables/useSalesType'
import type { SalesType } from '@/types/salesType.types'
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
  salesType: SalesType | null
}>()

const emit = defineEmits<{
  (e: 'update:open', value: boolean): void
  (e: 'success'): void
}>()

const { remove, loading } = useSalesType()

const isOpen = computed({
  get: () => props.open,
  set: (value) => emit('update:open', value),
})

async function handleDelete() {
  if (!props.salesType) return

  try {
    await remove(props.salesType.id)
    toast.success('Sales type deleted successfully')
    emit('success')
    isOpen.value = false
  } catch (err: any) {
    toast.error(err?.message || 'Failed to delete sales type')
  }
}
</script>

<template>
  <Dialog v-model:open="isOpen">
    <DialogContent class="sm:max-w-106.25">
      <DialogHeader>
        <DialogTitle>Delete Sales Type</DialogTitle>
        <DialogDescription>
          This will soft-delete the sales type <strong>{{ salesType?.name }}</strong>. You can
          restore it later if needed.
        </DialogDescription>
      </DialogHeader>
      <DialogFooter>
        <Button variant="outline" @click="isOpen = false">Cancel</Button>
        <Button
          variant="destructive"
          :disabled="loading"
          @click="handleDelete"
        >
          {{ loading ? 'Deleting...' : 'Delete' }}
        </Button>
      </DialogFooter>
    </DialogContent>
  </Dialog>
</template>
