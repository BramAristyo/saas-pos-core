<script setup lang="ts">
import { reactive, watch, onMounted } from 'vue'
import { Input } from '@/components/ui/input'
import { Button } from '@/components/ui/button'
import {
  Field,
  FieldContent,
  FieldLabel,
  FieldError,
} from '@/components/ui/field'
import { AmountInput } from '@/components/common/form/input/amount'
import { Checkbox } from '@/components/ui/checkbox'
import { useCategoryStore } from '@/stores/category.stores'
import { useModifierStore } from '@/stores/modifier.store'
import { useFormErrors } from '@/composables/common/useFormErrors'
import type { Product, StoreProductRequest } from '@/types/product.types'

const props = defineProps<{
  initialData?: Product | null
  loading?: boolean
}>()

const emit = defineEmits<{
  (e: 'submit', data: StoreProductRequest): void
}>()

const categoryStore = useCategoryStore()
const modifierStore = useModifierStore()
const { getErrorMessage, hasError, setErrors, clearErrors } = useFormErrors()

const form = reactive<{
  name: string
  description: string
  price: string
  cogs: string
  categoryId: string
  modifierGroupIds: string[]
  imageUrl: string
}>({
  name: '',
  description: '',
  price: '0',
  cogs: '0',
  categoryId: '',
  modifierGroupIds: [],
  imageUrl: '',
})

watch(
  () => props.initialData,
  (newData) => {
    if (newData) {
      form.name = newData.name
      form.description = newData.description || ''
      form.price = newData.price.toString()
      form.cogs = newData.cogs.toString()
      form.categoryId = newData.categoryId
      form.modifierGroupIds = newData.modifierGroups?.map((mg) => mg.id) || []
      form.imageUrl = newData.imageUrl || ''
    } else {
      form.name = ''
      form.description = ''
      form.price = '0'
      form.cogs = '0'
      form.categoryId = ''
      form.modifierGroupIds = []
      form.imageUrl = ''
    }
  },
  { immediate: true },
)

onMounted(async () => {
  await Promise.all([
    categoryStore.categories.length === 0 ? categoryStore.fetchAll() : Promise.resolve(),
    modifierStore.modifiers.length === 0 ? modifierStore.fetchAll() : Promise.resolve(),
  ])
})

function handleModifierToggle(id: string, checked: boolean) {
  if (checked) {
    form.modifierGroupIds.push(id)
  } else {
    form.modifierGroupIds = form.modifierGroupIds.filter((mgId) => mgId !== id)
  }
}

const handleSubmit = () => {
  emit('submit', {
    name: form.name,
    description: form.description,
    price: Number(form.price),
    cogs: Number(form.cogs),
    categoryId: form.categoryId,
    modifierGroupIds: form.modifierGroupIds,
    imageUrl: form.imageUrl || undefined,
  })
}

defineExpose({
  setErrors,
  clearErrors,
})
</script>

<template>
  <form @submit.prevent="handleSubmit" class="space-y-6">
    <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
      <div class="space-y-4">
        <Field>
          <FieldLabel>Product Name</FieldLabel>
          <FieldContent>
            <Input
              v-model="form.name"
              placeholder="Enter product name"
              required
              :aria-invalid="hasError('Name')"
            />
            <FieldError v-if="hasError('Name')" :errors="[getErrorMessage('Name')]" />
          </FieldContent>
        </Field>

        <Field>
          <FieldLabel>Category</FieldLabel>
          <FieldContent>
            <select
              v-model="form.categoryId"
              class="flex h-9 w-full rounded-md border border-input bg-background px-3 py-1 text-sm shadow-sm transition-colors file:border-0 file:bg-transparent file:text-sm file:font-medium placeholder:text-muted-foreground focus-visible:outline-none focus-visible:ring-1 focus-visible:ring-ring disabled:cursor-not-allowed disabled:opacity-50 aria-invalid:border-destructive"
              required
              :aria-invalid="hasError('CategoryId')"
            >
              <option value="" disabled>Select a category</option>
              <option
                v-for="category in categoryStore.categories"
                :key="category.id"
                :value="category.id"
              >
                {{ category.name }}
              </option>
            </select>
            <FieldError v-if="hasError('CategoryId')" :errors="[getErrorMessage('CategoryId')]" />
          </FieldContent>
        </Field>

        <Field>
          <FieldLabel>Description</FieldLabel>
          <FieldContent>
            <textarea
              v-model="form.description"
              placeholder="Enter product description"
              class="flex min-h-[100px] w-full rounded-md border border-input bg-background px-3 py-2 text-sm shadow-sm placeholder:text-muted-foreground focus-visible:outline-none focus-visible:ring-1 focus-visible:ring-ring disabled:cursor-not-allowed disabled:opacity-50 aria-invalid:border-destructive"
              :aria-invalid="hasError('Description')"
            ></textarea>
            <FieldError v-if="hasError('Description')" :errors="[getErrorMessage('Description')]" />
          </FieldContent>
        </Field>

        <Field>
          <FieldLabel>Modifier Groups</FieldLabel>
          <div class="grid grid-cols-1 gap-2 border rounded-md p-4 max-h-[200px] overflow-y-auto bg-muted/20">
            <div
              v-for="mg in modifierStore.modifiers"
              :key="mg.id"
              class="flex items-center space-x-2"
            >
              <Checkbox
                :id="mg.id"
                :checked="form.modifierGroupIds.includes(mg.id)"
                @update:checked="(val: boolean) => handleModifierToggle(mg.id, val)"
              />
              <label
                :for="mg.id"
                class="text-sm font-medium leading-none cursor-pointer peer-disabled:cursor-not-allowed peer-disabled:opacity-70"
              >
                {{ mg.name }}
              </label>
            </div>
            <div v-if="modifierStore.modifiers.length === 0" class="text-sm text-muted-foreground italic">
              No modifier groups available.
            </div>
          </div>
          <FieldError
            v-if="hasError('ModifierGroupIds')"
            :errors="[getErrorMessage('ModifierGroupIds')]"
          />
        </Field>
      </div>

      <div class="space-y-4">
        <Field>
          <FieldLabel>Price</FieldLabel>
          <FieldContent>
            <AmountInput
              v-model="form.price"
              mode="money"
              placeholder="0"
              :aria-invalid="hasError('Price')"
            />
            <FieldError v-if="hasError('Price')" :errors="[getErrorMessage('Price')]" />
          </FieldContent>
        </Field>

        <Field>
          <FieldLabel>COGS (Cost of Goods Sold)</FieldLabel>
          <FieldContent>
            <AmountInput
              v-model="form.cogs"
              mode="money"
              placeholder="0"
              :aria-invalid="hasError('Cogs')"
            />
            <FieldError v-if="hasError('Cogs')" :errors="[getErrorMessage('Cogs')]" />
          </FieldContent>
        </Field>

        <Field>
          <FieldLabel>Image URL</FieldLabel>
          <FieldContent>
            <Input
              v-model="form.imageUrl"
              placeholder="https://example.com/image.jpg"
              :aria-invalid="hasError('ImageUrl')"
            />
            <FieldError v-if="hasError('ImageUrl')" :errors="[getErrorMessage('ImageUrl')]" />
          </FieldContent>
        </Field>
      </div>
    </div>

    <div class="flex justify-end pt-4">
      <Button type="submit" :disabled="loading">
        {{ loading ? 'Saving...' : 'Save Product' }}
      </Button>
    </div>
  </form>
</template>
