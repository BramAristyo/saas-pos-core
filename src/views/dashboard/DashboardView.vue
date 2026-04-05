<script setup lang="ts">
import { useAuthStore } from '@/stores/auth.stores'

const authStore = useAuthStore()
</script>

<template>
  <div class="flex min-h-svh items-center justify-center">
    <h1 class="text-2xl font-bold">Hello, {{ authStore.user?.name }}!</h1>
  </div>
</template>
