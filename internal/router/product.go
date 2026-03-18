package router

import (
	"github.com/BramAristyo/go-pos-mawish/internal/handler"
	"github.com/gin-gonic/gin"
)

func ProductRoutes(r *gin.RouterGroup, h *handler.ProductHandler) {
	r.GET("/", h.Paginate)
	r.POST("/", h.Store)
	r.GET("/:id", h.FindById)
	r.PUT("/:id", h.Update)
	r.PATCH("/:id/status", h.ChangeStatus)
}
