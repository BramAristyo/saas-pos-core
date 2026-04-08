<script setup lang="ts">
import type { SidebarProps } from '@/components/ui/sidebar'
import { LucideFlower, Minus, Plus } from 'lucide-vue-next'
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
      ],
    },
  ],
}
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
      <SearchForm />
    </SidebarHeader>
    <SidebarContent>
      <SidebarGroup>
        <SidebarMenu>
          <Collapsible
            v-for="(item, index) in data.navMain"
            :key="item.title"
            :default-open="index === 1"
            class="group/collapsible"
          >
            <SidebarMenuItem>
              <CollapsibleTrigger as-child>
                <SidebarMenuButton>
                  {{ item.title }}
                  <Plus class="ml-auto group-data-[state=open]/collapsible:hidden" />
                  <Minus class="ml-auto group-data-[state=closed]/collapsible:hidden" />
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
