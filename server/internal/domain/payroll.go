package domain

import (
	"time"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type Payroll struct {
	ID             uuid.UUID       `gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	EmployeeID     uuid.UUID       `gorm:"type:uuid;not null"`
	Employee       *Employee       `gorm:"foreignKey:EmployeeID"`
	PeriodStart    time.Time       `gorm:"type:date;not null"`
	PeriodEnd      time.Time       `gorm:"type:date;not null"`
	BaseSalary     decimal.Decimal `gorm:"type:decimal(12,2);not null"`
	TotalDeduction decimal.Decimal `gorm:"type:decimal(12,2);default:0"` // calc by method receiver
	NetSalary      decimal.Decimal `gorm:"type:decimal(12,2);not null"`  // calc by method receiver
	Notes          *string         `gorm:"type:text"`
	CreatedAt      time.Time       `gorm:"autoCreateTime"`
	// UpdatedAt      time.Time       `gorm:"autoUpdateTime"`
	// DeletedAt      gorm.DeletedAt  `gorm:"index"`
}

// Calculate computes the total deduction and net salary based on the provided attendance records.
// It sums up all deductions from the attendances and calculates the net salary as:
// (BaseSalary * number of shifts) - TotalDeduction.
func (p *Payroll) Calculate(as []Attendance) {
	p.TotalDeduction = decimal.Zero

	for _, a := range as {
		p.TotalDeduction = p.TotalDeduction.Add(decimal.NewFromFloat(a.DeductionAmount))
	}

	totalShifts := decimal.NewFromInt(int64(len(as)))

	grossSalary := p.BaseSalary.Mul(totalShifts)

	p.NetSalary = grossSalary.Sub(p.TotalDeduction)
}
