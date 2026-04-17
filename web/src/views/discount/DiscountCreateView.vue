<script setup lang="ts">
import { ref } from 'vue'
import { useRouter } from 'vue-router'
import { useDiscountStore } from '@/stores/discount.store'
import AppLayout from '@/layouts/AppLayout.vue'
import DiscountForm from './DiscountForm.vue'
import type { CreateDiscountRequest } from '@/types/discount.types'
import { toast } from 'vue-sonner'

const discountStore = useDiscountStore()
const router = useRouter()
const loading = ref(false)
const formRef = ref<any>(null)

async function handleSubmit(data: any) {
  loading.value = true
  try {
    await discountStore.create(data as CreateDiscountRequest)
    toast.success('Discount created successfully')
    router.push({ name: 'discounts' })
  } catch (error: any) {
    toast.error(error?.message || 'Failed to create discount')
  } finally {
    loading.value = false
  }
}
</script>

<template>
  <AppLayout>
    <DiscountForm ref="formRef" :loading="loading" @submit="handleSubmit" />
  </AppLayout>
</template>
