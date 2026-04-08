<script setup lang="ts">
import AppSidebar from '@/components/AppSidebar.vue'
import { SidebarInset, SidebarProvider, SidebarTrigger } from '@/components/ui/sidebar'
import { Separator } from '@/components/ui/separator'
import { useAuthStore } from '@/stores/auth.stores'
import { computed } from 'vue'

const authStore = useAuthStore()

const user = computed(() => {
  return {
    name: authStore.user?.name || 'Jane Doe',
    email: authStore.user?.email || 'jane@example.com',
    avatar: 'https://avatar.iran.liara.run/public'
  }
})
</script>

<template>
  <SidebarProvider>
    <AppSidebar :user="user" />
    <SidebarInset>
      <header class="flex h-16 shrink-0 items-center gap-2 border-b px-4">
        <SidebarTrigger class="-ml-1" />
        <Separator orientation="vertical" class="mr-2 h-4" />
        <div class="flex items-center gap-2">
          <!-- Breadcrumb could go here if needed -->
          <h2 class="text-sm font-semibold">Dashboard</h2>
        </div>
      </header>
      <main class="flex flex-1 flex-col gap-4 p-4">
        <slot />
      </main>
    </SidebarInset>
  </SidebarProvider>
</template>
