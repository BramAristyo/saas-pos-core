package router

import (
	"github.com/BramAristyo/go-pos-mawish/internal/dependecy"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine, h *dependecy.Handlers) {
	api := r.Group("/api")
	v1 := api.Group("/v1")
	{
		users := v1.Group("/users")
		categories := v1.Group("/categories")

		UserRoutes(users, h.User)
		CategoryRoutes(categories, h.Category)
	}
}
