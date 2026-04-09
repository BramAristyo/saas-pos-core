import { watchDebounced } from '@vueuse/core'
import { ref } from 'vue'

export function useSearch(onSearch: () => void, debounce = 500) {
  const search = ref('')

  watchDebounced(search, onSearch, { debounce: debounce })

  return { search }
}
