package seeder

import (
	"github.com/BramAristyo/saas-pos-core/server/internal/domain"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func SeedUserData(db *gorm.DB) {
	users := []*domain.User{
		{
			Name:     "Alice Admin",
			Email:    "alice.admin@example.com",
			Password: "password1",
			Role:     domain.RoleAdmin,
		},
		{
			Name:     "Bob Cashier",
			Email:    "bob.cashier@example.com",
			Password: "password2",
			Role:     domain.RoleCashier,
		},
		{
			Name:     "Charlie Cashier",
			Email:    "charlie.cashier@example.com",
			Password: "password3",
			Role:     domain.RoleCashier,
		},
		{
			Name:     "Diana Admin",
			Email:    "diana.admin@example.com",
			Password: "password4",
			Role:     domain.RoleAdmin,
		},
		{
			Name:     "Eve Cashier",
			Email:    "eve.cashier@example.com",
			Password: "password5",
			Role:     domain.RoleCashier,
		},
	}

	for _, user := range users {
		hashed, _ := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
		user.Password = string(hashed)
	}

	db.Clauses(clause.OnConflict{DoNothing: true}).Create(&users)
}
