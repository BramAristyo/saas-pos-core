package router

import (
	"github.com/BramAristyo/go-pos-mawish/internal/api/handler"
	"github.com/gin-gonic/gin"
)

func CategoryRoutes(r *gin.RouterGroup, h *handler.CategoryHandler) {
	r.GET("/", h.Paginate)
	r.GET("/:id", h.FindById)
	r.POST("/", h.Store)
	r.PUT("/:id", h.Update)
}
