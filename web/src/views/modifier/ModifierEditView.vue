<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { useRouter, useRoute } from 'vue-router'
import { useModifierStore } from '@/stores/modifier.store'
import AppLayout from '@/layouts/AppLayout.vue'
import ModifierForm from './ModifierForm.vue'
import { Loader2 } from 'lucide-vue-next'

const modifierStore = useModifierStore()
const router = useRouter()
const route = useRoute()
const fetching = ref(true)

const id = route.params.id as string

onMounted(async () => {
  fetching.value = true
  try {
    await modifierStore.fetchById(id)
  } catch (error: any) {
    router.push({ name: 'modifiers' })
  } finally {
    fetching.value = false
  }
})

function handleSuccess() {
  router.push({ name: 'modifiers' })
}
</script>

<template>
  <AppLayout>
    <div v-if="fetching" class="flex flex-col items-center justify-center min-h-100">
      <Loader2 class="size-8 animate-spin text-primary mb-2" />
      <p class="text-muted-foreground">Loading modifier group...</p>
    </div>
    <ModifierForm
      v-else-if="modifierStore.selected"
      :initial-data="modifierStore.selected"
      @success="handleSuccess"
    />
  </AppLayout>
</template>
/template>
