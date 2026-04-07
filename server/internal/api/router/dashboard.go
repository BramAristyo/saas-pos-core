package router

import (
	"github.com/BramAristyo/saas-pos-core/server/internal/api/handler"
	"github.com/gin-gonic/gin"
)

func DashboardRoutes(r *gin.RouterGroup, h *handler.DashboardHandler) {
	r.GET("/sales-summary", h.SalesSummary)
}
