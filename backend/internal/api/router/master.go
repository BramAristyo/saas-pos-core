package router

import (
	"github.com/BramAristyo/saas-pos-core/backend/internal/api/handler"
	"github.com/gin-gonic/gin"
)

func CategoryRoutes(r *gin.RouterGroup, h *handler.CategoryHandler) {
	r.GET("get-all", h.GetAll)
	r.GET("", h.Paginate)
	r.GET("/:id", h.FindById)
	r.POST("", h.Store)
	r.PUT("/:id", h.Update)
	r.DELETE("/:id", h.Delete)
	r.PATCH("/:id/restore", h.Restore)
}

func ProductRoutes(r *gin.RouterGroup, h *handler.ProductHandler) {
	r.GET("", h.Paginate)
	r.POST("", h.Store)
	r.GET("/:id", h.FindById)
	r.PUT("/:id", h.Update)
	r.DELETE("/:id", h.Delete)
	r.PATCH("/:id/restore", h.Restore)
}

func ModifierRoutes(r *gin.RouterGroup, h *handler.ModifierGroupHandler) {
	r.GET("", h.Paginate)
	r.GET("/:id", h.FindById)
	r.POST("", h.Store)
	r.PUT("/:id", h.Update)
	r.DELETE("/:id", h.Delete)
	r.PATCH("/:id/restore", h.Restore)
}

func ModifierOptionRoutes(r *gin.RouterGroup, h *handler.ModifierOptionHandler) {
	r.GET("", h.Paginate)
	r.GET("/:id", h.FindById)
	r.POST("", h.Store)
	r.PUT("/:id", h.Update)
	r.DELETE("/:id", h.Delete)
	r.PATCH("/:id/restore", h.Restore)
}

func BundlingRoutes(r *gin.RouterGroup, h *handler.BundlingHandler) {
	r.GET("", h.Paginate)
	r.GET("/:id", h.FindById)
	r.POST("", h.Store)
	r.PUT("/:id", h.Update)
	r.DELETE("/:id", h.Delete)
	r.PATCH("/:id/restore", h.Restore)
}

func TaxRoutes(r *gin.RouterGroup, h *handler.TaxHandler) {
	r.GET("", h.Paginate)
	r.GET("/:id", h.FindById)
	r.POST("", h.Store)
	r.PUT("/:id", h.Update)
	r.DELETE("/:id", h.Delete)
	r.PATCH("/:id/restore", h.Restore)
}

func DiscountRoutes(r *gin.RouterGroup, h *handler.DiscountHandler) {
	r.GET("", h.Paginate)
	r.GET("/:id", h.FindById)
	r.POST("", h.Store)
	r.PUT("/:id", h.Update)
	r.DELETE("/:id", h.Delete)
	r.PATCH("/:id/restore", h.Restore)
}

func SalesTypeRoutes(r *gin.RouterGroup, h *handler.SalesTypeHandler) {
	r.GET("", h.Paginate)
	r.GET("/:id", h.FindById)
	r.POST("", h.Store)
	r.PUT("/:id", h.Update)
	r.DELETE("/:id", h.Delete)
	r.PATCH("/:id/restore", h.Restore)
}
