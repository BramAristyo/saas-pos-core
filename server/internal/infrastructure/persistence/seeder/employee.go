package seeder

import (
	"github.com/BramAristyo/saas-pos-core/server/internal/domain"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func SeedEmployeeData(db *gorm.DB) {
	pin1, _ := bcrypt.GenerateFromPassword([]byte("1234"), 10)
	pin2, _ := bcrypt.GenerateFromPassword([]byte("5678"), 10)

	employees := []domain.Employee{
		{
			Code:       "EMP-001",
			Name:       "John Doe",
			BaseSalary: 5000000,
			PinHash:    string(pin1),
		},
		{
			Code:       "EMP-002",
			Name:       "Jane Smith",
			BaseSalary: 4500000,
			PinHash:    string(pin2),
		},
	}

	for _, e := range employees {
		var existing domain.Employee
		if err := db.Where("code = ?", e.Code).First(&existing).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				db.Create(&e)
			}
		}
	}
}
