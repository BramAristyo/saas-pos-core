<script setup lang="ts">
import { computed, reactive, watch } from 'vue'
import { useCoaStore } from '@/stores/coa.store'
import type { Coa, CreateCoaRequest, UpdateCoaRequest } from '@/types/coa.types'
import type { ValidationError } from '@/types/common.types'

import {
  Dialog,
  DialogContent,
  DialogDescription,
  DialogFooter,
  DialogHeader,
  DialogTitle,
} from '@/components/ui/dialog'
import { Button } from '@/components/ui/button'
import { Input } from '@/components/ui/input'
import { Field, FieldContent, FieldLabel, FieldError } from '@/components/ui/field'
import { Toggle } from '@/components/common/form'
import { toast } from 'vue-sonner'
import { useFormErrors } from '@/composables/common/useFormErrors'

const props = defineProps<{
  open: boolean
  coa?: Coa | null
}>()

const emit = defineEmits<{
  (e: 'update:open', value: boolean): void
  (e: 'success'): void
}>()

const coaStore = useCoaStore()

const { setErrors, clearErrors, getErrorMessage, hasError } = useFormErrors()

const isOpen = computed({
  get: () => props.open,
  set: (value) => emit('update:open', value),
})

const isEdit = computed(() => !!props.coa)

const form = reactive<CreateCoaRequest>({
  name: '',
  type: 'in',
  isOperational: false,
})

const typeOptions = [
  { label: 'Income', value: 'in' },
  { label: 'Expenses', value: 'out' },
]

const operationalOptions = [
  { label: 'Operational', value: true },
  { label: 'Non-Operational', value: false },
]

watch(
  () => props.coa,
  (newCoa) => {
    clearErrors()
    if (newCoa) {
      form.name = newCoa.name
      form.type = newCoa.type
      form.isOperational = newCoa.IsOperational
    } else {
      form.name = ''
      form.type = 'in'
      form.isOperational = false
    }
  },
  { immediate: true },
)

async function handleSubmit() {
  clearErrors()
  try {
    if (isEdit.value && props.coa) {
      await coaStore.update(props.coa.id, form as UpdateCoaRequest)
      toast.success('Account updated successfully')
    } else {
      await coaStore.create(form)
      toast.success('Account created successfully')
    }
    emit('success')
    isOpen.value = false
  } catch (err: any) {
    if (err?.error && Array.isArray(err.error)) {
      setErrors(err.error as ValidationError[])
    } else {
      toast.error(err?.message || 'Failed to save account')
    }
  }
}
</script>

<template>
  <Dialog v-model:open="isOpen">
    <DialogContent class="sm:max-w-106.25">
      <form @submit.prevent="handleSubmit">
        <DialogHeader>
          <DialogTitle>{{ isEdit ? 'Edit Account' : 'Add Account' }}</DialogTitle>
          <DialogDescription>
            {{
              isEdit
                ? 'Update the details of your account here.'
                : 'Create a new account for your finances.'
            }}
          </DialogDescription>
        </DialogHeader>

        <div class="grid gap-4 py-4">
          <Field>
            <FieldLabel>Name</FieldLabel>
            <FieldContent>
              <Input
                v-model="form.name"
                placeholder="Account name"
                required
                :aria-invalid="hasError('Name')"
              />
              <FieldError v-if="hasError('Name')" :errors="[getErrorMessage('Name')]" />
            </FieldContent>
          </Field>

          <Field>
            <FieldLabel>Type</FieldLabel>
            <FieldContent>
              <Toggle
                v-model="form.type"
                :options="typeOptions"
              />
              <FieldError v-if="hasError('Type')" :errors="[getErrorMessage('Type')]" />
            </FieldContent>
          </Field>

          <Field>
            <FieldLabel>Operational Status</FieldLabel>
            <FieldContent>
              <Toggle
                v-model="form.isOperational"
                :options="operationalOptions"
              />
              <FieldError v-if="hasError('IsOperational')" :errors="[getErrorMessage('IsOperational')]" />
            </FieldContent>
          </Field>
        </div>

        <DialogFooter>
          <Button type="button" variant="outline" @click="isOpen = false"> Cancel </Button>
          <Button type="submit" :disabled="coaStore.loading">
            {{ coaStore.loading ? 'Saving...' : 'Save' }}
          </Button>
        </DialogFooter>
      </form>
    </DialogContent>
  </Dialog>
</template>
