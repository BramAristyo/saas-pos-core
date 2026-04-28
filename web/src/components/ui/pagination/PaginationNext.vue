<script setup lang="ts">
import type { PaginationNextProps } from 'reka-ui'
import { reactiveOmit } from '@vueuse/core'
import { ChevronRight } from 'lucide-vue-next'
import { PaginationNext, useForwardProps } from 'reka-ui'
import { Button } from '@/components/ui/button'

const props = defineProps<PaginationNextProps>()

const delegatedProps = reactiveOmit(props)
const forwarded = useForwardProps(delegatedProps)
</script>

<template>
  <PaginationNext v-bind="forwarded" as-child>
    <Button variant="outline" size="icon" class="h-8 w-8">
      <ChevronRight class="h-4 w-4" />
    </Button>
  </PaginationNext>
</template>
