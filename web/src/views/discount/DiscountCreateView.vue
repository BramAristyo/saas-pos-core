<script setup lang="ts">
import { useRouter } from 'vue-router'
import AppLayout from '@/layouts/AppLayout.vue'
import DiscountForm from './DiscountForm.vue'

const router = useRouter()

function handleSuccess() {
  router.push({ name: 'discounts' })
}
</script>

<template>
  <AppLayout>
    <DiscountForm @success="handleSuccess" />
  </AppLayout>
</template>
