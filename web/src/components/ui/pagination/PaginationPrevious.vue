<script setup lang="ts">
import type { PaginationPrevProps } from 'reka-ui'
import { reactiveOmit } from '@vueuse/core'
import { ChevronLeft } from 'lucide-vue-next'
import { PaginationPrev, useForwardProps } from 'reka-ui'
import { Button } from '@/components/ui/button'

const props = defineProps<PaginationPrevProps>()

const delegatedProps = reactiveOmit(props)
const forwarded = useForwardProps(delegatedProps)
</script>

<template>
  <PaginationPrev v-bind="forwarded" as-child>
    <Button variant="outline" size="icon" class="h-8 w-8">
      <ChevronLeft class="h-4 w-4" />
    </Button>
  </PaginationPrev>
</template>
