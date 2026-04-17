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

  return {
    formatRupiah,
    formatPercent,
  }
}
