<script setup lang="ts">
import { ref, watch, nextTick } from 'vue'
import { formatCurrency, parseCurrency } from '@/utils/currency'
import { cn } from '@/lib/utils'

interface Props {
  modelValue: string | number
  placeholder?: string
  disabled?: boolean
  class?: string
  prefix?: string
}

const props = withDefaults(defineProps<Props>(), {
  placeholder: '0',
  prefix: 'Rp',
})

const emit = defineEmits<{
  (e: 'update:modelValue', value: string): void
}>()

const inputRef = ref<HTMLInputElement | null>(null)
const displayValue = ref('')

// Watch for external changes
watch(
  () => props.modelValue,
  (newVal) => {
    const formatted = formatCurrency(newVal)
    if (formatted !== displayValue.value) {
      displayValue.value = formatted
    }
  },
  { immediate: true },
)

function onInput(e: Event) {
  const el = e.target as HTMLInputElement
  const rawValue = el.value

  // 1. Get numeric string (e.g., "12000000")
  const numericString = parseCurrency(rawValue)

  // 2. Get formatted string (e.g., "12,000,000")
  const formattedValue = formatCurrency(numericString)

  // 3. Save cursor state before updating DOM
  const selectionStart = el.selectionStart || 0

  // 4. Update internal state
  displayValue.value = formattedValue
  emit('update:modelValue', numericString)

  // 5. Manually update input value and restore cursor
  // This is more reliable than relying solely on Vue's :value
  nextTick(() => {
    if (el) {
      const oldLen = rawValue.length
      const newLen = formattedValue.length
      el.value = formattedValue

      const newPosition = selectionStart + (newLen - oldLen)
      el.setSelectionRange(newPosition, newPosition)
    }
  })
}

function onBlur() {
  if (!displayValue.value) {
    displayValue.value = '0'
    emit('update:modelValue', '0')
  }
}
</script>

<template>
  <div class="relative w-full group">
    <div
      class="absolute left-3 top-1/2 -translate-y-1/2 text-sm text-muted-foreground font-bold pointer-events-none transition-colors group-focus-within:text-primary z-10"
    >
      {{ prefix }}
    </div>
    <input
      ref="inputRef"
      type="text"
      inputmode="numeric"
      :value="displayValue"
      :placeholder="placeholder"
      :disabled="disabled"
      :class="
        cn(
          'flex h-9 w-full rounded-md border border-input bg-background px-3 py-1 text-base shadow-sms transition-colors file:border-0 file:bg-transparent file:text-sm file:font-medium placeholder:text-muted-foreground focus-visible:outline-none focus-visible:ring-1 focus-visible:ring-ring disabled:cursor-not-allowed disabled:opacity-50 pl-10 text-right ',
          props.class,
        )
      "
      @input="onInput"
      @blur="onBlur"
    />
  </div>
</template>
