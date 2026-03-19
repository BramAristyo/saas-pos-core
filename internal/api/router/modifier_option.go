package router

import (
	"github.com/BramAristyo/go-pos-mawish/internal/api/handler"
	"github.com/gin-gonic/gin"
)

func ModifierOptionRoutes(r *gin.RouterGroup, h *handler.ModifierOptionHandler) {
	r.GET("/", h.Paginate)
	r.GET("/:id", h.FindById)
	r.POST("/", h.Store)
	r.PUT("/:id", h.Update)
	r.PATCH("/:id/activate", h.Activate)
	r.PATCH("/:id/deactivate", h.Deactivate)
}
