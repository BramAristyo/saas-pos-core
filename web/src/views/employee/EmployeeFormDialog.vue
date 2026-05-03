<script setup lang="ts">
import { computed, reactive, watch } from 'vue'
import { useEmployeeStore } from '@/stores/employee.store'
import type { Employee, CreateEmployeeRequest, UpdateEmployeeRequest } from '@/types/employee.types'
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
import { Label } from '@/components/ui/label'
import { InputOTP, InputOTPGroup, InputOTPSlot } from '@/components/ui/input-otp'
import { toast } from 'vue-sonner'
import { useFormErrors } from '@/composables/common/useFormErrors'

const props = defineProps<{
  open: boolean
  employee?: Employee | null
}>()

const emit = defineEmits<{
  (e: 'update:open', value: boolean): void
  (e: 'success'): void
}>()

const employeeStore = useEmployeeStore()

const { setErrors, clearErrors, getErrorMessage, hasError } = useFormErrors()

const isOpen = computed({
  get: () => props.open,
  set: (value) => emit('update:open', value),
})

const isEdit = computed(() => !!props.employee)

const form = reactive<CreateEmployeeRequest>({
  name: '',
  phone: '',
  baseSalary: 0,
  pin: '',
})

watch(
  () => props.employee,
  (newEmployee) => {
    clearErrors()
    if (newEmployee) {
      form.name = newEmployee.name
      form.phone = newEmployee.phone
      form.baseSalary = newEmployee.baseSalary
      form.pin = ''
    } else {
      form.name = ''
      form.phone = ''
      form.baseSalary = 0
      form.pin = ''
    }
  },
  { immediate: true },
)

async function handleSubmit() {
  clearErrors()
  try {
    if (isEdit.value && props.employee) {
      await employeeStore.update(props.employee.id, form as UpdateEmployeeRequest)
      toast.success('Employee updated successfully')
    } else {
      await employeeStore.create(form)
      toast.success('Employee created successfully')
    }
    emit('success')
    isOpen.value = false
  } catch (err: any) {
    if (err?.error && Array.isArray(err.error)) {
      setErrors(err.error as ValidationError[])
    } else {
      toast.error(err?.message || 'Failed to save employee')
    }
  }
}
</script>

<template>
  <Dialog v-model:open="isOpen">
    <DialogContent class="sm:max-w-[425px]">
      <form @submit.prevent="handleSubmit">
        <DialogHeader>
          <DialogTitle>{{ isEdit ? 'Edit Employee' : 'Add Employee' }}</DialogTitle>
          <DialogDescription>
            {{
              isEdit
                ? 'Update the details of your employee here.'
                : 'Create a new employee to manage your team.'
            }}
          </DialogDescription>
        </DialogHeader>

        <div class="grid gap-4 py-4">
          <div class="grid gap-2">
            <Label for="name">Name</Label>
            <Input id="name" v-model="form.name" placeholder="Employee name" required />
            <span v-if="hasError('Name')" class="text-xs text-destructive">
              {{ getErrorMessage('Name') }}
            </span>
          </div>

          <div class="grid gap-2">
            <Label for="phone">Phone</Label>
            <Input id="phone" v-model="form.phone" placeholder="Phone number" />
            <span v-if="hasError('Phone')" class="text-xs text-destructive">
              {{ getErrorMessage('Phone') }}
            </span>
          </div>

          <div class="grid gap-2">
            <Label for="baseSalary">Base Salary</Label>
            <Input
              id="baseSalary"
              v-model.number="form.baseSalary"
              type="number"
              placeholder="0"
              required
            />
            <span v-if="hasError('BaseSalary')" class="text-xs text-destructive">
              {{ getErrorMessage('BaseSalary') }}
            </span>
          </div>

          <div class="grid gap-2">
            <Label for="pin">PIN (6 Digits)</Label>
            <InputOTP id="pin" v-model="form.pin" :maxlength="6" class="w-full">
              <InputOTPGroup class="w-full flex">
                <InputOTPSlot :index="0" class="flex-1 h-12" />
                <InputOTPSlot :index="1" class="flex-1 h-12" />
                <InputOTPSlot :index="2" class="flex-1 h-12" />
                <InputOTPSlot :index="3" class="flex-1 h-12" />
                <InputOTPSlot :index="4" class="flex-1 h-12" />
                <InputOTPSlot :index="5" class="flex-1 h-12" />
              </InputOTPGroup>
            </InputOTP>
            <span v-if="hasError('Pin')" class="text-xs text-destructive">
              {{ getErrorMessage('Pin') }}
            </span>
            <p v-if="isEdit" class="text-xs text-muted-foreground italic">
              Leave empty to keep current PIN.
            </p>
          </div>
        </div>

        <DialogFooter>
          <Button type="button" variant="outline" @click="isOpen = false"> Cancel </Button>
          <Button type="submit" :disabled="employeeStore.loading">
            {{ employeeStore.loading ? 'Saving...' : 'Save' }}
          </Button>
        </DialogFooter>
      </form>
    </DialogContent>
  </Dialog>
</template>
