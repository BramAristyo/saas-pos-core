package seeder

import (
	"time"
	"github.com/BramAristyo/saas-pos-core/server/internal/domain"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func SeedAttendanceData(db *gorm.DB) {
	date := time.Date(2024, 1, 1, 0, 0, 0, 0, time.Local)
	in1 := time.Date(2024, 1, 1, 8, 0, 0, 0, time.Local)
	out1 := time.Date(2024, 1, 1, 16, 0, 0, 0, time.Local)
	in2 := time.Date(2024, 1, 1, 14, 10, 0, 0, time.Local)
	
	schedule1 := uuid.MustParse("00000000-0000-0000-0000-000000000901")
	schedule2 := uuid.MustParse("00000000-0000-0000-0000-000000000902")

	attendances := []domain.Attendance{
		{
			ID:              uuid.MustParse("00000000-0000-0000-0000-000000001001"),
			EmployeeID:      uuid.MustParse("00000000-0000-0000-0000-000000000801"),
			Date:            date,
			CheckIn:         &in1,
			CheckOut:        &out1,
			ShiftScheduleID: &schedule1,
			LateMinutes:     0,
			DeductionAmount: 0,
		},
		{
			ID:              uuid.MustParse("00000000-0000-0000-0000-000000001002"),
			EmployeeID:      uuid.MustParse("00000000-0000-0000-0000-000000000802"),
			Date:            date,
			CheckIn:         &in2,
			ShiftScheduleID: &schedule2,
			LateMinutes:     10,
			DeductionAmount: 5000,
		},
	}
	db.Clauses(clause.OnConflict{DoNothing: true}).Create(&attendances)
}
