<script setup lang="ts">
import { computed, onMounted, ref, watch } from 'vue'
import { useProductStore } from '@/stores/product.store'
import {
  Dialog,
  DialogContent,
  DialogHeader,
  DialogTitle,
  DialogFooter,
} from '@/components/ui/dialog'
import { Button } from '@/components/ui/button'
import { Input } from '@/components/ui/input'
import { Search, Package } from 'lucide-vue-next'
import { ScrollArea } from '@/components/ui/scroll-area'
import { useFormatter } from '@/composables/common/useFormatter'

const props = defineProps<{
  open: boolean
  selectedIds: string[]
}>()

const emit = defineEmits<{
  (e: 'update:open', value: boolean): void
  (e: 'select', ids: string[]): void
}>()

const productStore = useProductStore()
const { formatRupiah } = useFormatter()
const searchQuery = ref('')
const localSelectedIds = ref<string[]>([])

const isOpen = computed({
  get: () => props.open,
  set: (value) => emit('update:open', value),
})

onMounted(async () => {
  await productStore.ensureDataLoaded()
})

watch(
  () => props.open,
  (newVal) => {
    if (newVal) {
      localSelectedIds.value = [...props.selectedIds]
    }
  },
  { immediate: true },
)

const filteredProducts = computed(() => {
  if (!searchQuery.value) return productStore.products
  return productStore.products.filter((p) =>
    p.name.toLowerCase().includes(searchQuery.value.toLowerCase()),
  )
})

function toggleProduct(id: string) {
  const index = localSelectedIds.value.indexOf(id)
  if (index === -1) {
    localSelectedIds.value = [...localSelectedIds.value, id]
  } else {
    localSelectedIds.value = localSelectedIds.value.filter((i) => i !== id)
  }
}

function handleSave() {
  emit('select', localSelectedIds.value)
  isOpen.value = false
}
</script>

<template>
  <Dialog v-model:open="isOpen">
    <DialogContent class="w-[95vw] sm:max-w-125 gap-0 p-0 overflow-hidden rounded-lg">
      <DialogHeader class="p-6 border-b">
        <DialogTitle>Select Products</DialogTitle>
      </DialogHeader>

      <div class="p-4 border-b">
        <div class="relative w-full">
          <Search class="absolute left-2.5 top-2.5 h-4 w-4 text-muted-foreground" />
          <Input v-model="searchQuery" placeholder="Search products..." class="pl-9" />
        </div>
      </div>

      <ScrollArea class="h-[60vh] sm:h-100 p-4">
        <div
          v-if="filteredProducts.length === 0"
          class="flex flex-col items-center justify-center py-10 text-muted-foreground"
        >
          <Package class="size-10 mb-2 opacity-20" />
          <p>No products found</p>
        </div>
        <div v-else class="space-y-1">
          <div
            v-for="product in filteredProducts"
            :key="product.id"
            :class="[
              'flex items-center space-x-3 p-3 rounded-lg cursor-pointer transition-colors',
              localSelectedIds.includes(product.id)
                ? 'bg-primary/10 text-primary hover:bg-primary/15'
                : 'hover:bg-accent/50 text-foreground',
            ]"
            @click="toggleProduct(product.id)"
          >
            <div class="flex-1 min-w-0">
              <p class="text-sm font-medium leading-none truncate">{{ product.name }}</p>
              <p
                :class="[
                  'text-xs mt-1',
                  localSelectedIds.includes(product.id) ? 'opacity-80' : 'text-muted-foreground',
                ]"
              >
                {{ product.category?.name }}
              </p>
            </div>
            <div class="text-sm font-medium">
              {{ formatRupiah(product.price) }}
            </div>
          </div>
        </div>
      </ScrollArea>

      <DialogFooter class="p-6 border-t bg-muted/50">
        <div class="flex items-center justify-between w-full">
          <p class="text-sm text-muted-foreground">
            {{ localSelectedIds.length }} product(s) selected
          </p>
          <div class="flex gap-2">
            <Button variant="outline" @click="isOpen = false">Cancel</Button>
            <Button @click="handleSave">Save Selection</Button>
          </div>
        </div>
      </DialogFooter>
    </DialogContent>
  </Dialog>
</template>
