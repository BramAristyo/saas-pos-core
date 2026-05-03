<script setup lang="ts">
import { computed } from 'vue'
import { useShiftScheduleStore } from '@/stores/shiftSchedule.store'
import type { ShiftSchedule } from '@/types/shiftSchedule.types'
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
  shiftSchedule: ShiftSchedule | null
}>()

const emit = defineEmits<{
  (e: 'update:open', value: boolean): void
  (e: 'success'): void
}>()

const shiftScheduleStore = useShiftScheduleStore()

const isOpen = computed({
  get: () => props.open,
  set: (value) => emit('update:open', value),
})

async function handleDelete() {
  if (!props.shiftSchedule) return
  try {
    await shiftScheduleStore.remove(props.shiftSchedule.id)
    toast.success('Shift schedule deleted successfully')
    emit('success')
    isOpen.value = false
  } catch (error: any) {
    toast.error(error?.message || 'Failed to delete shift schedule')
  }
}
</script>

<template>
  <Dialog v-model:open="isOpen">
    <DialogContent class="sm:max-w-[425px]">
      <DialogHeader>
        <DialogTitle>Delete Shift Schedule</DialogTitle>
        <DialogDescription>
          Are you sure you want to delete <strong>{{ shiftSchedule?.name }}</strong
          >? This action cannot be undone.
        </DialogDescription>
      </DialogHeader>

      <DialogFooter class="flex gap-2 justify-end">
        <Button type="button" variant="outline" @click="isOpen = false"> Cancel </Button>
        <Button variant="destructive" :disabled="shiftScheduleStore.loading" @click="handleDelete">
          {{ shiftScheduleStore.loading ? 'Deleting...' : 'Delete' }}
        </Button>
      </DialogFooter>
    </DialogContent>
  </Dialog>
</template>
