package seeder

import (
	"github.com/BramAristyo/saas-pos-core/server/internal/domain"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func SeedEmployeeData(db *gorm.DB) {
	pin1, _ := bcrypt.GenerateFromPassword([]byte("1234"), 10)
	pin2, _ := bcrypt.GenerateFromPassword([]byte("5678"), 10)

	employees := []domain.Employee{
		{
			ID:         uuid.MustParse("00000000-0000-0000-0000-000000000801"),
			Code:       "EMP-001",
			Name:       "John Doe",
			BaseSalary: 5000000,
			PinHash:    string(pin1),
		},
		{
			ID:         uuid.MustParse("00000000-0000-0000-0000-000000000802"),
			Code:       "EMP-002",
			Name:       "Jane Smith",
			BaseSalary: 4500000,
			PinHash:    string(pin2),
		},
	}
	db.Clauses(clause.OnConflict{DoNothing: true}).Create(&employees)
}
