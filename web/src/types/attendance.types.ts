export interface Attendance {
  id: string
  date: string
  employeeCode: string
  employeeName: string
  shiftScheduleName: string
  checkIn: string
  checkOut?: string
  lateMinutes: number
  totalWorkMinutes: string | number
  deductionAmount: string | number
  notes?: string
}

export interface CreateAttendanceRequest {
  employeeId: string
  date: string
  checkIn: string
  shiftScheduleId: string
  notes?: string
}
