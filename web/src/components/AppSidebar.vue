<script setup lang="ts">
import type { SidebarProps } from '@/components/ui/sidebar'
import { LucideFlower, ChevronDown, ChevronUp } from 'lucide-vue-next'
import SearchForm from '@/components/SearchForm.vue'
import { Collapsible, CollapsibleContent, CollapsibleTrigger } from '@/components/ui/collapsible'
import NavUser from '@/components/NavUser.vue'
import {
  Sidebar,
  SidebarContent,
  SidebarFooter,
  SidebarGroup,
  SidebarHeader,
  SidebarMenu,
  SidebarMenuButton,
  SidebarMenuItem,
  SidebarMenuSub,
  SidebarMenuSubButton,
  SidebarMenuSubItem,
  SidebarRail,
} from '@/components/ui/sidebar'
import { computed, ref, watch } from 'vue'
import { useRoute } from 'vue-router'

interface Props extends SidebarProps {
  user: {
    name: string
    email: string
    avatar: string
  }
}

const props = defineProps<Props>()

const route = useRoute()
const data = {
  navMain: [
    {
      title: 'Getting Started',
      url: '#',
      items: [
        {
          title: 'Dashboard',
          url: '/dashboard',
        },
      ],
    },
    {
      title: 'Master Data',
      url: '#',
      items: [
        {
          title: 'Category',
          url: '/categories',
        },
        {
          title: 'Tax',
          url: '/taxes',
        },
        {
          title: 'Sales Type',
          url: '/sales-types',
        },
      ],
    },
  ],
}

const searchQuery = ref('')
const openSectionTitle = ref<string | null>(null)

const filteredNavMain = computed(() => {
  const query = searchQuery.value.toLowerCase().trim()
  if (!query) return data.navMain

  return data.navMain
    .map((section) => ({
      ...section,
      items: section.items.filter((item) => item.title.toLowerCase().includes(query)),
    }))
    .filter((section) => section.items.length > 0)
})

function updateOpenSection() {
  const currentPath = route.path
  const activeSection = data.navMain.find((item) =>
    item.items.some((child) => child.url === currentPath),
  )
  openSectionTitle.value = activeSection ? activeSection.title : null
}

function handleToggle(title: string, open: boolean) {
  // Only allow manual toggle if search is empty
  if (searchQuery.value) return

  if (open) {
    openSectionTitle.value = title
  } else if (openSectionTitle.value === title) {
    openSectionTitle.value = null
  }
}

watch(
  () => route.path,
  () => {
    updateOpenSection()
  },
  { immediate: true },
)
</script>

<template>
  <Sidebar v-bind="props">
    <SidebarHeader>
      <SidebarMenu>
        <SidebarMenuItem>
          <SidebarMenuButton size="lg" as-child>
            <a href="#">
              <div
                class="flex aspect-square size-8 items-center justify-center rounded-lg bg-sidebar-primary text-sidebar-primary-foreground"
              >
                <LucideFlower class="size-4" />
              </div>
              <div class="flex flex-col gap-0.5 leading-none">
                <span class="font-medium">Point of Sales</span>
                <span class="">v1.0.0</span>
              </div>
            </a>
          </SidebarMenuButton>
        </SidebarMenuItem>
      </SidebarMenu>
      <SearchForm v-model="searchQuery" />
    </SidebarHeader>
    <SidebarContent>
      <SidebarGroup>
        <SidebarMenu>
          <Collapsible
            v-for="item in filteredNavMain"
            :key="item.title"
            :open="searchQuery ? true : openSectionTitle === item.title"
            @update:open="(val) => handleToggle(item.title, val)"
            class="group/collapsible"
          >
            <SidebarMenuItem>
              <CollapsibleTrigger as-child>
                <SidebarMenuButton>
                  {{ item.title }}
                  <ChevronDown
                    v-if="!searchQuery"
                    class="ml-auto group-data-[state=open]/collapsible:hidden"
                  />
                  <ChevronUp
                    v-if="!searchQuery"
                    class="ml-auto group-data-[state=closed]/collapsible:hidden"
                  />
                </SidebarMenuButton>
              </CollapsibleTrigger>
              <CollapsibleContent v-if="item.items.length">
                <SidebarMenuSub>
                  <SidebarMenuSubItem v-for="childItem in item.items" :key="childItem.title">
                    <RouterLink :to="childItem.url" v-slot="{ isActive }">
                      <SidebarMenuSubButton :is-active="isActive">
                        {{ childItem.title }}
                      </SidebarMenuSubButton>
                    </RouterLink>
                  </SidebarMenuSubItem>
                </SidebarMenuSub>
              </CollapsibleContent>
            </SidebarMenuItem>
          </Collapsible>
        </SidebarMenu>
      </SidebarGroup>
    </SidebarContent>
    <SidebarFooter>
      <NavUser :user="user" />
    </SidebarFooter>
    <SidebarRail />
  </Sidebar>
</template>
