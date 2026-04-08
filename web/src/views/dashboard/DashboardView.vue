<script setup lang="ts">
import { useAuthStore } from '@/stores/auth.stores'
import AppLayout from '@/layouts/AppLayout.vue'
const authStore = useAuthStore()
</script>

<template>
  <AppLayout>
    <div class="flex flex-col gap-4">
      <h1 class="text-2xl font-bold">Hello, {{ authStore.user?.name }}!</h1>
    </div>
  </AppLayout>
</template>
