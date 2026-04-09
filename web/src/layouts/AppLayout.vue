<script setup lang="ts">
import AppSidebar from '@/components/AppSidebar.vue'
import { SidebarInset, SidebarProvider, SidebarTrigger } from '@/components/ui/sidebar'
import { Separator } from '@/components/ui/separator'
import {
  Breadcrumb,
  BreadcrumbItem,
  BreadcrumbLink,
  BreadcrumbList,
  BreadcrumbPage,
  BreadcrumbSeparator,
} from '@/components/ui/breadcrumb'
import { useAuthStore } from '@/stores/auth.stores'
import { computed } from 'vue'
import { useRoute } from 'vue-router'

const authStore = useAuthStore()
const route = useRoute()

const user = computed(() => {
  return {
    name: authStore.user?.name || 'Jane Doe',
    email: authStore.user?.email || 'jane@example.com',
    avatar: 'https://avatar.iran.liara.run/public',
  }
})

const breadcrumbs = computed(() => {
  const matched = route.matched.filter((r) => r.meta?.title)

  return matched.map((record, index) => {
    return {
      title: record.meta.title,
      url: record.path,
      isLast: index === matched.length - 1,
    }
  })
})
</script>

<template>
  <SidebarProvider>
    <AppSidebar :user="user" />
    <SidebarInset>
      <header class="flex h-16 shrink-0 items-center gap-2 border-b px-4">
        <SidebarTrigger class="-ml-1" />
        <Separator orientation="vertical" class="mr-2 h-4" />
        <Breadcrumb>
          <BreadcrumbList>
            <template v-for="breadcrumb in breadcrumbs" :key="breadcrumb.url">
              <BreadcrumbItem>
                <BreadcrumbLink v-if="!breadcrumb.isLast" as-child>
                  <RouterLink :to="breadcrumb.url">
                    {{ breadcrumb.title }}
                  </RouterLink>
                </BreadcrumbLink>
                <BreadcrumbPage v-else>
                  {{ breadcrumb.title }}
                </BreadcrumbPage>
              </BreadcrumbItem>
              <BreadcrumbSeparator v-if="!breadcrumb.isLast" />
            </template>
          </BreadcrumbList>
        </Breadcrumb>
      </header>
      <main class="flex flex-1 flex-col overflow-y-auto">
        <div class="container max-w-7xl mx-auto p-4 lg:p-6">
          <slot />
        </div>
      </main>
    </SidebarInset>
  </SidebarProvider>
</template>
