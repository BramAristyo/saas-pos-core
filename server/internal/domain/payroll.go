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

func (p *Payroll) Calculate(as []Attendance) {
	for _, a := range as {
		p.TotalDeduction.Add(decimal.NewFromFloat(a.DeductionAmount))
	}

	p.NetSalary = p.BaseSalary.Sub(p.TotalDeduction)
}
