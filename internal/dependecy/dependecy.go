package dependecy

import (
	"github.com/BramAristyo/go-pos-mawish/internal/config"
	"github.com/BramAristyo/go-pos-mawish/internal/handler"
	"github.com/BramAristyo/go-pos-mawish/internal/repository"
	"github.com/BramAristyo/go-pos-mawish/internal/service"
	"gorm.io/gorm"
)

type Handlers struct {
	Auth     *handler.AuthHandler
	User     *handler.UserHandler
	Category *handler.CategoryHandler
}

func Bootstrap(db *gorm.DB, cfg *config.Config) *Handlers {
	userRepository := repository.NewUserRepository(db)

	authService := service.NewAuthService(userRepository, cfg)
	userService := service.NewUserService(userRepository)

	categoryRepository := repository.NewCategoryRepository(db)
	categoryService := service.NewCategoryService(categoryRepository)

	return &Handlers{
		Auth:     handler.NewAuthHandler(authService),
		User:     handler.NewUserHandler(userService),
		Category: handler.NewCategoryHandler(categoryService),
	}
}
