<script setup lang="ts">
import AppSidebar from '@/components/AppSidebar.vue'
import { SidebarInset, SidebarProvider, SidebarTrigger } from '@/components/ui/sidebar'
import NavUser from '@/components/NavUser.vue'
import { useAuthStore } from '@/stores/auth.stores'
import { computed } from 'vue'
import { useRoute } from 'vue-router'
import { useFormatter } from '@/composables/common/useFormatter'

const authStore = useAuthStore()
const route = useRoute()
const { formatDateOnly } = useFormatter()

const user = computed(() => {
  return {
    name: authStore.user?.name || 'Jane Doe',
    email: authStore.user?.email || 'jane@example.com',
    avatar: 'https://avatar.iran.liara.run/public',
  }
})

const pageTitle = computed(() => {
  return route.meta?.title || 'Overview'
})

const today = computed(() => {
  return formatDateOnly(new Date())
})
</script>

<template>
  <SidebarProvider>
    <AppSidebar />
    <SidebarInset class="bg-background">
      <!-- Professional Floating Header -->
      <header class="sticky top-4 z-40 mx-6 flex h-16 shrink-0 items-center justify-between rounded-2xl border border-border/40 bg-background/80 backdrop-blur-xl px-6 shadow-sm transition-all duration-300">
        <div class="flex items-center gap-4">
          <SidebarTrigger class="-ml-2 h-10 w-10 rounded-xl hover:bg-accent transition-colors" />
          <div class="h-4 w-px bg-border/60" />
          <h1 class="text-sm font-semibold text-foreground hidden md:block">
            {{ pageTitle }}
          </h1>
        </div>
        <div class="flex items-center gap-6">
          <div class="text-[11px] font-medium uppercase tracking-wider text-muted-foreground/80">
            {{ today }}
          </div>
          <NavUser :user="user" />
        </div>
      </header>
      
      <main class="flex flex-1 flex-col overflow-y-auto pt-8">
        <div class="container max-w-7xl mx-auto p-4 lg:p-6 pb-20">
          <slot />
        </div>
      </main>
    </SidebarInset>
  </SidebarProvider>
</template>
