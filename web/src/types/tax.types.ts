export interface Tax {
  id: string
  name: string
  percentage: string
}

export type CreateTaxRequest = Omit<Tax, 'id'>
export type UpdateTaxRequest = Omit<Tax, 'id'>
