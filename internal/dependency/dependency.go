package dependency

import (
	"github.com/BramAristyo/go-pos-mawish/internal/api/handler"
	"github.com/BramAristyo/go-pos-mawish/internal/infrastructure/config"
	"github.com/BramAristyo/go-pos-mawish/internal/repository"
	"github.com/BramAristyo/go-pos-mawish/internal/usecase"
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

	authUseCase := usecase.NewAuthUseCase(userRepository, cfg)
	userUseCase := usecase.NewUserUseCase(userRepository)

	categoryRepository := repository.NewCategoryRepository(db)
	categoryUseCase := usecase.NewCategoryUseCase(categoryRepository)

	productRepository := repository.NewProductRepository(db)
	productUseCase := usecase.NewProductUseCase(productRepository)

	modifierGroupRepository := repository.NewModifierGroupRepository(db)
	modifierGroupUseCase := usecase.NewModifierGroupUseCase(modifierGroupRepository)

	modifierOptionRepository := repository.NewModifierOptionRepository(db)
	modifierOptionUseCase := usecase.NewModifierOptionUseCase(modifierOptionRepository)

	return &Handlers{
		Auth:           handler.NewAuthHandler(authUseCase),
		User:           handler.NewUserHandler(userUseCase),
		Category:       handler.NewCategoryHandler(categoryUseCase),
		Product:        handler.NewProductHandler(productUseCase),
		ModifierGroup:  handler.NewModifierGroupHandler(modifierGroupUseCase),
		ModifierOption: handler.NewModifierOptionHandler(modifierOptionUseCase),
	}
}
