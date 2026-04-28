<script setup lang="ts">
import { computed } from 'vue'
import { useModifierStore } from '@/stores/modifier.store'
import type { ModifierGroup } from '@/types/modifier.types'
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
  modifier: ModifierGroup | null
}>()

const emit = defineEmits<{
  (e: 'update:open', value: boolean): void
  (e: 'success'): void
}>()

const modifierStore = useModifierStore()

const isOpen = computed({
  get: () => props.open,
  set: (value) => emit('update:open', value),
})

async function handleRestore() {
  if (!props.modifier) return
  try {
    await modifierStore.restore(props.modifier.id)
    toast.success('Modifier group restored successfully')
    emit('success')
    isOpen.value = false
  } catch (error: any) {
    toast.error(error?.message || 'Failed to restore modifier group')
  }
}
</script>

<template>
  <Dialog v-model:open="isOpen">
    <DialogContent class="sm:max-w-[425px]">
      <DialogHeader>
        <DialogTitle>Restore Modifier Group</DialogTitle>
        <DialogDescription>
          Are you sure you want to restore <strong>{{ modifier?.name }}</strong
          >? It will be active and available for use again.
        </DialogDescription>
      </DialogHeader>

      <DialogFooter class="flex gap-2 justify-end">
        <Button type="button" variant="outline" @click="isOpen = false"> Cancel </Button>
        <Button variant="default" :disabled="modifierStore.loading" @click="handleRestore">
          {{ modifierStore.loading ? 'Restoring...' : 'Restore' }}
        </Button>
      </DialogFooter>
    </DialogContent>
  </Dialog>
</template>
