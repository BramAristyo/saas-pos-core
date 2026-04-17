<script setup lang="ts">
import { computed } from 'vue'
import { useDiscountStore } from '@/stores/discount.store'
import type { Discount } from '@/types/discount.types'
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
  discount: Discount | null
}>()

const emit = defineEmits<{
  (e: 'update:open', value: boolean): void
  (e: 'success'): void
}>()

const discountStore = useDiscountStore()

const isOpen = computed({
  get: () => props.open,
  set: (value) => emit('update:open', value),
})

async function handleDelete() {
  if (!props.discount) return
  try {
    await discountStore.remove(props.discount.id)
    toast.success('Discount deleted successfully')
    emit('success')
    isOpen.value = false
  } catch (error: any) {
    toast.error(error?.message || 'Failed to delete discount')
  }
}
</script>

<template>
  <Dialog v-model:open="isOpen">
    <DialogContent class="sm:max-w-[425px]">
      <DialogHeader>
        <DialogTitle>Delete Discount</DialogTitle>
        <DialogDescription>
          Are you sure you want to delete <strong>{{ discount?.name }}</strong
          >? This action cannot be undone.
        </DialogDescription>
      </DialogHeader>

      <DialogFooter class="flex gap-2 justify-end">
        <Button type="button" variant="outline" @click="isOpen = false"> Cancel </Button>
        <Button variant="destructive" :disabled="discountStore.loading" @click="handleDelete">
          {{ discountStore.loading ? 'Deleting...' : 'Delete' }}
        </Button>
      </DialogFooter>
    </DialogContent>
  </Dialog>
</template>
