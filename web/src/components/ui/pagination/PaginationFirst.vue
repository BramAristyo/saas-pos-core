<script setup lang="ts">
import type { PaginationFirstProps } from 'reka-ui'
import { reactiveOmit } from '@vueuse/core'
import { ChevronsLeft } from 'lucide-vue-next'
import { PaginationFirst, useForwardProps } from 'reka-ui'
import { Button } from '@/components/ui/button'

const props = defineProps<PaginationFirstProps>()

const delegatedProps = reactiveOmit(props)
const forwarded = useForwardProps(delegatedProps)
</script>

<template>
  <PaginationFirst v-bind="forwarded" as-child>
    <Button variant="outline" size="icon" class="h-8 w-8">
      <ChevronsLeft class="h-4 w-4" />
    </Button>
  </PaginationFirst>
</template>
