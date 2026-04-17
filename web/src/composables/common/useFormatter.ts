export function useFormatter() {
  const formatRupiah = (value: number | string): string => {
    const number = typeof value === 'string' ? parseFloat(value) : value
    return new Intl.NumberFormat('id-ID', {
      style: 'currency',
      currency: 'IDR',
      minimumFractionDigits: 0,
      maximumFractionDigits: 0,
    }).format(number || 0)
  }

  const formatPercent = (value: number | string): string => {
    const number = typeof value === 'string' ? parseFloat(value) : value
    return `${number || 0}%`
  }

  const formatDate = (value: string | Date | null | undefined): string => {
    if (!value) return '-'

    const date = new Date(value)
    if (isNaN(date.getTime())) return '-'

    return new Intl.DateTimeFormat('en-US', {
      month: 'short',
      day: 'numeric',
      year: 'numeric',
      hour: 'numeric',
      minute: '2-digit',
      hour12: true,
    }).format(date)
  }

  return {
    formatRupiah,
    formatPercent,
    formatDate,
  }
}
