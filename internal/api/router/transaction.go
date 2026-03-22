package router

import (
	"github.com/BramAristyo/go-pos-mawish/internal/api/handler"
	"github.com/gin-gonic/gin"
)

func ShiftRoutes(r *gin.RouterGroup, h *handler.ShiftHandler) {
	r.GET("", h.Paginate)
	r.GET("/:id", h.FindById)
	r.GET("/open", h.FindOpenShiftByCurrentUser)
	r.POST("/open", h.OpenShift)
	r.POST("/close", h.CloseShift)
	r.PUT("/expenses", h.UpsertExpenses)
}
