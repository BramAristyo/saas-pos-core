package seeder

import (
	"github.com/BramAristyo/go-pos-mawish/internal/model"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func SeedUserData(db *gorm.DB) {
	users := []*model.User{
		{
			ID:       uuid.New(),
			Name:     "Alice Admin",
			Email:    "alice.admin@example.com",
			Password: "password1",
			Role:     model.RoleAdmin,
			IsActive: true,
		},
		{
			ID:       uuid.New(),
			Name:     "Bob Cashier",
			Email:    "bob.cashier@example.com",
			Password: "password2",
			Role:     model.RoleCashier,
			IsActive: true,
		},
		{
			ID:       uuid.New(),
			Name:     "Charlie Cashier",
			Email:    "charlie.cashier@example.com",
			Password: "password3",
			Role:     model.RoleCashier,
			IsActive: true,
		},
		{
			ID:       uuid.New(),
			Name:     "Diana Admin",
			Email:    "diana.admin@example.com",
			Password: "password4",
			Role:     model.RoleAdmin,
			IsActive: true,
		},
		{
			ID:       uuid.New(),
			Name:     "Eve Cashier",
			Email:    "eve.cashier@example.com",
			Password: "password5",
			Role:     model.RoleCashier,
			IsActive: true,
		},
	}

	for _, user := range users {
		hashed, _ := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
		user.Password = string(hashed)
	}

	db.Clauses(clause.OnConflict{DoNothing: true}).Create(&users)
}
