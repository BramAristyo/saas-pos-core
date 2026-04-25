package router

import (
	"github.com/BramAristyo/saas-pos-core/server/internal/api/handler"
	"github.com/gin-gonic/gin"
)

func AttendanceRoutes(r *gin.RouterGroup, h *handler.AttendanceHandler) {
	r.GET("", h.Paginate)
}
