<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { useRouter, useRoute } from 'vue-router'
import { useModifierStore } from '@/stores/modifier.store'
import AppLayout from '@/layouts/AppLayout.vue'
import ModifierForm from './ModifierForm.vue'
import type { UpdateModifierGroupRequest } from '@/types/modifier.types'
import { toast } from 'vue-sonner'
import { Loader2 } from 'lucide-vue-next'

const modifierStore = useModifierStore()
const router = useRouter()
const route = useRoute()
const loading = ref(false)
const fetching = ref(true)

const id = route.params.id as string

onMounted(async () => {
  fetching.value = true
  try {
    await modifierStore.fetchById(id)
  } catch (error: any) {
    toast.error('Failed to fetch modifier group details')
    router.push({ name: 'modifiers' })
  } finally {
    fetching.value = false
  }
})

async function handleSubmit(data: any) {
  loading.value = true
  try {
    await modifierStore.update(id, data as UpdateModifierGroupRequest)
    toast.success('Modifier group updated successfully')
    router.push({ name: 'modifiers' })
  } catch (error: any) {
    toast.error(error?.message || 'Failed to update modifier group')
  } finally {
    loading.value = false
  }
}
</script>

<template>
  <AppLayout>
    <div v-if="fetching" class="flex flex-col items-center justify-center min-h-[400px]">
      <Loader2 class="size-8 animate-spin text-primary mb-2" />
      <p class="text-muted-foreground">Loading modifier group...</p>
    </div>
    <ModifierForm
      v-else-if="modifierStore.selected"
      :initial-data="modifierStore.selected"
      :loading="loading"
      @submit="handleSubmit"
    />
  </AppLayout>
</template>
