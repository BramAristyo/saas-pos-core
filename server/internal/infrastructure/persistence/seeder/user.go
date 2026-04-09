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
			Name:     "Camelia Admin",
			Email:    "camelia.admin@cameliawhite.my.id",
			Password: "password",
			Role:     domain.RoleAdmin,
		}, {
			Name:     "White Admin",
			Email:    "white.admin@cameliawhite.my.id",
			Password: "password",
			Role:     domain.RoleAdmin,
		},
		{
			Name:     "John Cashier",
			Email:    "john.cashier@cameliawhite.my.id",
			Password: "password",
			Role:     domain.RoleCashier,
		},
	}

	for _, user := range users {
		hashed, _ := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
		user.Password = string(hashed)
	}

	db.Clauses(clause.OnConflict{DoNothing: true}).Create(&users)
}
