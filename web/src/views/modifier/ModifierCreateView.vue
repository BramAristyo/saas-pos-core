<script setup lang="ts">
import { ref } from 'vue'
import { useRouter } from 'vue-router'
import { useModifierStore } from '@/stores/modifier.store'
import AppLayout from '@/layouts/AppLayout.vue'
import ModifierForm from './ModifierForm.vue'
import type { CreateModifierGroupRequest } from '@/types/modifier.types'
import { toast } from 'vue-sonner'

const modifierStore = useModifierStore()
const router = useRouter()
const loading = ref(false)

async function handleSubmit(data: any) {
  loading.value = true
  try {
    await modifierStore.create(data as CreateModifierGroupRequest)
    toast.success('Modifier group created successfully')
    router.push({ name: 'modifiers' })
  } catch (error: any) {
    toast.error(error?.message || 'Failed to create modifier group')
  } finally {
    loading.value = false
  }
}
</script>

<template>
  <AppLayout>
    <ModifierForm :loading="loading" @submit="handleSubmit" />
  </AppLayout>
</template>
