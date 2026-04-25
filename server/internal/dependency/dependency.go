package dependency

import (
	"github.com/BramAristyo/saas-pos-core/server/internal/api/handler"
	"github.com/BramAristyo/saas-pos-core/server/internal/api/validation"
	"github.com/BramAristyo/saas-pos-core/server/internal/infrastructure/config"
	"github.com/BramAristyo/saas-pos-core/server/internal/repository"
	"github.com/BramAristyo/saas-pos-core/server/internal/usecase"
	"gorm.io/gorm"
)

type Handlers struct {
	Auth          *handler.AuthHandler
	User          *handler.UserHandler
	Category      *handler.CategoryHandler
	Product       *handler.ProductHandler
	ModifierGroup *handler.ModifierGroupHandler
	Bundling      *handler.BundlingHandler
	Tax           *handler.TaxHandler
	Discount      *handler.DiscountHandler
	Shift         *handler.ShiftHandler
	SalesType     *handler.SalesTypeHandler
	Order         *handler.OrderHandler
	Report        *handler.ReportHandler
	Dashboard     *handler.DashboardHandler
	Expense       *handler.ExpenseHandler
	COA           *handler.COAHandler
	Employee      *handler.EmployeeHandler
	Attendance    *handler.AttendanceHandler
}

func Bootstrap(db *gorm.DB, cfg *config.Config) *Handlers {
	validation.RegisterCustomValidators()

	userRepository := repository.NewUserRepository(db)
	auditLogRepository := repository.NewAuditLogRepository(db)

	auditLogUseCase := usecase.NewAuditLogUseCase(auditLogRepository)
	authUseCase := usecase.NewAuthUseCase(userRepository, cfg, auditLogUseCase)
	userUseCase := usecase.NewUserUseCase(userRepository, auditLogUseCase)

	categoryRepository := repository.NewCategoryRepository(db)
	categoryUseCase := usecase.NewCategoryUseCase(categoryRepository, auditLogUseCase)

	productRepository := repository.NewProductRepository(db)
	productUseCase := usecase.NewProductUseCase(productRepository, auditLogUseCase)

	modifierGroupRepository := repository.NewModifierGroupRepository(db)
	modifierGroupUseCase := usecase.NewModifierGroupUseCase(modifierGroupRepository, auditLogUseCase)

	modifierOptionRepository := repository.NewModifierOptionRepository(db)

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

	ledgerRepository := repository.NewLedgerRepository(db)

	expenseRepository := repository.NewExpenseRepository(db)
	expenseUseCase := usecase.NewExpenseUseCase(expenseRepository, ledgerRepository, auditLogUseCase)

	coaRepository := repository.NewCOARepository(db)
	coaUseCase := usecase.NewCOAUseCase(coaRepository, auditLogUseCase)

	employeeRepository := repository.NewEmployeeRepository(db)
	employeeUseCase := usecase.NewEmployeeUseCase(employeeRepository, auditLogUseCase)

	attendanceRepository := repository.NewAttendanceRepository(db)
	attendanceUseCase := usecase.NewAttendanceUseCase(attendanceRepository)

	orderRepository := repository.NewOrderRepository(db)
	orderUseCase := usecase.NewOrderUseCase(
		orderRepository,
		shiftRepository,
		salesTypeRepository,
		taxRepository,
		discountRepository,
		productRepository,
		bundlingRepository,
		modifierOptionRepository,
		auditLogUseCase,
	)

	reportUseCase := usecase.NewReportUseCase(orderRepository, shiftRepository, discountRepository)
	dashboardUseCase := usecase.NewDashboardUseCase(orderRepository)

	return &Handlers{
		Auth:          handler.NewAuthHandler(authUseCase),
		User:          handler.NewUserHandler(userUseCase),
		Category:      handler.NewCategoryHandler(categoryUseCase),
		Product:       handler.NewProductHandler(productUseCase),
		ModifierGroup: handler.NewModifierGroupHandler(modifierGroupUseCase),
		Bundling:      handler.NewBundlingHandler(bundlingUseCase),
		Tax:           handler.NewTaxHandler(taxUseCase),
		Discount:      handler.NewDiscountHandler(discountUseCase),
		Shift:         handler.NewShiftHandler(shiftUseCase),
		SalesType:     handler.NewSalesTypeHandler(salesTypeUseCase),
		Order:         handler.NewOrderHandler(orderUseCase),
		Report:        handler.NewReportHandler(reportUseCase),
		Dashboard:     handler.NewDashboardHandler(dashboardUseCase),
		Expense:       handler.NewExpenseHandler(expenseUseCase),
		COA:           handler.NewCOAHandler(coaUseCase),
		Employee:      handler.NewEmployeeHandler(employeeUseCase),
		Attendance:    handler.NewAttendanceHandler(attendanceUseCase),
	}
}
