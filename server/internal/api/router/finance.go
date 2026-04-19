package router

import (
	"github.com/BramAristyo/saas-pos-core/server/internal/api/handler"
	"github.com/gin-gonic/gin"
)

func ExpenseRoutes(r *gin.RouterGroup, h *handler.ExpenseHandler) {
	r.GET("", h.Paginate)
	r.GET("/:id", h.FindById)
	r.POST("", h.Store)
	r.PUT("/:id", h.Update)
	r.DELETE("/:id", h.Delete)
	r.PATCH("/:id/restore", h.Restore)
}

func ShiftExpensesRoutes(r *gin.RouterGroup, h *handler.ShiftExpensesHandler) {
	r.GET("", h.Paginate)
	r.GET("/:id", h.FindById)
}

func COARoutes(r *gin.RouterGroup, h *handler.COAHandler) {
	r.GET("get-all", h.GetAll)
	r.GET("", h.Paginate)
	r.GET("/:id", h.FindById)
	r.POST("", h.Store)
	r.PUT("/:id", h.Update)
	r.DELETE("/:id", h.Delete)
	r.PATCH("/:id/restore", h.Restore)
}
