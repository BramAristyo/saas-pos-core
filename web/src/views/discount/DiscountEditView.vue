<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { useRouter, useRoute } from 'vue-router'
import { useDiscountStore } from '@/stores/discount.store'
import AppLayout from '@/layouts/AppLayout.vue'
import DiscountForm from './DiscountForm.vue'
import type { UpdateDiscountRequest } from '@/types/discount.types'
import { toast } from 'vue-sonner'
import { Loader2 } from 'lucide-vue-next'

const discountStore = useDiscountStore()
const router = useRouter()
const route = useRoute()
const loading = ref(false)
const fetching = ref(true)

const id = route.params.id as string

onMounted(async () => {
  fetching.value = true
  try {
    await discountStore.fetchById(id)
  } catch (error: any) {
    toast.error('Failed to fetch discount details')
    router.push({ name: 'discounts' })
  } finally {
    fetching.value = false
  }
})

async function handleSubmit(data: any) {
  loading.value = true
  try {
    await discountStore.update(id, data as UpdateDiscountRequest)
    toast.success('Discount updated successfully')
    router.push({ name: 'discounts' })
  } catch (error: any) {
    toast.error(error?.message || 'Failed to update discount')
  } finally {
    loading.value = false
  }
}
</script>

<template>
  <AppLayout>
    <div v-if="fetching" class="flex flex-col items-center justify-center min-h-[400px]">
      <Loader2 class="size-8 animate-spin text-primary mb-2" />
      <p class="text-muted-foreground">Loading discount...</p>
    </div>
    <DiscountForm
      v-else-if="discountStore.selected"
      :initial-data="discountStore.selected"
      :loading="loading"
      @submit="handleSubmit"
    />
  </AppLayout>
</template>
