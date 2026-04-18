<script setup lang="ts">
import { ref, watch, nextTick, computed } from 'vue'
import { formatAmount, parseAmount } from '@/utils/currency'
import { cn } from '@/lib/utils'

defineOptions({
  inheritAttrs: false,
})

interface Props {
  modelValue: string | number
  mode?: 'money' | 'percentage' | 'number'
  placeholder?: string
  disabled?: boolean
  class?: string
  prefix?: string
  suffix?: string
  decimals?: number
}

const props = withDefaults(defineProps<Props>(), {
  mode: 'money',
  placeholder: '0',
  decimals: (props) => (props.mode === 'money' ? 0 : 2),
})

const emit = defineEmits<{
  (e: 'update:modelValue', value: string): void
}>()

const inputRef = ref<HTMLInputElement | null>(null)
const displayValue = ref('')

// Computed values based on mode
const effectivePrefix = computed(() => {
  if (props.prefix !== undefined) return props.prefix
  return props.mode === 'money' ? 'Rp' : ''
})

const effectiveSuffix = computed(() => {
  if (props.suffix !== undefined) return props.suffix
  return props.mode === 'percentage' ? '%' : ''
})

// Sync internal display value with external modelValue
watch(
  () => props.modelValue,
  (newVal) => {
    const formatted = formatAmount(newVal, props.decimals)
    if (formatted !== displayValue.value) {
      displayValue.value = formatted
    }
  },
  { immediate: true },
)

function onInput(e: Event) {
  const el = e.target as HTMLInputElement
  const rawValue = el.value

  // Clean it up (allow only digits and decimal separator if decimals > 0)
  // We use regex to extract numeric parts.
  const regex = props.decimals > 0 ? /[^0-9.]/g : /[^0-9]/g
  let cleanValue = rawValue.replace(regex, '')

  // Ensure only one dot
  if (props.decimals > 0) {
    const parts = cleanValue.split('.')
    if (parts.length > 2) {
      cleanValue = parts[0] + '.' + parts.slice(1).join('')
    }
  }

  // Format it for display
  const formattedValue = formatAmount(cleanValue, props.decimals)

  // Update internal and emit raw
  displayValue.value = formattedValue
  emit('update:modelValue', cleanValue)

  // Important: manual sync to the element to handle cursor
  const cursor = el.selectionStart || 0
  const oldLen = rawValue.length

  nextTick(() => {
    if (el) {
      el.value = formattedValue
      const newLen = formattedValue.length
      const diff = newLen - oldLen
      const newPos = Math.max(0, cursor + diff)
      el.setSelectionRange(newPos, newPos)
    }
  })
}

function onBlur() {
  if (!displayValue.value) {
    displayValue.value = props.decimals > 0 ? formatAmount('0', props.decimals) : '0'
    emit('update:modelValue', '0')
  } else {
    // Format to ensure decimal places are consistent
    displayValue.value = formatAmount(props.modelValue, props.decimals)
  }
}
</script>

<template>
  <div class="relative w-full group">
    <div
      v-if="effectivePrefix"
      class="absolute left-3 top-1/2 -translate-y-1/2 text-sm text-muted-foreground font-bold pointer-events-none transition-colors group-focus-within:text-primary z-10"
    >
      {{ effectivePrefix }}
    </div>

    <input
      ref="inputRef"
      type="text"
      inputmode="decimal"
      v-bind="$attrs"
      :value="displayValue"
      :placeholder="placeholder"
      :disabled="disabled"
      :class="
        cn(
          'flex h-9 w-full rounded-md border border-input bg-background px-3 py-1 text-sm shadow-sm transition-all file:border-0 file:bg-transparent file:text-sm file:font-medium placeholder:text-muted-foreground focus-visible:outline-none focus-visible:ring-1 focus-visible:ring-ring disabled:cursor-not-allowed disabled:opacity-50 text-right font-bold',
          'aria-invalid:ring-destructive/20 dark:aria-invalid:ring-destructive/40 aria-invalid:border-destructive',
          effectivePrefix ? 'pl-10' : 'pl-3',
          effectiveSuffix ? 'pr-8' : 'pr-3',
          props.class,
        )
      "
      @input="onInput"
      @blur="onBlur"
    />

    <div
      v-if="effectiveSuffix"
      class="absolute right-3 top-1/2 -translate-y-1/2 text-sm text-muted-foreground font-bold pointer-events-none transition-colors group-focus-within:text-primary z-10"
    >
      {{ effectiveSuffix }}
    </div>
  </div>
</template>
