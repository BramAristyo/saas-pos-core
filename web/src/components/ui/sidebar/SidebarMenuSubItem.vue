<script setup lang="ts">
import type { HTMLAttributes } from 'vue'
import { cn } from '@/lib/utils'

const props = defineProps<{
  class?: HTMLAttributes['class']
}>()
</script>

<template>
  <li
    data-slot="sidebar-menu-sub-item"
    data-sidebar="menu-sub-item"
    :class="
      cn(
        'group/menu-sub-item relative before:absolute before:left-[-10.5px] before:top-1/2 before:h-px before:w-2.5',
        props.class,
      )
    "
  >
    <slot />
  </li>
</template>
