<script setup lang="ts">
import { ref, onMounted, computed } from 'vue'
import { useRouter } from 'vue-router'
import { Button } from '@/components/ui/button'
import { Input } from '@/components/ui/input'
import { Label } from '@/components/ui/label'
import { Switch } from '@/components/ui/switch'
import { Card, CardContent, CardHeader, CardTitle } from '@/components/ui/card'
import { Plus, Trash2, Package } from 'lucide-vue-next'
import type {
  ModifierGroupDetail,
  CreateModifierGroupRequest,
  UpdateModifierGroupRequest,
  CreateModifierOptionRequest,
  UpdateModifierOptionRequest,
} from '@/types/modifier.types'
import ProductSelectModal from '@/components/common/product/ProductSelectModal.vue'
import { useProductStore } from '@/stores/product.store'
import { toast } from 'vue-sonner'

const props = defineProps<{
  initialData?: ModifierGroupDetail | null
  loading?: boolean
}>()

const emit = defineEmits<{
  (e: 'submit', data: CreateModifierGroupRequest | UpdateModifierGroupRequest): void
}>()

const router = useRouter()
const productStore = useProductStore()

const name = ref(props.initialData?.name || '')
const isRequired = ref(props.initialData?.isRequired || false)
const options = ref<(CreateModifierOptionRequest | UpdateModifierOptionRequest)[]>(
  props.initialData?.options.map(o => ({
    id: o.id,
    name: o.name,
    priceAdjustment: o.priceAdjustment,
    cogsAdjustment: o.cogsAdjustment
  })) || [{ name: '', priceAdjustment: '0', cogsAdjustment: '0' }]
)
const selectedProductIds = ref<string[]>(
  props.initialData?.productModifiers.map((p) => p.id) || []
)

const isProductModalOpen = ref(false)

const selectedProducts = computed(() => {
  return productStore.products.filter((p) => selectedProductIds.value.includes(p.id))
})

function addOption() {
  options.value.push({ name: '', priceAdjustment: '0', cogsAdjustment: '0' })
}

function removeOption(index: number) {
  if (options.value.length > 1) {
    options.value.splice(index, 1)
  } else {
    toast.error('At least one option is required')
  }
}

function handleProductSelect(ids: string[]) {
  selectedProductIds.value = ids
}

function handleSubmit() {
  if (!name.value) {
    toast.error('Name is required')
    return
  }

  const payload: CreateModifierGroupRequest = {
    name: name.value,
    isRequired: isRequired.value,
    options: options.value.map((o) => ({
      ...o,
      priceAdjustment: o.priceAdjustment.toString(),
      cogsAdjustment: o.cogsAdjustment.toString(),
    })),
    productModifiers: selectedProductIds.value.length > 0 ? selectedProductIds.value : null,
  }

  emit('submit', payload)
}

onMounted(async () => {
  await productStore.ensureDataLoaded()
})
</script>

<template>
  <div class="space-y-6 max-w-4xl mx-auto pb-10">
    <div class="flex items-center justify-between">
      <h2 class="text-xl font-semibold">{{ initialData ? 'Edit' : 'Create' }} Modifier Group</h2>
      <div class="flex gap-2">
        <Button variant="outline" @click="router.back()">Cancel</Button>
        <Button :disabled="loading" @click="handleSubmit">
          {{ loading ? 'Saving...' : 'Save Modifier Group' }}
        </Button>
      </div>
    </div>

    <div class="grid grid-cols-1 md:grid-cols-3 gap-6">
      <!-- General Info -->
      <Card class="md:col-span-2">
        <CardHeader>
          <CardTitle>General Information</CardTitle>
        </CardHeader>
        <CardContent class="space-y-4">
          <div class="space-y-2">
            <Label for="name">Group Name</Label>
            <Input id="name" v-model="name" placeholder="e.g. Extra Toppings, Sugar Level" />
          </div>
          <div class="flex items-center justify-between p-4 border rounded-lg">
            <div class="space-y-0.5">
              <Label>Required</Label>
              <p class="text-sm text-muted-foreground">
                Customer must select at least one option
              </p>
            </div>
            <Switch v-model:checked="isRequired" />
          </div>
        </CardContent>
      </Card>

      <!-- Assignment -->
      <Card>
        <CardHeader class="flex flex-row items-center justify-between space-y-0">
          <CardTitle>Products</CardTitle>
          <Button variant="ghost" size="sm" @click="isProductModalOpen = true">
            <Plus class="size-4 mr-1" />
            Assign
          </Button>
        </CardHeader>
        <CardContent>
          <div v-if="selectedProducts.length === 0" class="text-center py-6 text-muted-foreground">
            <Package class="size-8 mx-auto mb-2 opacity-20" />
            <p class="text-xs">No products assigned</p>
          </div>
          <div v-else class="space-y-2">
            <div
              v-for="product in selectedProducts"
              :key="product.id"
              class="flex items-center justify-between text-sm p-2 bg-accent/50 rounded"
            >
              <span class="truncate pr-2">{{ product.name }}</span>
              <Button
                variant="ghost"
                size="icon"
                class="size-6 text-destructive h-auto w-auto p-1"
                @click="selectedProductIds = selectedProductIds.filter(id => id !== product.id)"
              >
                <Trash2 class="size-3" />
              </Button>
            </div>
          </div>
        </CardContent>
      </Card>

      <!-- Options -->
      <Card class="md:col-span-3">
        <CardHeader class="flex flex-row items-center justify-between space-y-0">
          <CardTitle>Modifier Options</CardTitle>
          <Button variant="outline" size="sm" @click="addOption">
            <Plus class="size-4 mr-1" />
            Add Option
          </Button>
        </CardHeader>
        <CardContent>
          <div class="space-y-4">
            <div
              v-for="(option, index) in options"
              :key="index"
              class="grid grid-cols-1 md:grid-cols-12 gap-4 items-end border-b pb-4 last:border-0 last:pb-0"
            >
              <div class="md:col-span-5 space-y-2">
                <Label v-if="index === 0">Option Name</Label>
                <Input v-model="option.name" placeholder="e.g. Extra Cheese" />
              </div>
              <div class="md:col-span-3 space-y-2">
                <Label v-if="index === 0">Price Adj.</Label>
                <Input v-model="option.priceAdjustment" type="number" step="0.01" />
              </div>
              <div class="md:col-span-3 space-y-2">
                <Label v-if="index === 0">COGS Adj.</Label>
                <Input v-model="option.cogsAdjustment" type="number" step="0.01" />
              </div>
              <div class="md:col-span-1 flex justify-end">
                <Button
                  variant="ghost"
                  size="icon"
                  class="text-destructive"
                  :disabled="options.length === 1"
                  @click="removeOption(index)"
                >
                  <Trash2 class="size-4" />
                </Button>
              </div>
            </div>
          </div>
        </CardContent>
      </Card>
    </div>

    <ProductSelectModal
      v-model:open="isProductModalOpen"
      :selected-ids="selectedProductIds"
      @select="handleProductSelect"
    />
  </div>
</template>
