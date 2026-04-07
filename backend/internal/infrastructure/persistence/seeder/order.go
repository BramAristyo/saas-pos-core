package seeder

import (
	"fmt"
	"time"

	"github.com/BramAristyo/saas-pos-core/backend/internal/domain"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

func SeedOrderData(db *gorm.DB) {
	var shifts []domain.Shift
	db.Limit(1).Find(&shifts)
	if len(shifts) == 0 {
		return
	}
	shift := shifts[0]

	var users []domain.User
	db.Limit(1).Find(&users)
	if len(users) == 0 {
		return
	}
	user := users[0]

	var salesTypes []domain.SalesType
	db.Limit(1).Find(&salesTypes)
	if len(salesTypes) == 0 {
		return
	}
	salesType := salesTypes[0]

	var products []domain.Product
	db.Limit(2).Find(&products)
	if len(products) == 0 {
		return
	}

	var taxes []domain.Tax
	db.Limit(1).Find(&taxes)
	var taxID *string
	if len(taxes) > 0 {
		tID := taxes[0].ID.String()
		taxID = &tID
	}

	var discounts []domain.Discount
	db.Limit(1).Find(&discounts)
	var discountID *string
	if len(discounts) > 0 {
		dID := discounts[0].ID.String()
		discountID = &dID
	}

	dateStr := time.Now().Format("20060102")
	orderNumber := fmt.Sprintf("MW/%s/00001", dateStr)

	order := domain.Order{
		ShiftID:     shift.ID,
		CashierID:   user.ID,
		SalesTypeID: salesType.ID,
		OrderNumber: orderNumber,
		Status:      domain.OrderCompleted,
		Items: []domain.OrderItem{
			{
				ProductID:    &products[0].ID,
				ProductName:  products[0].Name,
				ProductPrice: products[0].Price,
				ProductCogs:  products[0].Cogs,
				Quantity:     2,
			},
		},
		Payments: []domain.Payment{
			{
				Method: domain.Cash,
				Amount: products[0].Price.Mul(decimal.NewFromInt(2)),
			},
		},
	}

	if taxID != nil {
		order.TaxID = &taxes[0].ID
		order.Tax = &taxes[0]
	}
	if discountID != nil {
		order.DiscountID = &discounts[0].ID
		order.Discount = &discounts[0]
	}

	order.CalculateAll()
	db.Create(&order)

	// Create a voided order
	orderNumber2 := fmt.Sprintf("MW/%s/00002", dateStr)
	voidReason := "Wrong items"
	voidedAt := time.Now()
	order2 := domain.Order{
		ShiftID:     shift.ID,
		CashierID:   user.ID,
		SalesTypeID: salesType.ID,
		OrderNumber: orderNumber2,
		Status:      domain.OrderVoided,
		VoidReason:  &voidReason,
		VoidedBy:    &user.ID,
		VoidedAt:    &voidedAt,
		Items: []domain.OrderItem{
			{
				ProductID:    &products[0].ID,
				ProductName:  products[0].Name,
				ProductPrice: products[0].Price,
				ProductCogs:  products[0].Cogs,
				Quantity:     1,
			},
		},
		Payments: []domain.Payment{
			{
				Method: domain.Cash,
				Amount: products[0].Price,
			},
		},
	}
	order2.CalculateAll()
	db.Create(&order2)
}
