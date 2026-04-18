import type { ValidationError } from '@/types/common.types'
import { ref } from 'vue'

export function useFormErrors() {
  const errors = ref<ValidationError[]>([])

  const setErrors = (newErrors: ValidationError[]) => {
    errors.value = newErrors
  }

  const clearErrors = () => {
    errors.value = []
  }

  const getErrorMessage = (property: string) => {
    if (!property) return ''

    const error = errors.value.find(
      (e) => e.property && e.property.toLowerCase() === property.toLowerCase(),
    )

    if (!error) return ''
    return error.message || `${error.property} failed on ${error.tag}`
  }

  const hasError = (property: string) => {
    if (!property) return false

    return errors.value.some(
      (e) => e.property && e.property.toLowerCase() === property.toLowerCase(),
    )
  }

  return {
    errors,
    setErrors,
    clearErrors,
    getErrorMessage,
    hasError,
  }
}
