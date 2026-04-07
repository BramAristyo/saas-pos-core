package router

import (
	"github.com/BramAristyo/saas-pos-core/backend/internal/api/handler"
	"github.com/gin-gonic/gin"
)

func UserRoutes(r *gin.RouterGroup, h *handler.UserHandler) {
	r.GET("", h.GetAll)
	r.GET("/:id", h.FindById)
	r.POST("", h.Store)
	r.PUT("/:id", h.Update)
	r.DELETE("/:id", h.Delete)
	r.PATCH("/:id/restore", h.Restore)
}
