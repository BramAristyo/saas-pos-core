package seeder

import (
	"github.com/BramAristyo/saas-pos-core/server/internal/domain"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func SeedShiftScheduleData(db *gorm.DB) {
	schedules := []domain.ShiftSchedule{
		{
			ID:                  uuid.MustParse("00000000-0000-0000-0000-000000000901"),
			Name:                "Pagi",
			StartTime:           "08:00:00",
			EndTime:             "16:00:00",
			ToleranceMinutes:    15,
			LateIntervalMinutes: 10,
			LateDeductionAmount: 5000,
		},
		{
			ID:                  uuid.MustParse("00000000-0000-0000-0000-000000000902"),
			Name:                "Sore",
			StartTime:           "14:00:00",
			EndTime:             "22:00:00",
			ToleranceMinutes:    15,
			LateIntervalMinutes: 10,
			LateDeductionAmount: 5000,
		},
	}
	db.Clauses(clause.OnConflict{DoNothing: true}).Create(&schedules)
}
