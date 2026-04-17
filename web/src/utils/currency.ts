/**
 * Formats a number or string into a formatted string with separators
 * e.g. 12000000 -> 12,000,000
 */
export function formatAmount(value: number | string | undefined | null, decimals = 0): string {
  if (value === undefined || value === null || value === '') return ''
  
  const numericValue = typeof value === 'string' 
    ? parseFloat(value.replace(/[^0-9.-]+/g, '')) 
    : value
  
  if (isNaN(numericValue)) return ''
  
  return new Intl.NumberFormat('en-US', {
    minimumFractionDigits: decimals,
    maximumFractionDigits: decimals,
  }).format(numericValue)
}

/**
 * Formats a number or string into currency format with commas
 * e.g. 12000000 -> 12,000,000
 */
export function formatCurrency(value: number | string | undefined | null): string {
  return formatAmount(value, 0)
}

/**
 * Parses a formatted string back to a numeric string for API
 * e.g. "12,000,000.00" -> "12000000.00"
 */
export function parseAmount(value: string | undefined | null): string {
  if (!value) return ''
  // Keep numbers, dots (for decimals), and minus sign
  // But we usually only want one dot. For now, just remove commas.
  return value.replace(/,/g, '')
}

/**
 * Parses a currency string back to a numeric string for API
 * e.g. "-12,000,000" -> "-12000000"
 */
export function parseCurrency(value: string | undefined | null): string {
  if (!value) return ''
  const isNegative = value.startsWith('-')
  const clean = value.replace(/[^0-9]/g, '')
  return isNegative ? `-${clean}` : clean
}
