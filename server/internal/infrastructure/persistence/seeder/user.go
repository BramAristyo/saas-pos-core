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
			Name:     "Camelia",
			Email:    "camelia@cameliawhite.my.id",
			Password: "password",
			Role:     domain.RoleAdmin,
		}, {
			Name:     "White",
			Email:    "white@cameliawhite.my.id",
			Password: "password",
			Role:     domain.RoleAdmin,
		},
		{
			Name:     "John",
			Email:    "john@cameliawhite.my.id",
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
