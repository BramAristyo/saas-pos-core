package seeder

import (
	"time"

	"github.com/BramAristyo/saas-pos-core/server/internal/domain"
	"gorm.io/gorm"
)

func SeedAttendanceData(db *gorm.DB) {
	var employees []domain.Employee
	db.Find(&employees)

	if len(employees) < 1 {
		return
	}

	var schedules []domain.ShiftSchedule
	db.Find(&schedules)

	if len(schedules) < 1 {
		return
	}

	now := time.Now()
	checkIn := now.Add(-8 * time.Hour)
	checkOut := now

	attendances := []domain.Attendance{
		{
			EmployeeID:      employees[0].ID,
			Date:            now,
			CheckIn:         &checkIn,
			CheckOut:        &checkOut,
			ShiftScheduleID: &schedules[0].ID,
			LateMinutes:     0,
			DeductionAmount: 0,
		},
		{
			EmployeeID:      employees[1].ID,
			Date:            now,
			CheckIn:         &checkIn,
			ShiftScheduleID: &schedules[1].ID,
			LateMinutes:     10,
			DeductionAmount: 5000,
		},
	}

	for _, a := range attendances {
		var existing domain.Attendance
		if err := db.Where("employee_id = ? AND date = ?", a.EmployeeID, a.Date.Format("2006-01-02")).First(&existing).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				db.Create(&a)
			}
		}
	}
}
