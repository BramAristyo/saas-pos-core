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
	Product  *handler.ProductHandler
}

func Bootstrap(db *gorm.DB, cfg *config.Config) *Handlers {
	userRepository := repository.NewUserRepository(db)

	authService := service.NewAuthService(userRepository, cfg)
	userService := service.NewUserService(userRepository)

	categoryRepository := repository.NewCategoryRepository(db)
	categoryService := service.NewCategoryService(categoryRepository)

	productRepository := repository.NewProductRepository(db)
	productService := service.NewProductService(productRepository)

	return &Handlers{
		Auth:     handler.NewAuthHandler(authService),
		User:     handler.NewUserHandler(userService),
		Category: handler.NewCategoryHandler(categoryService),
		Product:  handler.NewProductHandler(productService),
	}
}
