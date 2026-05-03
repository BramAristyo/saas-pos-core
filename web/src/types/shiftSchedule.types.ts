export interface ShiftSchedule {
  id: string
  name: string
  startTime: string
  endTime: string
  toleranceMinutes: number
  lateIntervalMinutes: number
  lateDeductionAmount: number
  createdAt: string
  updatedAt: string
  deletedAt?: string
}

export interface CreateShiftScheduleRequest {
  name: string
  startTime: string
  endTime: string
  toleranceMinutes: number
  lateIntervalMinutes: number
  lateDeductionAmount: number
}

export interface UpdateShiftScheduleRequest {
  name: string
  startTime: string
  endTime: string
  toleranceMinutes: number
  lateIntervalMinutes: number
  lateDeductionAmount: number
}
