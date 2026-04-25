package dto

import (
	"fmt"

	"github.com/BramAristyo/saas-pos-core/server/internal/domain"
	"github.com/BramAristyo/saas-pos-core/server/pkg/filter"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type AttendanceResponse struct {
	ID                uuid.UUID       `json:"id"`
	EmployeeCode      string          `json:"employeeCode"`
	EmployeeName      string          `json:"employeeName"`
	Date              string          `json:"date"`
	CheckIn           string          `json:"checkIn"`
	CheckOut          *string         `json:"checkOut"`
	TotalWorkMinutes  *string         `json:"totalWorkMinutes"`
	ShiftScheduleName string          `json:"shiftScheduleName"`
	Notes             *string         `json:"notes"`
	LateMinutes       int             `json:"lateMinutes"`
	DeductionAmount   decimal.Decimal `json:"deductionAmount"`
}

type AttendanceResponsePagination struct {
	Data []AttendanceResponse `json:"data"`
	Meta filter.Meta          `json:"meta"`
}

func ToAttendanceResponses(as []domain.Attendance) []AttendanceResponse {
	res := make([]AttendanceResponse, 0, len(as))

	for _, a := range as {
		var empCode, empName string
		if a.Employee != nil {
			empCode = a.Employee.Code
			empName = a.Employee.Name
		}

		dateStr := a.Date.Format("2006-01-02")

		var checkInStr string
		if a.CheckIn != nil {
			checkInStr = a.CheckIn.Format("15:04")
		}

		var checkOutStr *string
		var totalWorkMinsStr *string

		if a.CheckOut != nil {
			co := a.CheckOut.Format("15:04")
			checkOutStr = &co

			if a.CheckIn != nil {
				duration := a.CheckOut.Sub(*a.CheckIn)
				strMins := fmt.Sprintf("%d", int(duration.Minutes()))
				totalWorkMinsStr = &strMins
			}
		}

		var shiftName string
		if a.ShiftSchedule != nil {
			shiftName = a.ShiftSchedule.Name
		}

		res = append(res, AttendanceResponse{
			ID:                a.ID,
			EmployeeCode:      empCode,
			EmployeeName:      empName,
			Date:              dateStr,
			CheckIn:           checkInStr,
			CheckOut:          checkOutStr,
			TotalWorkMinutes:  totalWorkMinsStr,
			ShiftScheduleName: shiftName,
			Notes:             a.Notes,
			LateMinutes:       a.LateMinutes,
			DeductionAmount:   decimal.NewFromFloat(a.DeductionAmount),
		})
	}

	return res
}

func ToAttendanceResponsePagination(a []AttendanceResponse, p filter.PaginationWithInputFilter, totalRows int64) AttendanceResponsePagination {
	return AttendanceResponsePagination{
		Data: a,
		Meta: p.ToMeta(totalRows),
	}
}
