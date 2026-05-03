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

  const formatDateOnly = (value: string | Date | null | undefined): string => {
    if (!value) return '-'

    const date = new Date(value)
    if (isNaN(date.getTime())) return '-'

    return new Intl.DateTimeFormat('en-US', {
      month: 'short',
      day: 'numeric',
      year: 'numeric',
    }).format(date)
  }

  const formatDuration = (minutes: number | string | undefined | null): string => {
    if (minutes === undefined || minutes === null) return '-'
    const totalMinutes = typeof minutes === 'string' ? parseInt(minutes) : minutes
    if (isNaN(totalMinutes)) return '-'

    const h = Math.floor(totalMinutes / 60)
    const m = totalMinutes % 60

    if (h === 0) return `${m}m`
    if (m === 0) return `${h}h`
    return `${h}h ${m}m`
  }

  return {
    formatRupiah,
    formatPercent,
    formatDate,
    formatDateOnly,
    formatDuration,
  }
}
