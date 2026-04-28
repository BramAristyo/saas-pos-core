<script setup lang="ts">
import type { PaginationLastProps } from 'reka-ui'
import { reactiveOmit } from '@vueuse/core'
import { ChevronsRight } from 'lucide-vue-next'
import { PaginationLast, useForwardProps } from 'reka-ui'
import { Button } from '@/components/ui/button'

const props = defineProps<PaginationLastProps>()

const delegatedProps = reactiveOmit(props)
const forwarded = useForwardProps(delegatedProps)
</script>

<template>
  <PaginationLast v-bind="forwarded" as-child>
    <Button variant="outline" size="icon" class="h-8 w-8">
      <ChevronsRight class="h-4 w-4" />
    </Button>
  </PaginationLast>
</template>
