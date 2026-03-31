package dependency

import (
	"github.com/BramAristyo/go-pos-mawish/internal/api/handler"
	"github.com/BramAristyo/go-pos-mawish/internal/infrastructure/config"
	"github.com/BramAristyo/go-pos-mawish/internal/repository"
	"github.com/BramAristyo/go-pos-mawish/internal/usecase"
	"gorm.io/gorm"
)

type Handlers struct {
	Auth           *handler.AuthHandler
	User           *handler.UserHandler
	Category       *handler.CategoryHandler
	Product        *handler.ProductHandler
	ModifierGroup  *handler.ModifierGroupHandler
	ModifierOption *handler.ModifierOptionHandler
	Bundling       *handler.BundlingHandler
	Tax            *handler.TaxHandler
	Discount       *handler.DiscountHandler
	Shift          *handler.ShiftHandler
	SalesType      *handler.SalesTypeHandler
	Order          *handler.OrderHandler
	Report         *handler.ReportHandler
}

func Bootstrap(db *gorm.DB, cfg *config.Config) *Handlers {
	userRepository := repository.NewUserRepository(db)
	auditLogRepository := repository.NewAuditLogRepository(db)

	auditLogUseCase := usecase.NewAuditLogUseCase(auditLogRepository)
	authUseCase := usecase.NewAuthUseCase(userRepository, cfg)
	userUseCase := usecase.NewUserUseCase(userRepository, auditLogUseCase)

	categoryRepository := repository.NewCategoryRepository(db)
	categoryUseCase := usecase.NewCategoryUseCase(categoryRepository, auditLogUseCase)

	productRepository := repository.NewProductRepository(db)
	productUseCase := usecase.NewProductUseCase(productRepository, auditLogUseCase)

	modifierGroupRepository := repository.NewModifierGroupRepository(db)
	modifierGroupUseCase := usecase.NewModifierGroupUseCase(modifierGroupRepository, auditLogUseCase)

	modifierOptionRepository := repository.NewModifierOptionRepository(db)
	modifierOptionUseCase := usecase.NewModifierOptionUseCase(modifierOptionRepository, auditLogUseCase)

	bundlingRepository := repository.NewBundlingRepository(db)
	bundlingUseCase := usecase.NewBundlingUseCase(bundlingRepository, auditLogUseCase)

	taxRepository := repository.NewTaxRepository(db)
	taxUseCase := usecase.NewTaxUseCase(taxRepository, auditLogUseCase)

	discountRepository := repository.NewDiscountRepository(db)
	discountUseCase := usecase.NewDiscountUseCase(discountRepository, auditLogUseCase)

	shiftRepository := repository.NewShiftRepository(db)
	shiftUseCase := usecase.NewShiftUseCase(shiftRepository, auditLogUseCase)

	salesTypeRepository := repository.NewSalesTypeRepository(db)
	salesTypeUseCase := usecase.NewSalesTypeUseCase(salesTypeRepository, auditLogUseCase)

	orderRepository := repository.NewOrderRepository(db)
	orderUseCase := usecase.NewOrderUseCase(
		orderRepository,
		shiftRepository,
		salesTypeRepository,
		taxRepository,
		discountRepository,
		productRepository,
		bundlingRepository,
		auditLogUseCase,
	)

	reportUseCase := usecase.NewReportUseCase(orderRepository, shiftRepository, discountRepository)

	return &Handlers{
		Auth:           handler.NewAuthHandler(authUseCase),
		User:           handler.NewUserHandler(userUseCase),
		Category:       handler.NewCategoryHandler(categoryUseCase),
		Product:        handler.NewProductHandler(productUseCase),
		ModifierGroup:  handler.NewModifierGroupHandler(modifierGroupUseCase),
		ModifierOption: handler.NewModifierOptionHandler(modifierOptionUseCase),
		Bundling:       handler.NewBundlingHandler(bundlingUseCase),
		Tax:            handler.NewTaxHandler(taxUseCase),
		Discount:       handler.NewDiscountHandler(discountUseCase),
		Shift:          handler.NewShiftHandler(shiftUseCase),
		SalesType:      handler.NewSalesTypeHandler(salesTypeUseCase),
		Order:          handler.NewOrderHandler(orderUseCase),
		Report:         handler.NewReportHandler(reportUseCase),
	}
}
