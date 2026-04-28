<script setup lang="ts">
import type { SidebarProps } from '@/components/ui/sidebar'
import { Store, ChevronDown, ChevronUp, LayoutDashboard, Package, Receipt } from 'lucide-vue-next'
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
      title: 'Overview',
      url: '#',
      icon: LayoutDashboard,
      items: [
        {
          title: 'Dashboard',
          url: '/dashboard',
        },
      ],
    },
    {
      title: 'Catalog',
      url: '#',
      icon: Package,
      items: [
        {
          title: 'Categories',
          url: '/categories',
        },
        {
          title: 'Modifiers',
          url: '/modifiers',
        },
      ],
    },
    {
      title: 'Transactions',
      url: '#',
      icon: Receipt,
      items: [
        {
          title: 'Taxes',
          url: '/taxes',
        },
        {
          title: 'Sales Types',
          url: '/sales-types',
        },
        {
          title: 'Discounts',
          url: '/discounts',
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
  const activeSection = data.navMain.find((item) =>
    item.items.some((child) => isChildActive(child.url)),
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

function isChildActive(itemUrl: string) {
  // Exact match
  if (route.path === itemUrl) return true

  // Special handling for master data sub-routes
  // e.g. /modifiers/create should active /modifiers
  if (itemUrl !== '/' && itemUrl !== '/dashboard' && route.path.startsWith(itemUrl)) {
    return true
  }

  return false
}

const isParentActive = (item: any) => {
  return item.items?.some((child: any) => isChildActive(child.url))
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
                class="flex aspect-square size-8 items-center justify-center rounded-lg bg-sidebar-primary text-sidebar-primary-foreground">
                <Store class="size-4" />
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
          <Collapsible v-for="item in filteredNavMain" :key="item.title"
            :open="searchQuery ? true : openSectionTitle === item.title"
            @update:open="(val) => handleToggle(item.title, val)" class="group/collapsible">
            <SidebarMenuItem>
              <CollapsibleTrigger as-child>
                <SidebarMenuButton :is-active="isParentActive(item)" :class="{ 'font-bold': isParentActive(item) }">
                  <component :is="item.icon" v-if="item.icon" class="size-4 mr-2" />
                  {{ item.title }}
                  <ChevronDown v-if="!searchQuery" class="ml-auto group-data-[state=open]/collapsible:hidden" />
                  <ChevronUp v-if="!searchQuery" class="ml-auto group-data-[state=closed]/collapsible:hidden" />
                </SidebarMenuButton>
              </CollapsibleTrigger>
              <CollapsibleContent v-if="item.items.length">
                <SidebarMenuSub>
                  <SidebarMenuSubItem v-for="childItem in item.items" :key="childItem.title">
                    <RouterLink :to="childItem.url">
                      <SidebarMenuSubButton :is-active="isChildActive(childItem.url)">
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
