<script setup lang="ts">
import {
  Dialog,
  DialogContent,
  DialogDescription,
  DialogFooter,
  DialogHeader,
  DialogTitle,
} from '@/components/ui/dialog'
import { Button } from '@/components/ui/button'
import { AlertTriangle } from 'lucide-vue-next'

const props = defineProps<{
  open: boolean
}>()

const emit = defineEmits<{
  (e: 'update:open', value: boolean): void
  (e: 'confirm'): void
}>()

function handleConfirm() {
  emit('confirm')
  emit('update:open', false)
}
</script>

<template>
  <Dialog :open="open" @update:open="emit('update:open', $event)">
    <DialogContent class="sm:max-w-md">
      <DialogHeader>
        <div class="flex items-center gap-2">
          <div class="p-2 rounded-full bg-destructive/10 text-destructive">
            <AlertTriangle class="size-5" />
          </div>
          <DialogTitle>Discard Changes?</DialogTitle>
        </div>
        <DialogDescription class="pt-2">
          Are you sure you want to discard your changes? This action cannot be undone and all unsaved progress will be lost.
        </DialogDescription>
      </DialogHeader>
      <DialogFooter class="mt-4 flex flex-col-reverse sm:flex-row sm:justify-end gap-2">
        <Button variant="ghost" @click="emit('update:open', false)">
          Keep Editing
        </Button>
        <Button variant="destructive" @click="handleConfirm">
          Discard Changes
        </Button>
      </DialogFooter>
    </DialogContent>
  </Dialog>
</template>
