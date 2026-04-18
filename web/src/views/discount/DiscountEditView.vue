<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { useRouter, useRoute } from 'vue-router'
import { useDiscountStore } from '@/stores/discount.store'
import AppLayout from '@/layouts/AppLayout.vue'
import DiscountForm from './DiscountForm.vue'
import { Loader2 } from 'lucide-vue-next'

const discountStore = useDiscountStore()
const router = useRouter()
const route = useRoute()
const fetching = ref(true)

const id = route.params.id as string

onMounted(async () => {
  fetching.value = true
  try {
    await discountStore.fetchById(id)
  } catch (error: any) {
    router.push({ name: 'discounts' })
  } finally {
    fetching.value = false
  }
})

function handleSuccess() {
  router.push({ name: 'discounts' })
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
      @success="handleSuccess"
    />
  </AppLayout>
</template>
