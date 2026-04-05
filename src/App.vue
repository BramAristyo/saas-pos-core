<script setup lang="ts">
import { onMounted } from 'vue'
import { useAuthStore } from '@/stores/auth.stores'
import { Toaster } from '@/components/ui/sonner'
const authStore = useAuthStore()

onMounted(async () => {
  if (authStore.token) {
    await authStore.fetchMe()
  }
})
</script>

<template>
  <Toaster position="top-center" rich-colors />
  <RouterView />
</template>
