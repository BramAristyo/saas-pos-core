package dependency

import (
	"github.com/BramAristyo/go-pos-mawish/internal/api/handler"
	"github.com/BramAristyo/go-pos-mawish/internal/infrastructure/config"
	"github.com/BramAristyo/go-pos-mawish/internal/repository"
	"github.com/BramAristyo/go-pos-mawish/internal/service"
	"gorm.io/gorm"
)

type Handlers struct {
	Auth          *handler.AuthHandler
	User          *handler.UserHandler
	Category      *handler.CategoryHandler
	Product       *handler.ProductHandler
	ModifierGroup  *handler.ModifierGroupHandler
	ModifierOption *handler.ModifierOptionHandler
}

func Bootstrap(db *gorm.DB, cfg *config.Config) *Handlers {
	userRepository := repository.NewUserRepository(db)

	authService := service.NewAuthService(userRepository, cfg)
	userService := service.NewUserService(userRepository)

	categoryRepository := repository.NewCategoryRepository(db)
	categoryService := service.NewCategoryService(categoryRepository)

	productRepository := repository.NewProductRepository(db)
	productService := service.NewProductService(productRepository)

	modifierGroupRepository := repository.NewModifierGroupRepository(db)
	modifierGroupService := service.NewModifierGroupRepository(modifierGroupRepository)

	modifierOptionRepository := repository.NewModifierOptionRepository(db)
	modifierOptionService := service.NewModifierOptionService(modifierOptionRepository)

	return &Handlers{
		Auth:           handler.NewAuthHandler(authService),
		User:           handler.NewUserHandler(userService),
		Category:       handler.NewCategoryHandler(categoryService),
		Product:        handler.NewProductHandler(productService),
		ModifierGroup:  handler.NewModifierGroupHandler(modifierGroupService),
		ModifierOption: handler.NewModifierOptionHandler(modifierOptionService),
	}
}
