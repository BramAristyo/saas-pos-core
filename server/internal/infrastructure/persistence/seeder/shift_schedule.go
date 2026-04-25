package seeder

import (
	"github.com/BramAristyo/saas-pos-core/server/internal/domain"
	"gorm.io/gorm"
)

func SeedShiftScheduleData(db *gorm.DB) {
	schedules := []domain.ShiftSchedule{
		{
			Name:                "Pagi",
			StartTime:           "08:00:00",
			EndTime:             "16:00:00",
			ToleranceMinutes:    15,
			LateIntervalMinutes: 10,
			LateDeductionAmount: 5000,
		},
		{
			Name:                "Sore",
			StartTime:           "14:00:00",
			EndTime:             "22:00:00",
			ToleranceMinutes:    15,
			LateIntervalMinutes: 10,
			LateDeductionAmount: 5000,
		},
	}

	for _, s := range schedules {
		var existing domain.ShiftSchedule
		if err := db.Where("name = ?", s.Name).First(&existing).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				db.Create(&s)
			}
		}
	}
}
