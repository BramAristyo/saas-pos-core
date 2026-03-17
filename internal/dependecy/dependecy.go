package dependecy

import (
	"github.com/BramAristyo/go-pos-mawish/internal/handler"
	"github.com/BramAristyo/go-pos-mawish/internal/repository"
	"github.com/BramAristyo/go-pos-mawish/internal/service"
	"gorm.io/gorm"
)

type Handlers struct {
	User     *handler.UserHandler
	Category *handler.CategoryHandler
}

func Bootstrap(db *gorm.DB) *Handlers {
	userRepository := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepository)

	categoryRepository := repository.NewCategoryRepository(db)
	categoryService := service.NewCategoryService(categoryRepository)

	return &Handlers{
		User:     handler.NewUserHandler(userService),
		Category: handler.NewCategoryHandler(categoryService),
	}
}
