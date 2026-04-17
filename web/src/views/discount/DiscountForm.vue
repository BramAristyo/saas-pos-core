<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { useRouter } from 'vue-router'
import { Button } from '@/components/ui/button'
import { Input } from '@/components/ui/input'
import { Label } from '@/components/ui/label'
import { Card, CardContent, CardHeader, CardTitle } from '@/components/ui/card'
import { X } from 'lucide-vue-next'
import type { CreateDiscountRequest, Discount, UpdateDiscountRequest } from '@/types/discount.types'
import { CancelModal } from '@/components/common/cancel'
import { AmountInput } from '@/components/common/form/input/amount'
import { Toggle } from '@/components/common/form'
import { toast } from 'vue-sonner'

const props = defineProps<{
  initialData?: Discount | null
  loading?: boolean
}>()

const emit = defineEmits<{
  (e: 'submit', data: CreateDiscountRequest | UpdateDiscountRequest): void
}>()

const router = useRouter()

const formatForInputDate = (value: string | null | undefined): string => {
  if (!value) return ''
  const date = new Date(value)
  if (isNaN(date.getTime())) return ''
  return date.toISOString().split('T')[0]
}

const name = ref(props.initialData?.name || '')
const type = ref<'fixed' | 'percentage'>(props.initialData?.type || 'fixed')
const value = ref(props.initialData?.value || '0')
const startDate = ref(formatForInputDate(props.initialData?.startDate))
const endDate = ref(formatForInputDate(props.initialData?.endDate))

const isCancelModalOpen = ref(false)

function reset() {
  name.value = props.initialData?.name || ''
  type.value = props.initialData?.type || 'fixed'
  value.value = props.initialData?.value || '0'
  startDate.value = formatForInputDate(props.initialData?.startDate)
  endDate.value = formatForInputDate(props.initialData?.endDate)
}

defineExpose({ reset })

function handleSubmit() {
  if (!name.value) {
    toast.error('Name is required')
    return
  }

  const payload: CreateDiscountRequest = {
    name: name.value,
    type: type.value,
    value: value.value.toString(),
    startDate: startDate.value || null,
    endDate: endDate.value || null,
  }

  emit('submit', payload)
}
</script>

<template>
  <div class="space-y-6 max-w-2xl mx-auto pb-10">
    <div class="flex items-center justify-between">
      <h2 class="text-xl font-semibold">{{ initialData ? 'Edit' : 'Create' }} Discount</h2>
      <div class="flex gap-2">
        <Button variant="outline" @click="isCancelModalOpen = true">
          <X class="size-4 mr-2" />
          Cancel
        </Button>
        <Button :disabled="loading" @click="handleSubmit">
          {{ loading ? 'Saving...' : 'Save Discount' }}
        </Button>
      </div>
    </div>

    <Card>
      <CardHeader>
        <CardTitle>General Information</CardTitle>
      </CardHeader>
      <CardContent class="space-y-4">
        <div class="space-y-2">
          <Label for="name">Discount Name</Label>
          <Input id="name" v-model="name" placeholder="e.g. Member Discount, Seasonal Sale" />
        </div>

        <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
          <Toggle
            v-model="type"
            label="Adjustment Type"
            :options="[
              { label: 'Fixed', value: 'fixed' },
              { label: 'Percentage', value: 'percentage' },
            ]"
          />

          <div class="space-y-2">
            <Label for="value">Value</Label>
            <AmountInput v-model="value" :mode="type === 'fixed' ? 'money' : 'percentage'" />
          </div>
        </div>

        <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
          <div class="space-y-2">
            <Label for="startDate">Start Date</Label>
            <Input id="startDate" v-model="startDate" type="date" />
          </div>
          <div class="space-y-2">
            <Label for="endDate">End Date</Label>
            <Input id="endDate" v-model="endDate" type="date" />
          </div>
        </div>
      </CardContent>
    </Card>

    <CancelModal v-model:open="isCancelModalOpen" @confirm="router.back()" />
  </div>
</template>
