<script setup lang="ts">
import { useRouter } from 'vue-router'
import AppLayout from '@/layouts/AppLayout.vue'
import ModifierForm from './ModifierForm.vue'

const router = useRouter()

function handleSuccess() {
  router.push({ name: 'modifiers' })
}
</script>

<template>
  <AppLayout>
    <ModifierForm @success="handleSuccess" />
  </AppLayout>
</template>
