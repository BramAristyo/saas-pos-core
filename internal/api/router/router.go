package router

import (
	"github.com/BramAristyo/go-pos-mawish/internal/infrastructure/config"
	"github.com/BramAristyo/go-pos-mawish/internal/dependency"
	"github.com/BramAristyo/go-pos-mawish/internal/api/middleware"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine, h *dependency.Handlers, cfg *config.Config) {
	api := r.Group("/api")
	v1 := api.Group("/v1")
	{
		users := v1.Group("/users", middleware.Authentication(cfg))
		categories := v1.Group("/categories", middleware.Authentication(cfg))
		products := v1.Group("/products", middleware.Authentication(cfg))

		v1.POST("/", h.Auth.Login)

		UserRoutes(users, h.User)
		CategoryRoutes(categories, h.Category)
		ProductRoutes(products, h.Product)
	}
}
